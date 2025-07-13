package main

// TODO: generate blog list page
// TODO: generate home page

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gomarkdown/markdown"
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

// TODO: read header from a file
func appendPageHeader(contents string, header PostHeader) (string, error) {
	var opts struct {
		Header PostHeader
		Post   template.HTML
	}

	opts.Header = header
	opts.Post = template.HTML(contents)

	headerTempl := `
	<div id="post-header">
		{{if .Header.Title}}
			<div id="post-header-title">
				{{.Header.Title}}
			</div>
		{{end}}

		{{if .Header.Description}}
			<div id="post-header-description">
				{{.Header.Description}}
			</div>
		{{end}}
	</div>

	<div id="post-content">
		{{.Post}}
	</div>
	`
	templ, err := template.New("header").Parse(headerTempl)
	if err != nil {
		return "", err
	}

	var headerBuilder strings.Builder
	if err := templ.Execute(&headerBuilder, opts); err != nil {
		return "", err
	}
	return headerBuilder.String(), nil
}

// TODO: read body from a file
// TODO: include css file in html
func appendBody(contents string, header PostHeader) (string, error) {
	type Body struct {
		Header PostHeader
		Child  template.HTML
	}

	bodyFields := Body{
		Header: header,
		Child:  template.HTML(contents),
	}

	bodyTemplate := `
	<!DOCTYPE html>
	<html lang="en">

	<head>
	  <meta charset="utf-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1">
	  <title>RaphGL | {{.Header.Title}}</title>
	</head>
	<body>
		{{.Child}}
	</body>
	</html>
	`

	templ, err := template.New("body").Parse(bodyTemplate)
	if err != nil {
		return "", err
	}

	var bodyBuilder strings.Builder
	if err := templ.Execute(&bodyBuilder, bodyFields); err != nil {
		return "", err
	}
	return bodyBuilder.String(), nil
}

func compileToHTML(mdFile string) ([]byte, error) {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	contentsBytes, err := os.ReadFile(mdFile)
	if err != nil {
		return nil, err
	}
	contents := string(contentsBytes)

	header, headerEnd, err := getMetadataHeader(contents)
	if err != nil {
		return nil, err
	}
	contents = contents[headerEnd:]

	parsedMd := p.Parse([]byte(contents))

	// TODO: set more options
	htmlOpts := html.RendererOptions{
		Flags: html.CommonFlags,
	}

	renderer := html.NewRenderer(htmlOpts)
	contentsHtml := string(markdown.Render(parsedMd, renderer))
	contentsHtml, err = appendPageHeader(contentsHtml, header)
	if err != nil {
		return nil, err
	}
	contentsHtml, err = appendBody(contentsHtml, header)
	return []byte(contentsHtml), err
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("expected more args")
		return
	}

	dir := os.Args[1]
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

		htmlArtifact, err := compileToHTML(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		pathComponents := strings.Split(path, string(filepath.Separator))[1:]
		destComponents := slices.Insert(pathComponents, 0, TargetDirName)

		destPath := strings.Join(destComponents, string(filepath.Separator))
		destExt := filepath.Ext(destPath)
		destPath = destPath[:len(destPath)-len(destExt)] + ".html"
		parentDirPath := strings.Join(destComponents[:len(destComponents)-1], string(filepath.Separator))

		if err := os.MkdirAll(parentDirPath, 0755); err != nil {
			fmt.Println(err)
			return nil
		}

		if err := os.WriteFile(destPath, htmlArtifact, 0644); err != nil {
			fmt.Println(err)
			return nil
		}

		return nil
	})
}
