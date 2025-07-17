package main

// TODO: generate home page

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

// TODO: make paths absolute and then do the resolution so that it works
// consistently
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

	files := make([]string, 0)
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		files = append(files, path)
		return nil
	})

	// === Generate posts ===
	var wg sync.WaitGroup
	renderedPosts := make([]Post, 0)
	// locks writes to renderedPosts
	var rendMux sync.Mutex
	for _, filePath := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			post, err := NewPost(filePath)
			if err != nil {
				fmt.Println(err)
				return
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
}
