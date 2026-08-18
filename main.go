package main

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
const WebURL = "https://raphgl.github.io"

// contains the base html body to which the body will be added to
type Base struct {
	Post     Post
	BodyHTML template.HTML
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

func GetStyles() (string, error) {
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

	// if on windows it will use backslash instead so we need to convert it
	parentPath = filepath.ToSlash(parentPath)
	destPath = filepath.ToSlash(destPath)
	return
}

func CopyStaticFile(path string) error {
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

func GenerateRSSFromPosts(posts []Post) string {
	var rss strings.Builder
	rss.WriteString(`<?xml version="1.0" encoding="UTF-8" ?>`)
	rss.WriteString(`<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:content="http://purl.org/rss/1.0/modules/content/"><channel>`)
	{
		rss.WriteString(`<atom:link href="`)
		rss.WriteString(WebURL + "/rss.xml")
		rss.WriteString(`" rel="self" type="application/rss+xml" />`)

		rss.WriteString("<title>RaphGL</title>")
		rss.WriteString("<link>" + WebURL + "</link>")
		rss.WriteString("<description>RaphGL's Blog</description>")

		rss.WriteString("<pubDate>")
		rss.WriteString(time.Now().UTC().Format(time.RFC1123Z))
		rss.WriteString("</pubDate>")

		for idx, post := range posts {
			// to avoid rss bloating the feed only the most recent posts are shown
			if idx >= 20 {
				break
			}

			rss.WriteString("<item>")
			{
				rss.WriteString("<title>")
				rss.WriteString(post.Title)
				rss.WriteString("</title>")

				postURL := WebURL + post.SourceFilePath
				rss.WriteString("<link>")
				rss.WriteString(postURL)
				rss.WriteString("</link>")

				rss.WriteString("<guid>")
				rss.WriteString(postURL)
				rss.WriteString("</guid>")

				parsedDate, err := time.Parse("2006-01-02", post.Date)
				if err == nil {
					rss.WriteString("<pubDate>")
					rss.WriteString(parsedDate.UTC().Format(time.RFC1123Z))
					rss.WriteString("</pubDate>")
				}

				rss.WriteString("<description>")
				rss.WriteString(post.Description)
				rss.WriteString("</description>")

				rss.WriteString("<content:encoded><![CDATA[")
				rss.WriteString(post.Content)
				rss.WriteString("]]></content:encoded>")
			}
			rss.WriteString("</item>")
		}
	}
	rss.WriteString("</channel></rss>")
	return rss.String()
}

func main() {
	// we remove all files in target dir first to prevent previous
	// compilation items from being left in the final website artifacts
	os.RemoveAll(TargetDirName)

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

	// === Taxonomize all website files ===
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
	filepath.WalkDir("./static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if d.IsDir() {
			return nil
		}

		staticFiles = append(staticFiles, path)
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
			if post.Draft {
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

	rssFeed := GenerateRSSFromPosts(postList.Posts)
	os.WriteFile(TargetDirName+"/rss.xml", []byte(rssFeed), 0664)

	listHTML, err := postList.Render()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = os.WriteFile(TargetDirName+"/index.html", []byte(listHTML), 0644); err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range staticFiles {
		if err := CopyStaticFile(file); err != nil {
			fmt.Println(err)
		}
	}

	styles, err := GetStyles()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.WriteFile(TargetDirName+"/styles.css", []byte(styles), 0644); err != nil {
		fmt.Println(err)
		return
	}
}
