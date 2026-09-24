package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/pprof"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/alecthomas/chroma/v2"
	fmtHTML "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
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

func GetSyntaxHighlighter() (*fmtHTML.Formatter, *chroma.Style) {
	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}
	return fmtHTML.New(fmtHTML.WithClasses(true)), style
}

func GetStyles() (string, error) {
	cssReset, err := os.ReadFile("./layout/reset.css")
	if err != nil {
		return "", err
	}
	bodyStyles, err := os.ReadFile("./layout/styles.css")
	if err != nil {
		return "", err
	}

	formatter, style := GetSyntaxHighlighter()
	var cssBuilder strings.Builder
	if err := formatter.WriteCSS(&cssBuilder, style); err != nil {
		return "", err
	}

	return fmt.Sprintln(string(cssReset), string(bodyStyles), cssBuilder.String()), nil
}

// TODO fix
func GetTargetPath(path string) (parentPath, destPath string) {
	pathComponents := strings.Split(path, string(filepath.Separator))[1:]
	destComponents := slices.Insert(pathComponents, 0, TargetDirName)
	parentPath = strings.Join(destComponents[:len(destComponents)-1], string(filepath.Separator))
	destPath = strings.Join(destComponents, string(filepath.Separator))
	return
}

// TODO fix
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

func CopyStaticFile(to, from string) error {
	contents, err := os.ReadFile(from)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(to), 0755); err != nil {
		return err
	}

	if err := os.WriteFile(to, contents, 0755); err != nil {
		return err
	}

	return nil
}

func GenerateWebsite(isDebugMode bool, targetDirPath string) {
	// we remove all files in target dir first to prevent previous
	// compilation items from being left in the final website artifacts
	os.RemoveAll(targetDirPath)

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
		targetDir, err := os.Open(targetDirPath)
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
			os.Remove(targetDirPath)
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
		wg.Go(func() {
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
				if !isDebugMode {
					post.AddCheckerHook(CheckLinkIsReachable)
				}
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
		})
	}
	wg.Wait()

	// TODO: consider rendering different contents dirs separately
	postList, err := NewList(renderedPosts)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile(filepath.Join(targetDirPath, "rss.xml"), []byte(postList.GenerateRSSFromPosts()), 0664)

	listHTML, err := postList.Render()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = os.WriteFile(filepath.Join(targetDirPath, "index.html"), []byte(listHTML), 0644); err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range staticFiles {
		_, destPath := GetTargetPath(file)
		if err := CopyStaticFile(destPath, file); err != nil {
			fmt.Println(err)
		}
	}

	styles, err := GetStyles()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.WriteFile(filepath.Join(targetDirPath, "styles.css"), []byte(styles), 0644); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	devF := flag.Bool("dev", false, "Enable development mode")
	helpF := flag.Bool("help", false, "Show help message")
	profileF := flag.Bool("profile", false, "Generate program profile data")
	flag.Parse()

	if *helpF {
		flag.Usage()
		return
	}

	if *profileF {
		f, err := os.Create("cpu_profile.pprof")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var targetDir string
	if *devF {
		tmpDir, err := os.MkdirTemp("", filepath.Base(os.Args[0]))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer os.RemoveAll(tmpDir)
		targetDir = tmpDir
	} else {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		targetDir = filepath.Join(wd, TargetDirName)
	}

	GenerateWebsite(*devF, targetDir)
}
