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

	listPosts := make([]Post, len(posts))
	for i := range len(posts) {
		post := posts[i]
		_, destPath := GetCompiledTargetPath(post.SourceFilePath)
		post.SourceFilePath = strings.TrimLeft(destPath, TargetDirName)
		listPosts[i] = post
	}

	return List{
		Posts: listPosts,
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

	b := Base{
		BodyHTML: template.HTML(listBuilder.String()),
	}

	page, err := b.Render()
	if err != nil {
		return "", err
	}
	return page, nil
}

func (l List) GenerateRSSFromPosts() string {
	posts := l.Posts

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
