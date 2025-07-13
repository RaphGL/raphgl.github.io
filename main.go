package main

import (
	"errors"
	"fmt"
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

func getMetadataHeader(contents string) (header string, headerEnd int) {
	const DashLen = 3
	if string(contents[:DashLen]) != "---" {
		return "", -1
	}
	contents = contents[DashLen:]
	headerEnd = strings.Index(contents, "---")
	if headerEnd == -1 {
		return "", -1
	}

	header = string(contents[:headerEnd])
	// exclude both the beginning and ending dashes
	headerEnd += DashLen * 2
	return
}

func appendPageHeader(contents string, header string) (string, error) {
	_ = contents
	for entry := range strings.SplitSeq(header, "\n") {
		entry = strings.TrimSpace(entry)
		if len(entry) == 0 {
			continue
		}

		entrySep := strings.Index(entry, ":")
		if entrySep == -1 {
			// TODO track what file stuff comes from
			return "", errors.New("expected `:` separated key-values in metadata header")
		}
		key := entry[:entrySep]
		// value := entry[entrySep+1:]

		switch key {
		case "title":
		case "description":
		default:
			return "", errors.New("invalid key in metadata header")
		}
	}

	return "", nil
}

func compileToHTML(mdFile string) ([]byte, error) {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	contentsBytes, err := os.ReadFile(mdFile)
	if err != nil {
		return nil, err
	}
	contents := string(contentsBytes)

	metadata, metadataEnd := getMetadataHeader(contents)
	if metadataEnd == -1 {
		return nil, errors.New("could not find a metadata header")
	}
	contents = contents[metadataEnd:]

	parsedMd := p.Parse([]byte(contents))

	// TODO: set more options
	htmlOpts := html.RendererOptions{
		Flags: html.CommonFlags,
	}

	renderer := html.NewRenderer(htmlOpts)
	contentsHtml := string(markdown.Render(parsedMd, renderer))
	contentsHtml, err = appendPageHeader(contentsHtml, metadata)
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
