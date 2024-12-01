package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	textTemplate "text/template"
	"time"

	"github.com/yuin/goldmark"

	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

const docs = "docs"

var dateTimeInFilename = regexp.MustCompile(`\d{4}-\d{2}-\d{2}-\d{2}-\d{2}`)
var dateInFilename = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

// TODO just for fun, profile this after done
// hardcoded, idempotent, blunt error handling, not incremental, not parallel, duplicated data in memory
func main() {
	// github  pages wants to serve either / or docs/, let's go with docs/
	must(os.RemoveAll(docs))
	must(os.Mkdir(docs, 0775))

	// docs/assets
	must(os.Mkdir(filepath.Join(docs, "assets"), 0775))
	filepath.WalkDir("assets", func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			must(os.MkdirAll(filepath.Join(docs, path), 0775))
			return nil
		}
		src := must1(os.Open(path))
		dst := must1(os.Create(filepath.Join(docs, path)))
		_ = must1(io.Copy(dst, src))
		return nil
	})

	// parse posts
	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}, extension.Linkify),
	)

	indexHtmlData := IndexHtmlData{}
	allPosts := make([]Post, 0)
	postTemplate := template.Must(template.New("post.html").
		Funcs(template.FuncMap{
			"rfc3339": func(t time.Time) string {
				return t.Format(time.RFC3339)
			},
		}).
		ParseGlob("cmd/blog/post.html"))
	filepath.WalkDir("_posts", func(path string, d os.DirEntry, err error) error {
		if path == "_posts" {
			return nil
		}

		markdownFile := must1(os.ReadFile(path))
		var outputHtml bytes.Buffer
		ctx := parser.NewContext()
		must(md.Convert(markdownFile, &outputHtml, parser.WithContext(ctx)))
		postsFrontMatter := frontmatter.Get(ctx)

		if postsFrontMatter == nil {
			panic(fmt.Sprintf("no frontmatter found in %s\n", path))
		}

		var post Post
		must(postsFrontMatter.Decode(&post))
		post.Content = template.HTML(outputHtml.String())

		var publishedAt time.Time
		if post.Date == "" {
			// this is stupid, jekyll was bc-safe: date could be in the filename rather than frontmatter
			// this should be a one-time pass to migrate the old posts.. later?
			match := dateTimeInFilename.Find([]byte(d.Name()))
			if match == nil {
				panic(fmt.Sprintf("error parsing date from filename %s for path %s\n", d.Name(), path))
			}
			timeParts := bytes.Split(match, []byte{'-'})
			publishedAt = must1(time.Parse(time.DateTime, fmt.Sprintf("%s-%s-%s %s:%s:00", timeParts[0], timeParts[1], timeParts[2], timeParts[3], timeParts[4])))
		} else {
			// the frontmatter's "date" granularity is in minutes
			publishedAt, err = time.Parse(time.DateTime, post.Date+":00")
			if err != nil {
				panic(fmt.Sprintf("error parsing date from frontmatter for path %s: %v\n", path, err))
			}
		}
		post.PublishedAt = publishedAt
		post.PublishedFriendly = publishedAt.Format("Jan 2, 2006")

		if post.Permalink == "" {
			// variant 1: 2011-12-31-23-59-59-hello-world.md
			post.Permalink = dateTimeInFilename.ReplaceAllString(d.Name(), "")

			// variant 2: 2011-12-31-hello-world.md
			post.Permalink = dateInFilename.ReplaceAllString(d.Name(), "")

			// there's a leading hyphen, bridging the date and the title
			post.Permalink = strings.TrimPrefix(post.Permalink, "-")

			post.Permalink = strings.TrimSuffix(post.Permalink, ".md")
		} else {
			// rather than formatting permalink as a "slug", it's actually a relative path to a directory
			post.Permalink = strings.ReplaceAll(post.Permalink, "/", "")
		}

		allPosts = append(allPosts, post)

		indexHtmlData.Posts = append(indexHtmlData.Posts, IndexPost{
			Title:             post.Title,
			Permalink:         post.Permalink,
			PublishedFriendly: post.PublishedFriendly,
			timestamp:         int(publishedAt.Unix()),
		})

		must(os.Mkdir(filepath.Join(docs, post.Permalink), 0775))

		if _, err := os.Stat(filepath.Join(docs, post.Permalink, "index.html")); !errors.Is(err, os.ErrNotExist) {
			panic(fmt.Sprintf("post %s conflicts with some other post: %v\n", post.Permalink, err))
		}
		postIndexOutFile := must1(os.Create(filepath.Join(docs, post.Permalink, "index.html")))
		must(postTemplate.Execute(postIndexOutFile, post))

		return nil
	})

	// docs/index.html
	indexTemplate := template.Must(template.New("index.html").ParseGlob("cmd/blog/index.html"))
	indexOutFile := must1(os.Create(filepath.Join(docs, "index.html")))
	slices.SortFunc(indexHtmlData.Posts, func(i, j IndexPost) int {
		// sort by date, descending
		if i.timestamp > j.timestamp {
			return -1
		}
		if i.timestamp < j.timestamp {
			return 1
		}
		return 0
	})
	must(indexTemplate.Execute(indexOutFile, indexHtmlData))

	// atom feed: 10 entries
	// docs/atom.xml
	atomTemplate := textTemplate.Must(textTemplate.New("atom.xml").
		Funcs(textTemplate.FuncMap{
			"rfc3339": func(t time.Time) string {
				return t.Format(time.RFC3339)
			},
		}).
		ParseGlob("cmd/blog/atom.xml"))
	atomOutFile := must1(os.Create(filepath.Join(docs, "atom.xml")))
	slices.SortFunc(allPosts, func(i, j Post) int {
		// sort by date, descending
		asc := i.PublishedAt.Compare(j.PublishedAt)
		switch asc {
		case -1:
			return 1
		case 1:
			return -1
		}
		return 0
	})
	atomData := AtomData{
		FeedUpdated: allPosts[0].PublishedAt,
		Posts:       allPosts[:10],
	}
	must(atomTemplate.Execute(atomOutFile, atomData))

	// docs/CNAME
	must(os.WriteFile(filepath.Join(docs, "CNAME"), []byte("iamnearlythere.com"), 0664))
}

type IndexHtmlData struct {
	Posts []IndexPost
}

type IndexPost struct {
	Title             string
	Permalink         string
	PublishedFriendly string
	timestamp         int // for sorting reasons only
}

type AtomData struct {
	FeedUpdated time.Time
	Posts       []Post
}

type Post struct {
	Content           template.HTML
	Title             string
	PublishedAt       time.Time
	PublishedFriendly string
	Date              string
	Published         bool
	Permalink         string
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func must1[T any](something T, err error) T {
	must(err)
	return something
}
