package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2"
	formatterHTML "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	SourceFilePath     string
	Title              string
	Description        string
	Date               string
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
		return Post{}, errors.New("missing header")
	}
	contents = contents[DashLen:]
	headerEnd := strings.Index(contents, HeaderSeparator)
	if headerEnd == -1 {
		return Post{}, errors.New("missing header terminator")
	}

	headerStr := string(contents[:headerEnd])
	// exclude both the beginning and ending dashes
	headerEnd += DashLen*2 - 1

	post := Post{
		// content without header
		Content:        contents[headerEnd:],
		SourceFilePath: path,
	}

	for entry := range strings.SplitSeq(headerStr, "\n") {
		entry = strings.TrimSpace(entry)
		if len(entry) == 0 {
			continue
		}

		entrySep := strings.Index(entry, ":")
		if entrySep == -1 {
			return Post{}, fmt.Errorf("expected `:` separated key-values in metadata header in `%v`", post.SourceFilePath)
		}
		key := strings.ToLower(strings.TrimSpace(entry[:entrySep]))
		value := strings.TrimSpace(entry[entrySep+1:])

		switch key {
		case "title":
			post.Title = value
		case "description":
			post.Description = value
		case "date":
			post.Date = value
		default:
			return Post{}, errors.New("invalid key in metadata header")
		}
	}

	return post, nil
}

func (p Post) getStyles() (template.CSS, error) {
	styles, err := GetGlobalStyles()
	if err != nil {
		return "", err
	}

	finalStyles := fmt.Sprintln(styles, p.SyntaxHighlightCSS)
	return template.CSS(finalStyles), nil
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

	stylesCSS, err := p.getStyles()
	if err != nil {
		return "", err
	}
	page := Base{
		Post:      p,
		BodyHTML:  template.HTML(bodyBuilder.String()),
		StylesCSS: stylesCSS,
	}
	pageHTML, err := page.Render()
	if err != nil {
		return "", err
	}
	return pageHTML, nil
}

func (post *Post) Render() (template.HTML, error) {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	// === Add syntax highlighting for code blocks ===
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatterHTML.New(formatterHTML.WithClasses(true))
	var cssBuilder strings.Builder
	if err := formatter.WriteCSS(&cssBuilder, style); err != nil {
		return "", err
	}
	post.SyntaxHighlightCSS = template.CSS(cssBuilder.String())

	htmlOpts := html.RendererOptions{
		Flags: html.CommonFlags,
		// TODO: separate closure from here to a potential file that has all of our ast checkers modifiers whatever
		RenderNodeHook: func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
			if code, ok := node.(*ast.CodeBlock); ok && entering {
				// we're trimming because if there's trailing spaces syntax highlighting stops working
				lang := strings.TrimSpace(string(code.Info))
				lexer := lexers.Get(lang)
				if lexer == nil {
					lexer = lexers.Fallback
				}
				lexer = chroma.Coalesce(lexer)

				it, err := lexer.Tokenise(nil, string(code.Literal))
				if err != nil {
					return ast.GoToNext, false
				}
				if formatter.Format(w, style, it) != nil {
					return ast.GoToNext, false
				}

				return ast.GoToNext, true
			}

			return ast.GoToNext, false
		},
	}

	// === convert post from Markdown to HTML ===
	renderer := html.NewRenderer(htmlOpts)
	parsedMd := p.Parse([]byte(post.Content))
	postContents := string(markdown.Render(parsedMd, renderer))
	post.Content = postContents
	page, err := post.getPostHTML()
	return page, err
}
