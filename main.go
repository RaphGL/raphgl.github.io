package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
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

// converts the path to its equivalent in the `TargetDirName` directory
func GetTargetPath(path string) (destPath string) {
	pathComponents := strings.Split(path, string(filepath.Separator))[1:]
	destComponents := slices.Insert(pathComponents, 0, TargetDirName)
	return strings.Join(destComponents, string(filepath.Separator))
}

// returns the path of the compiled file without the `TargetDirName`
// this is useful for URLs.
//
// Note: an error will be returned if the file does not have a `.md` extension
func GetCompiledTargetPath(path string) (string, error) {
	target := strings.Join(strings.Split(path, string(filepath.Separator))[1:], string(filepath.Separator))
	ext := filepath.Ext(target)
	if ext != ".md" {
		return "", fmt.Errorf("Expected file with a `.md` extension but found `%s`", ext)
	}
	target = string(filepath.Separator) + strings.Replace(target, ext, ".html", 1)
	return target, nil
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

// writes post to `destPath` and returns the post for further inspection if needed
func GeneratePost(isDevMode bool, destPath, postPath string) (Post, error) {
	post, err := NewPost(postPath)
	if err != nil {
		return Post{}, err
	}
	if post.Draft {
		return Post{}, nil
	}
	// Post Hooks
	{
		if !isDevMode {
			post.AddCheckerHook(CheckLinkIsReachable)
		}
	}

	htmlArtifact, err := post.Render()
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return Post{}, err
	}

	if err := os.WriteFile(destPath, []byte(htmlArtifact), 0644); err != nil {
		return Post{}, err
	}

	return post, nil
}

// writes the entire website to `targetDirPath`
func GenerateWebsite(targetDirPath string) {
	// we remove all files in target dir first to prevent previous
	// compilation items from being left in the final website artifacts
	if err := os.RemoveAll(targetDirPath); err != nil {
		fmt.Println(err)
		return
	}
	if err := os.MkdirAll(targetDirPath, 0755); err != nil {
		fmt.Println(err)
		return
	}

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

	var wg sync.WaitGroup
	var renderedPosts []Post
	/* Generate posts */ {
		postChan := make(chan Post, 16)
		for _, filePath := range mdFiles {
			wg.Go(func() {
				destPath, err := GetCompiledTargetPath(filePath)
				if err != nil {
					fmt.Println(err)
					return
				}
				destPath = filepath.Join(targetDirPath, destPath)
				post, err := GeneratePost(false, destPath, filePath)
				if err != nil {
					fmt.Println(err)
					return
				}

				postChan <- post
			})
		}
		for range len(mdFiles) {
			renderedPosts = append(renderedPosts, <-postChan)
		}
		wg.Wait()
		close(postChan)
	}

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
		// Note: Copying files take time, might as well try and make it concurrent
		wg.Go(func() {
			destPath := GetTargetPath(file)
			if err := CopyStaticFile(destPath, file); err != nil {
				fmt.Println(err)
			}
		})
	}
	defer wg.Wait()

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

func ServeWebsiteWithHotReload() error {
	tmpDir, err := os.MkdirTemp("", filepath.Base(os.Args[0]))
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	docsDir := filepath.Join(tmpDir, TargetDirName)

	go func() {
		cssTarget := filepath.Join(docsDir, "styles.css")

		for {
			if len(DetectDirFilesChanged("./layout", ".html")) != 0 {
				GenerateWebsite(docsDir)
				continue
			}

			for _, f := range DetectDirFilesChanged("./content", ".md") {
				destPath, err := GetCompiledTargetPath(f)
				if err != nil {
					fmt.Println(err)
					continue
				}
				destPath = filepath.Join(docsDir, destPath)
				_, err = GeneratePost(true, destPath, f)
				if err != nil {
					fmt.Println(err)
				}
			}

			if len(DetectDirFilesChanged("./layout", ".css")) != 0 {
				css, err := GetStyles()
				if err != nil {
					fmt.Println(err)
				}

				if err := os.WriteFile(cssTarget, []byte(css), 0644); err != nil {
					fmt.Println(err)
				}
			}

			time.Sleep(400 * time.Millisecond)
		}
	}()

	fmt.Println("Using ", tmpDir, "as temporary directory")
	fmt.Println("Starting server at http://localhost:8080")

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/", http.FileServer(http.Dir(docsDir)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	return nil
}

func main() {
	devF := flag.Bool("dev", false, "Enable development mode")
	profileF := flag.Bool("profile", false, "Generate program profile data")
	helpF := flag.Bool("help", false, "Show help message")
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

	if *devF {
		err := ServeWebsiteWithHotReload()
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		targetDir := filepath.Join(wd, TargetDirName)
		GenerateWebsite(targetDir)
	}
}
