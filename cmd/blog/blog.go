package main

import (
	"io"
	"os"
	"path/filepath"
)

const docs = "docs"

// hardcoded, idempotent, blunt error handling, not incremental
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
