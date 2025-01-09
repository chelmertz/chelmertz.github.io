// "inspired" by https://git.icyphox.sh/vite/blob/master/atom/feed.go,
// https://pkg.go.dev/github.com/gorilla/feeds and some prompting
package main

import (
	"encoding/xml"
	"slices"
	"time"
)

type AtomLink struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr,omitempty"`
	Title   string   `xml:"title,attr,omitempty"`
}

type AtomSummary struct {
	XMLName xml.Name `xml:"summary"`
	Content string   `xml:",chardata"`
	Type    string   `xml:"type,attr,omitempty"`
}

type AtomContent struct {
	XMLName xml.Name `xml:"content"`
	Content string   `xml:",cdata"`
	Type    string   `xml:"type,attr,omitempty"`
}

type AtomAuthor struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

type AtomEntry struct {
	XMLName   xml.Name  `xml:"entry"`
	Title     string    `xml:"title"`
	Published string    `xml:"published"`
	Updated   string    `xml:"updated"`
	ID        string    `xml:"id"`
	Link      *AtomLink // TODO remove pointers?
	Summary   *AtomSummary
	Content   AtomContent
	Author    *AtomAuthor `xml:"author"`
}

type AtomFeed struct {
	XMLName  xml.Name `xml:"feed"`
	Xmlns    string   `xml:"xmlns,attr"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	ID       string   `xml:"id"`
	Updated  string   `xml:"updated"`
	Link     []AtomLink
	Author   *AtomAuthor `xml:"author"`
	Entries  []AtomEntry
}

func atomFeedOfPosts(posts []Post) ([]byte, error) {
	entries := []AtomEntry{}

	slices.SortFunc(posts, func(i, j Post) int {
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
	posts = posts[:10]

	feedUpdated := posts[0].PublishedAt

	author := &AtomAuthor{
		Name:  "Carl Helmertz",
		Email: "helmertz@gmail.com",
	}

	for _, post := range posts {
		publishedAt := post.PublishedAt.Format(time.RFC3339)
		permalink := "https://iamnearlythere.com/" + post.Permalink

		entry := AtomEntry{
			Title:     post.Title,
			Published: publishedAt,
			Updated:   publishedAt,
			ID:        permalink, // we assume published permalinks won't change
			Link: &AtomLink{
				Href:  permalink,
				Rel:   "alternate",
				Type:  "text/html",
				Title: post.Title,
			},
			Content: AtomContent{
				Type:    "html",
				Content: string(post.Content),
			},
			Author: author,
		}
		if post.Summary != "" {
			entry.Summary = &AtomSummary{
				Content: post.Summary,
			}
		}
		entries = append(entries, entry)
	}

	feed := &AtomFeed{
		Xmlns:    "http://www.w3.org/2005/Atom",
		ID:       "https://iamnearlythere.com/atom.xml",
		Title:    "iamnearlythere.com",
		Subtitle: "iamnearlythere.com, a blog on software development",
		Link: []AtomLink{
			{

				Href: "https://iamnearlythere.com/atom.xml",
				Rel:  "self",
				Type: "application/atom+xml",
			},
			{

				Href: "https://iamnearlythere.com/",
				Rel:  "alternate",
				Type: "text/html",
			},
		},
		Author:  author,
		Updated: feedUpdated.Format(time.RFC3339),
		Entries: entries,
	}

	feedXML, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return nil, err
	}

	return append([]byte(xml.Header), feedXML...), nil
}
