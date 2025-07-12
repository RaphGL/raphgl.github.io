package main

import (
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

func compileToHTML(mdFile string) ([]byte, error) {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)

	contents, err := os.ReadFile(mdFile)
	if err != nil {
		return []byte{}, err
	}

	parsedMd := p.Parse(contents)

	// TODO: set more options
	htmlOpts := html.RendererOptions{
		Flags: html.CommonFlags,
	}

	renderer := html.NewRenderer(htmlOpts)
	return markdown.Render(parsedMd, renderer), nil
}

func main() {
	const TargetDirName = "docs"

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

	// compiled, err := compileToHTML(file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
