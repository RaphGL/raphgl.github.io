package main

// TODO: generate blog list page
// TODO: generate home page

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

const TargetDirName = "docs"

func GetGlobalStyles() (string, error) {
	cssReset, err := os.ReadFile("./layout/reset.css")
	if err != nil {
		return "", err
	}
	styles, err := os.ReadFile("./layout/styles.css")
	if err != nil {
		return "", err
	}

	return fmt.Sprintln(string(cssReset), string(styles)), nil
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

	var wg sync.WaitGroup
	for _, filePath := range files {
		// for now all files are independent of each other so it's easily parallelized
		// TODO: in the future we're going to build a blog list so we'll have to use a mutex for that
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
	wg.Wait()
}
