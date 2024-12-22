package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/url"
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
var csvTable = regexp.MustCompile(`<<csvtable:(.*?)>>`)

var viewFuncs = template.FuncMap{
	"rfc3339": func(t time.Time) string {
		return t.Format(time.RFC3339)
	},
}

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
		goldmark.WithExtensions(&frontmatter.Extender{}, extension.Linkify, extension.Table),
	)

	indexHtmlData := IndexHtmlData{}
	allPosts := make([]Post, 0)
	postTemplate := template.Must(template.New("post.html").
		Funcs(viewFuncs).
		ParseGlob("cmd/blog/post.html"))
	filepath.WalkDir("_posts", func(path string, d os.DirEntry, err error) error {
		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		markdownFile := must1(os.ReadFile(path))
		markdownContent := csvTableReplacer(markdownFile)

		var outputHtml bytes.Buffer
		ctx := parser.NewContext()
		must(md.Convert(markdownContent, &outputHtml, parser.WithContext(ctx)))
		postsFrontMatter := frontmatter.Get(ctx)

		if postsFrontMatter == nil {
			panic(fmt.Sprintf("no frontmatter found in %s\n", path))
		}

		var post Post
		must(postsFrontMatter.Decode(&post))
		post.Content = template.HTML(outputHtml.String())

		if post.Date == "" {
			panic(fmt.Sprintf("date not found in frontmatter for path %s\n", path))
		}

		// the frontmatter's "date" granularity is in minutes, i.e. "2024-12-31 23:59"
		publishedAt := must1(time.Parse(time.DateTime, post.Date+":00"))
		post.PublishedAt = publishedAt
		post.PublishedFriendly = publishedAt.Format("Jan 2, 2006")

		if post.Permalink == "" {
			panic(fmt.Sprintf("permalink not found in frontmatter for path %s\n", path))
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

	// docs/atom.xml
	atomTemplate := textTemplate.Must(textTemplate.New("atom.xml").
		Funcs(viewFuncs).
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
	Summary           string
}

// extension to allow markdown file to include a CSV file, formatted as a GFM table
func csvTableReplacer(rawMarkdown []byte) []byte {
	replaced := csvTable.ReplaceAllStringFunc(string(rawMarkdown), func(match string) string {
		matches := csvTable.FindStringSubmatch(match)
		if len(matches) > 1 {
			filename := "file://" + matches[1]
			options := csvOptions{}
			if fileUrl, err := url.Parse(filename); err == nil {
				filename = fileUrl.Hostname()

				query := fileUrl.Query()
				options.urlFrom = query.Get("urlFrom")
				options.urlTo = query.Get("urlTo")
			}
			gfmTable := gfmTableOfCsv(filename, options)
			return gfmTable
		}
		return match
	})

	return []byte(replaced)
}

type csvOptions struct {
	urlFrom, urlTo string
}

// Naive function that takes placeholder-with-filename -> opens csv -> parses
// csv In Good Faith™️ -> replaces the placeholder with a GFM (Github flavored
// markdown) table
//
// Will break if csv contains delimiters or otherwise gets parsed weirdly by
// encoding/csv
func gfmTableOfCsv(filename string, options csvOptions) string {
	reader := csv.NewReader(must1(os.Open(filepath.Join("tables", filename))))

	csvRows, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("could not read csv %s, got err %v", filename, err))
	}
	if len(csvRows) < 2 {
		// empty table, or just a header row
		return ""
	}

	// options
	fromIndex, toIndex := -1, -1
	doUrlTransformation := false
	if options.urlFrom != "" && options.urlTo != "" {
		for i, header := range csvRows[0] {
			if header == options.urlFrom {
				fromIndex = i
			}
			if header == options.urlTo {
				toIndex = i
			}
		}
		if fromIndex == -1 || toIndex == -1 {
			panic(fmt.Sprintf("could not find url headers in csv %s, with the options %v; headers: %v", filename, options, csvRows[0]))
		}
		doUrlTransformation = true

		csvRows[0] = append(csvRows[0][:fromIndex], csvRows[0][fromIndex+1:]...)
	}

	// header
	rows := make([]string, 0)
	row := fmt.Sprintf("| %s |", strings.Join(csvRows[0], " | "))
	rows = append(rows, row)

	headerUnderline := "| " + strings.Repeat(" --- |", len(csvRows[0]))
	rows = append(rows, headerUnderline)

	// columns
	for _, columns := range csvRows[1:] {
		if doUrlTransformation {
			if columns[fromIndex] != "" {
				// the URL column is not empty, make a link
				columns[toIndex] = fmt.Sprintf("[%s](%s)", columns[toIndex], columns[fromIndex])
			}
			columns = append(columns[:fromIndex], columns[fromIndex+1:]...)
		}
		rows = append(rows, fmt.Sprintf("| %s |", strings.Join(columns, " | ")))
	}

	return strings.Join(rows, "\n")
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
