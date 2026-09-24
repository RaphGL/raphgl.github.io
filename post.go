package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/alecthomas/chroma/v2"
	fmtHTML "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func EstimateReadTime(content string) time.Duration {
	const WordsPerMinute = 220
	totalWords := 0
	for word := range strings.SplitSeq(content, " ") {
		if len(word) == 0 {
			continue
		}

		r, rSize := utf8.DecodeRuneInString(word)
		if len(word) == 1 && rSize >= 1 {
			if unicode.IsSymbol(r) {
				continue
			}

			if unicode.IsSpace(r) {
				continue
			}
		}

		totalWords += 1
	}

	return time.Duration(totalWords/WordsPerMinute) * time.Minute
}

type CheckerHook func(string, ast.Node) error

type Post struct {
	// hooks that are used just to inspect and verify markdown tokens
	checkerHooks []CheckerHook

	SourceFilePath     string
	Title              string
	Description        string
	ReadDuration       string
	Date               string
	Draft              bool
	Content            string
	SyntaxHighlightCSS template.CSS
}

func NewPost(path string) (Post, error) {
	contentsBytes, err := os.ReadFile(path)
	if err != nil {
		return Post{}, err
	}
	contents := string(contentsBytes)

	const HeaderSeparator = "---"
	const DashLen = len(HeaderSeparator)
	if string(contents[:DashLen]) != HeaderSeparator {
		return Post{}, errors.New("missing header in " + path)
	}
	contents = contents[DashLen:]
	headerEnd := strings.Index(contents, HeaderSeparator)
	if headerEnd == -1 {
		return Post{}, errors.New("missing header terminator in " + path)
	}

	headerStr := string(contents[:headerEnd])
	// exclude both the beginning and ending dashes
	headerEnd += DashLen*2 - 1

	post := Post{
		// content without header
		SourceFilePath: path,
	}
	if len(contents) > headerEnd {
		post.Content = contents[headerEnd:]
	}
	post.ReadDuration = fmt.Sprintf("%d min", EstimateReadTime(post.Content)/time.Minute)

	// == Parsing md header
	for entry := range strings.SplitSeq(headerStr, "\n") {
		entry = strings.TrimSpace(entry)
		if len(entry) == 0 {
			continue
		}

		key, value, hasSep := strings.Cut(entry, ":")
		if !hasSep {
			return Post{}, errors.New("expected `:` separated key-values in metadata header in " + post.SourceFilePath)
		}
		key = strings.ToLower(strings.TrimSpace(key))
		value = strings.TrimSpace(value)

		switch key {
		case "title":
			post.Title = value
		case "description":
			post.Description = value
		case "date":
			post.Date = value
		case "draft":
			isDraft, err := strconv.ParseBool(value)
			if err != nil {
				return Post{}, fmt.Errorf("expected `draft` to be a bool, found `%v`", value)
			}
			post.Draft = isDraft
		default:
			return Post{}, errors.New("invalid key '" + key + "' in metadata header")
		}
	}

	_, err = time.Parse("2006-01-02", post.Date)
	if err != nil || len(post.Date) == 0 {
		return Post{}, fmt.Errorf("`%v` is missing date of in the format YYYY-MM-DD: %v", path, err)
	}

	return post, nil
}

func (p *Post) AddCheckerHook(hook CheckerHook) {
	p.checkerHooks = append(p.checkerHooks, hook)
}

func (p Post) getHeaderHTML() (template.HTML, error) {
	headerTempl, err := os.ReadFile("./layout/header.html")
	if err != nil {
		return "", err
	}
	templ, err := template.New("header").Parse(string(headerTempl))
	if err != nil {
		return "", err
	}

	var headerBuilder strings.Builder
	if err := templ.Execute(&headerBuilder, p); err != nil {
		return "", err
	}
	return template.HTML(headerBuilder.String()), nil
}

func (p Post) getPostHTML() (template.HTML, error) {
	type Body struct {
		HeaderHTML template.HTML
		PostHTML   template.HTML
	}

	headerHTML, err := p.getHeaderHTML()
	if err != nil {
		return "", err
	}

	bodyFields := Body{
		HeaderHTML: headerHTML,
		PostHTML:   template.HTML(p.Content),
	}

	indexTempl, err := os.ReadFile("./layout/post.html")
	if err != nil {
		return "", err
	}

	templ, err := template.New("body").Parse(string(indexTempl))
	if err != nil {
		return "", err
	}

	var bodyBuilder strings.Builder
	if err := templ.Execute(&bodyBuilder, bodyFields); err != nil {
		return "", err
	}

	page := Base{
		Post:     p,
		BodyHTML: template.HTML(bodyBuilder.String()),
	}
	pageHTML, err := page.Render()
	if err != nil {
		return "", err
	}
	return pageHTML, nil
}

func renderCodeblock(w io.Writer, formatter *fmtHTML.Formatter, style *chroma.Style, code *ast.CodeBlock) bool {
	// we're trimming because if there's trailing spaces syntax highlighting stops working
	lang := strings.TrimSpace(string(code.Info))
	lexer := lexers.Get(lang)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	iter, err := lexer.Tokenise(nil, string(code.Literal))
	if err != nil {
		return false
	}
	if formatter.Format(w, style, iter) != nil {
		return false
	}

	return true
}

type PostSection struct {
	Title string
	ID    string
	Level int
}

// consumes the sections one by one and returns the generated HTML string
// and a new post section slice from where it stopped consuming the sections
//
// TODO: rewrite function so not be recursive anymore.
// We can easily implement this functionality by simply detecting when the `level` changes.
// once that happens we can insert insert a <ul> tag
func consumePostSectionInLevel(level int, sections []PostSection) (string, []PostSection) {
	var sb strings.Builder
	sb.WriteString("<ul>")

	var i int
	foundSibling := false
	for i = range len(sections) {
		if i >= len(sections) {
			break
		}
		s := sections[i]
		if s.Level < level {
			foundSibling = true
			break
		}

		if s.Level > level {
			var contents string
			contents, sections = consumePostSectionInLevel(s.Level, sections[i:])
			sb.WriteString(contents)
			continue
		}

		sb.WriteString("<li>")
		{
			sb.WriteString(`<a href="#`)
			sb.WriteString(s.ID)
			sb.WriteString(`">`)
			sb.WriteString(s.Title)
			sb.WriteString("</a>")
		}
		sb.WriteString("</li>")
	}

	sb.WriteString("</ul>")
	if !foundSibling {
		i++
		if i < len(sections) {
			sections = sections[i:]
		} else {
			sections = nil
		}

	}
	return sb.String(), sections
}

type TableOfContents struct {
	TOC template.HTML
}

func renderTableOfContents(sections []PostSection) (string, error) {
	contents, _ := consumePostSectionInLevel(1, sections)

	tocTempl, err := os.ReadFile("./layout/table-of-contents.html")
	if err != nil {
		return "", err
	}

	templ, err := template.New("toc").Parse(string(tocTempl))
	if err != nil {
		return "", err
	}
	toc := TableOfContents{
		TOC: template.HTML(contents),
	}

	var sb strings.Builder
	if err := templ.Execute(&sb, toc); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func (post *Post) Render() (template.HTML, error) {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	// === Add syntax highlighting for code blocks ===
	formatter, style := GetSyntaxHighlighter()

	var postSections []PostSection

	htmlOpts := html.RendererOptions{
		Flags: html.CommonFlags,
		RenderNodeHook: func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
			if !entering {
				return ast.GoToNext, false
			}

			var ok bool
			switch _node := node.(type) {
			case *ast.CodeBlock:
				ok = renderCodeblock(w, formatter, style, _node)

			case *ast.Heading:
				if len(_node.Children) != 0 {
					title := string(_node.Children[0].AsLeaf().Literal)
					id := strings.ReplaceAll(strings.ToLower(title), " ", "-")
					_node.HeadingID = id
					postSections = append(postSections, PostSection{
						Title: title,
						ID:    id,
						Level: _node.Level,
					})
				}
				ok = false
			}
			return ast.GoToNext, ok

		},
	}

	// === convert post from Markdown to HTML ===
	renderer := html.NewRenderer(htmlOpts)
	parsedMd := p.Parse([]byte(post.Content))

	// checker hooks
	var allErrs error
	ast.WalkFunc(parsedMd, func(node ast.Node, entering bool) ast.WalkStatus {
		for _, hook := range post.checkerHooks {
			if !entering {
				return ast.GoToNext
			}

			if err := hook(post.SourceFilePath, node); err != nil {
				allErrs = errors.Join(allErrs, err)
				return ast.GoToNext
			}
		}

		return ast.GoToNext
	})
	if allErrs != nil {
		return "", allErrs
	}

	postContents := string(markdown.Render(parsedMd, renderer))

	post.Content = postContents
	if len(postSections) != 0 {
		toc, err := renderTableOfContents(postSections)
		if err != nil {
			return "", err
		}
		post.Content = toc + post.Content
	}
	page, err := post.getPostHTML()
	return page, err
}
