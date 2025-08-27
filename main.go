package main

// TODO: figure out a way to use images and whatnot with /static/ or any other way
// TODO: add flag to disable deadlink checker so that we don't get limited by servers while developing this generator

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

const TargetDirName = "docs"

// contains the base html body to which the body will be added to
type Base struct {
	Post      Post
	BodyHTML  template.HTML
	StylesCSS template.CSS
}

func (b Base) Render() (template.HTML, error) {
	indexTempl, err := os.ReadFile("./layout/index.html")
	if err != nil {
		return "", err
	}
	templ, err := template.New("index").Parse(string(indexTempl))
	if err != nil {
		return "", err
	}

	type BaseExt struct {
		Base
		BuildYear int
	}

	bExt := BaseExt{
		Base:      b,
		BuildYear: time.Now().Year(),
	}

	var indexBuilder strings.Builder
	if err = templ.Execute(&indexBuilder, bExt); err != nil {
		return "", err
	}
	return template.HTML(indexBuilder.String()), nil
}

func GetGlobalStyles() (template.CSS, error) {
	cssReset, err := os.ReadFile("./layout/reset.css")
	if err != nil {
		return "", err
	}
	styles, err := os.ReadFile("./layout/styles.css")
	if err != nil {
		return "", err
	}

	return template.CSS(fmt.Sprintln(string(cssReset), string(styles))), nil
}

func GetTargetPath(path string) (parentPath, destPath string) {
	pathComponents := strings.Split(path, string(filepath.Separator))[1:]
	destComponents := slices.Insert(pathComponents, 0, TargetDirName)
	parentPath = strings.Join(destComponents[:len(destComponents)-1], string(filepath.Separator))
	destPath = strings.Join(destComponents, string(filepath.Separator))
	return
}

func GetCompiledTargetPath(path string) (parentPath, destPath string) {
	pathComponents := strings.Split(path, string(filepath.Separator))[1:]
	destComponents := slices.Insert(pathComponents, 0, TargetDirName)

	destPath = strings.Join(destComponents, string(filepath.Separator))
	destExt := filepath.Ext(destPath)
	destPath = destPath[:len(destPath)-len(destExt)] + ".html"
	parentPath = strings.Join(destComponents[:len(destComponents)-1], string(filepath.Separator))
	return
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
			fmt.Println(err)
			return
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

	mdFiles := make([]string, 0)
	staticFiles := make([]string, 0)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			staticFiles = append(staticFiles, path)
			return nil
		}

		mdFiles = append(mdFiles, path)
		return nil
	})

	// === Generate posts ===
	var wg sync.WaitGroup
	renderedPosts := make([]Post, 0)
	// locks writes to renderedPosts
	var rendMux sync.Mutex
	for _, filePath := range mdFiles {
		wg.Add(1)
		go func() {
			defer wg.Done()
			post, err := NewPost(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Post Hooks
			{
				post.AddCheckerHook(CheckLinkIsReachable)
			}

			htmlArtifact, err := post.Render()
			if err != nil {
				fmt.Println(err)
				return
			}

			parentDirPath, destPath := GetCompiledTargetPath(filePath)

			if err := os.MkdirAll(parentDirPath, 0755); err != nil {
				fmt.Println(err)
				return
			}

			if err := os.WriteFile(destPath, []byte(htmlArtifact), 0644); err != nil {
				fmt.Println(err)
				return
			}

			rendMux.Lock()
			renderedPosts = append(renderedPosts, post)
			rendMux.Unlock()
		}()
	}
	wg.Wait()

	// TODO: consider rendering different contents dirs separately
	postList, err := NewList(renderedPosts)
	if err != nil {
		fmt.Println(err)
		return
	}
	listHTML, err := postList.Render()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = os.WriteFile("./docs/index.html", []byte(listHTML), 0644); err != nil {
		fmt.Println(err)
		return
	}

	// === Copy static files ===
	copyStaticFile := func(path string) error {
		contents, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		parentPath, destPath := GetTargetPath(path)
		if err := os.MkdirAll(parentPath, 0755); err != nil {
			fmt.Println(err)
			return nil
		}

		if err := os.WriteFile(destPath, contents, 0755); err != nil {
			fmt.Println(err)
			return nil
		}

		return nil
	}

	filepath.WalkDir("./static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if d.IsDir() {
			os.MkdirAll(path, 0)
			return nil
		}

		if err := copyStaticFile(path); err != nil {
			fmt.Println(err)
			return nil
		}
		return nil
	})
	for _, file := range staticFiles {
		if err := copyStaticFile(file); err != nil {
			fmt.Println(err)
		}
	}
}
