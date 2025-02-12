package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var postTemplateText = `---
permalink: "todo-should-be-a-lowercase-slug-here"
title: "TODO fill this in"
summary: "TODO fill this in"
date: "{{.Now}}"
---

this is some content
`

func randomSuffix() string {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func main() {
	filedate := time.Now().Format("2006-01-02")

	path := filepath.Join("_posts", filedate+"-todo-"+randomSuffix()+".md")
	outfile, createErr := os.Create(path)
	if createErr != nil {
		panic(fmt.Sprintf("error creating temp post file: %v", createErr))
	}

	postTemplate := template.Must(template.New("post").Parse(postTemplateText))
	templateErr := postTemplate.Execute(outfile, map[string]string{
		"Now": time.Now().Format("2006-01-02 15:04"),
	})
	if templateErr != nil {
		panic(fmt.Sprintf("error executing template: %v", templateErr))
	}
	fmt.Printf("created %s\n", path)
}
