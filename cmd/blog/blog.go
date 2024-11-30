package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/yuin/goldmark"

	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

const docs = "docs"

var dateTimeInFilename = regexp.MustCompile(`\d{4}-\d{2}-\d{2}-\d{2}-\d{2}`)
var dateInFilename = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

// TODO just for fun, profile this after done
// hardcoded, idempotent, blunt error handling, not incremental, not parallel
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
		goldmark.WithExtensions(&frontmatter.Extender{}),
	)

	indexHtmlData := IndexHtmlData{}
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

		var decodedFrontMatter PostMeta
		must(postsFrontMatter.Decode(&decodedFrontMatter))

		var publishedAt time.Time
		if decodedFrontMatter.Date == "" {
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
			publishedAt, err = time.Parse(time.DateTime, decodedFrontMatter.Date+":00")
			if err != nil {
				panic(fmt.Sprintf("error parsing date from frontmatter for path %s: %v\n", path, err))
			}
		}
		decodedFrontMatter.PublishedAt = publishedAt

		if decodedFrontMatter.Permalink == "" {
			// variant 1: 2011-12-31-23-59-59-hello-world.md
			decodedFrontMatter.Permalink = dateTimeInFilename.ReplaceAllString(d.Name(), "")

			// variant 2: 2011-12-31-hello-world.md
			decodedFrontMatter.Permalink = dateInFilename.ReplaceAllString(d.Name(), "")

			// there's a leading hyphen, bridging the date and the title
			decodedFrontMatter.Permalink = strings.TrimPrefix(decodedFrontMatter.Permalink, "-")

			decodedFrontMatter.Permalink = strings.TrimSuffix(decodedFrontMatter.Permalink, ".md")
		} else {
			decodedFrontMatter.Permalink = strings.ReplaceAll(decodedFrontMatter.Permalink, "/", "")
		}

		indexHtmlData.Posts = append(indexHtmlData.Posts, IndexPost{
			Title:       decodedFrontMatter.Title,
			Permalink:   decodedFrontMatter.Permalink,
			PublishedAt: publishedAt.Format("Jan 2, 2006"),
			timestamp:   int(publishedAt.Unix()),
		})

		must(os.Mkdir(filepath.Join(docs, decodedFrontMatter.Permalink), 0775))
		must(os.WriteFile(filepath.Join(docs, decodedFrontMatter.Permalink, "index.html"), outputHtml.Bytes(), 0664))

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
}

type IndexHtmlData struct {
	Posts []IndexPost
}
type IndexPost struct {
	Title       string
	Permalink   string
	PublishedAt string
	timestamp   int // for sorting reasons only
}

type PostMeta struct {
	Title       string
	PublishedAt time.Time
	Date        string
	Published   bool
	Permalink   string
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
