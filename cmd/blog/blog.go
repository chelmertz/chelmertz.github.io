package main

import (
	"bytes"
	"cmp"
	"encoding/csv"
	"encoding/json"
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

	tfidf "github.com/dkgv/go-tf-idf"

	"github.com/yuin/goldmark"

	highlighting "github.com/yuin/goldmark-highlighting/v2"
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
		goldmark.WithExtensions(&frontmatter.Extender{}, extension.Linkify, extension.Table, highlighting.Highlighting),
	)

	indexHtmlData := IndexHtmlData{}
	indexHtmlData.PostsByYear = make(map[int][]IndexPost)

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
		post.RawMarkdown = markdownContent

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

		year := publishedAt.Year()
		indexHtmlData.PostsByYear[year] = append(indexHtmlData.PostsByYear[year], IndexPost{
			Title:             post.Title,
			Permalink:         post.Permalink,
			PublishedFriendly: post.PublishedFriendly,
			timestamp:         int(publishedAt.Unix()),
		})

		return nil
	})

	relatedPosts(allPosts)
	createSnippets(allPosts)

	{
		// docs/posts
		for _, post := range allPosts {
			must(os.Mkdir(filepath.Join(docs, post.Permalink), 0775))

			if _, err := os.Stat(filepath.Join(docs, post.Permalink, "index.html")); !errors.Is(err, os.ErrNotExist) {
				panic(fmt.Sprintf("post %s conflicts with some other post: %v\n", post.Permalink, err))
			}
			postIndexOutFile := must1(os.Create(filepath.Join(docs, post.Permalink, "index.html")))
			must(postTemplate.Execute(postIndexOutFile, post))
		}
	}

	{
		// docs/index.html
		indexTemplate := template.Must(template.New("index.html").ParseGlob("cmd/blog/index.html"))
		indexOutFile := must1(os.Create(filepath.Join(docs, "index.html")))
		indexHtmlData.Years = make([]int, 0)
		for ip := range indexHtmlData.PostsByYear {
			slices.SortFunc(indexHtmlData.PostsByYear[ip], func(i, j IndexPost) int {
				// sort by date, descending
				if i.timestamp > j.timestamp {
					return -1
				}
				if i.timestamp < j.timestamp {
					return 1
				}
				return 0
			})
			indexHtmlData.Years = append(indexHtmlData.Years, ip)
		}
		slices.SortFunc(indexHtmlData.Years, func(i, j int) int {
			return cmp.Compare(j, i)
		})

		must(indexTemplate.Execute(indexOutFile, indexHtmlData))
	}

	{
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
	}

	// docs/CNAME
	must(os.WriteFile(filepath.Join(docs, "CNAME"), []byte("iamnearlythere.com"), 0664))
}

type IndexHtmlData struct {
	PostsByYear map[int][]IndexPost
	Years       []int
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
	RawMarkdown       []byte
	Title             string
	PublishedAt       time.Time
	PublishedFriendly string
	Date              string
	Published         bool
	Permalink         string
	Summary           string
	RelatedPosts      []RelatedPost
}

type RelatedPost struct {
	Permalink        string
	Title            string
	CosineSimilarity float64
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

// relatedPosts modifies all posts to include related posts.
func relatedPosts(posts []Post) {
	rawMarkdowns := make([]string, len(posts))
	for i := range posts {
		rawMarkdowns[i] = string(posts[i].RawMarkdown)
	}
	comparator := tfidf.New(tfidf.WithDefaultStopWords(),
		tfidf.WithComparator(tfidf.CosineComparator),
		tfidf.WithDocuments(rawMarkdowns),
	)

	for i := range posts {
		similarities := make([]RelatedPost, 0)
		for j := range posts {
			if posts[i].Permalink == posts[j].Permalink {
				continue
			}
			similarity := must1(comparator.Compare(string(posts[i].RawMarkdown), string(posts[j].RawMarkdown)))
			similarities = append(similarities, RelatedPost{
				Permalink:        "/" + posts[j].Permalink,
				Title:            posts[j].Title,
				CosineSimilarity: similarity,
			})
		}

		slices.SortFunc(similarities, func(i, j RelatedPost) int {
			return cmp.Compare(j.CosineSimilarity, i.CosineSimilarity)
		})

		posts[i].RelatedPosts = make([]RelatedPost, 5)
		for sim := range 5 {
			posts[i].RelatedPosts[sim] = similarities[sim]
		}
	}
}

type Snippet struct {
	Scope       string `json:"scope"`
	Prefix      string `json:"prefix"`
	Body        string `json:"body"`
	Description string `json:"description"`
}

func createSnippets(posts []Post) {
	snippets := make(map[string]Snippet)

	for _, post := range posts {
		snippets[post.Title] = Snippet{
			Scope:       "markdown",
			Prefix:      fmt.Sprintf("@@%s - %s", post.Permalink, post.Title),
			Body:        fmt.Sprintf("[${1:%s}](/%s) ", post.Title, post.Permalink),
			Description: post.Title,
		}
	}

	asJson, jsonErr := json.MarshalIndent(snippets, "", "  ")
	if jsonErr != nil {
		panic(fmt.Sprintf("error marshalling vscode snippets for : %v", jsonErr))
	}
	snippetsFile := filepath.Join(".vscode", "posts.code-snippets")
	writeErr := os.WriteFile(snippetsFile, asJson, 0644)

	if writeErr != nil {
		panic(fmt.Sprintf("error writing vscode snippets to file %s : %v", snippetsFile, writeErr))
	}
}

/*
This is what vscode generates in the snippets boilerplate:

{
	// Place your chelmertz.github.io workspace snippets here. Each snippet is defined under a snippet name and has a scope, prefix, body and
	// description. Add comma separated ids of the languages where the snippet is applicable in the scope field. If scope
	// is left empty or omitted, the snippet gets applied to all languages. The prefix is what is
	// used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
	// $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders.
	// Placeholders with the same ids are connected.
	// Example:
	// "Print to console": {
	// 	"scope": "javascript,typescript",
	// 	"prefix": "log",
	// 	"body": [
	// 		"console.log('$1');",
	// 		"$2"
	// 	],
	// 	"description": "Log output to console"
	// }
}
*/

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func must1[T any](something T, err error) T {
	must(err)
	return something
}
