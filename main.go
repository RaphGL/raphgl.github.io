package main

// TODO: generate blog list page
// TODO: generate home page

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/alecthomas/chroma/v2"
	formatterHtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const TargetDirName = "docs"

type PostHeader struct {
	Title       string
	Description string
	Date        string
}

func getMetadataHeader(contents string) (header PostHeader, headerEnd int, err error) {
	const DashLen = 3
	if string(contents[:DashLen]) != "---" {
		return PostHeader{}, -1, errors.New("missing header")
	}
	contents = contents[DashLen:]
	headerEnd = strings.Index(contents, "---")
	if headerEnd == -1 {
		return PostHeader{}, -1, errors.New("missing header terminator")
	}

	headerStr := string(contents[:headerEnd])
	// exclude both the beginning and ending dashes
	headerEnd += DashLen * 2

	for entry := range strings.SplitSeq(headerStr, "\n") {
		entry = strings.TrimSpace(entry)
		if len(entry) == 0 {
			continue
		}

		entrySep := strings.Index(entry, ":")
		if entrySep == -1 {
			// TODO track what file stuff comes from
			return PostHeader{}, -1, errors.New("expected `:` separated key-values in metadata header")
		}
		key := strings.ToLower(strings.TrimSpace(entry[:entrySep]))
		value := strings.TrimSpace(entry[entrySep+1:])

		switch key {
		case "title":
			header.Title = value
		case "description":
			header.Description = value
		case "date":
			header.Date = value
		default:
			return PostHeader{}, -1, errors.New("invalid key in metadata header")
		}
	}

	return
}

func getStyles() (template.CSS, error) {
	styles, err := os.ReadFile("./layout/styles.css")
	if err != nil {
		return "", err
	}
	cssReset, err := os.ReadFile("./layout/reset.css")
	if err != nil {
		return "", err
	}

	finalStyles := fmt.Sprintln(string(cssReset), string(styles))
	return template.CSS(finalStyles), nil
}

func getHeaderHTML(header PostHeader) (template.HTML, error) {
	// TODO: generate date
	headerTempl, err := os.ReadFile("./layout/header.html")
	if err != nil {
		return "", err
	}
	templ, err := template.New("header").Parse(string(headerTempl))
	if err != nil {
		return "", err
	}

	var headerBuilder strings.Builder
	if err := templ.Execute(&headerBuilder, header); err != nil {
		return "", err
	}
	return template.HTML(headerBuilder.String()), nil
}

func getPageHTML(contents string, header PostHeader) (template.HTML, error) {
	type Body struct {
		Header     PostHeader
		HeaderHTML template.HTML
		PostHTML   template.HTML
		StylesCSS  template.CSS
	}

	headerHTML, err := getHeaderHTML(header)
	if err != nil {
		return "", err
	}
	stylesCSS, err := getStyles()
	if err != nil {
		return "", err
	}

	bodyFields := Body{
		Header:     header,
		HeaderHTML: headerHTML,
		PostHTML:   template.HTML(contents),
		StylesCSS:  stylesCSS,
	}

	indexTempl, err := os.ReadFile("./layout/index.html")
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
	return template.HTML(bodyBuilder.String()), nil
}

func compileToHTML(mdFile string) (template.HTML, error) {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	contentsBytes, err := os.ReadFile(mdFile)
	if err != nil {
		return "", err
	}
	contents := string(contentsBytes)

	header, headerEnd, err := getMetadataHeader(contents)
	if err != nil {
		return "", err
	}
	contents = contents[headerEnd:]

	parsedMd := p.Parse([]byte(contents))

	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatterHtml.New(formatterHtml.WithClasses(true))

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

				// TODO: refactor code so that we can append a single highlight CSS
				// in the page's header
				w.Write([]byte("<style>"))
				formatter.WriteCSS(w, style)
				w.Write([]byte("</style>"))

				return ast.GoToNext, true
			}

			return ast.GoToNext, false
		},
	}

	renderer := html.NewRenderer(htmlOpts)
	postContents := string(markdown.Render(parsedMd, renderer))
	page, err := getPageHTML(postContents, header)
	return page, err
}

func main() {
	dir := "./content"
	dirStat, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !dirStat.IsDir() {
		fmt.Println("Argument has to be a directory.")
		return
	}

	defer func() {
		targetDir, err := os.Open(TargetDirName)
		if err != nil {
			fmt.Println(targetDir)
		}

		dirnames, err := targetDir.Readdirnames(1)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(dirnames) == 0 {
			os.Remove(TargetDirName)
		}
	}()

	files := make([]string, 1)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		files = append(files, path)
		return nil
	})

	for _, filePath := range files {
		// for now all files are independent of each other so it's easily parallelized
		// TODO: in the future we're going to build a blog list so we'll have to use a mutex for that
		go func() {
			htmlArtifact, err := compileToHTML(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}

			pathComponents := strings.Split(filePath, string(filepath.Separator))[1:]
			destComponents := slices.Insert(pathComponents, 0, TargetDirName)

			destPath := strings.Join(destComponents, string(filepath.Separator))
			destExt := filepath.Ext(destPath)
			destPath = destPath[:len(destPath)-len(destExt)] + ".html"
			parentDirPath := strings.Join(destComponents[:len(destComponents)-1], string(filepath.Separator))

			if err := os.MkdirAll(parentDirPath, 0755); err != nil {
				fmt.Println(err)
				return
			}

			if err := os.WriteFile(destPath, []byte(htmlArtifact), 0644); err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
}
