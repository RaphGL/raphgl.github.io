package main

import (
	"html/template"
	"os"
	"slices"
	"strings"
	"time"
)

type List struct {
	Posts []Post
}

func NewList(posts []Post) (List, error) {
	// sort in descending order
	var err error
	slices.SortFunc(posts, func(p1 Post, p2 Post) int {
		date1, err := time.Parse("2006-01-02", p1.Date)
		if err != nil {
			return 0
		}
		date2, err := time.Parse("2006-01-02", p2.Date)
		if err != nil {
			return 0
		}

		if date1.After(date2) {
			return -1
		}

		if date1.Before(date2) {
			return 1
		}

		return 0
	})

	if err != nil {
		return List{}, err
	}

	for _, post := range posts {
		_, destPath := GetCompiledTargetPath(post.SourceFilePath)
		post.SourceFilePath = destPath
	}

	return List{
		Posts: posts,
	}, nil
}

func (l List) Render() (template.HTML, error) {
	listTempl, err := os.ReadFile("./layout/list.html")
	if err != nil {
		return "", err
	}
	templ, err := template.New("list").Parse(string(listTempl))
	if err != nil {
		return "", err
	}

	var listBuilder strings.Builder
	if err = templ.Execute(&listBuilder, l); err != nil {
		return "", err
	}

	styles, err := GetGlobalStyles()
	if err != nil {
		return "", err
	}
	b := Base{
		BodyHTML:  template.HTML(listBuilder.String()),
		StylesCSS: styles,
	}

	page, err := b.Render()
	if err != nil {
		return "", err
	}
	return page, nil
}
