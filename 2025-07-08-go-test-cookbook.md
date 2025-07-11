---
permalink: "go-test-patterns"
title: "Patterns for testing in Go"
summary: "Cookbook of patterns I've ran into"
date: "2025-07-08 13:05"
---

Go is 1.24.3 throughout this page.

## Executing all fuzz tests

`go test -fuzz=Fuzz ./...` bails with "cannot use -fuzz flag with multiple packages".

I use this workaround:

```shell
#!/bin/bash

# stolen from https://github.com/golang/go/issues/46312#issuecomment-1486038966

set -e

fuzzTime=${1:-10}

files=$(grep -r --include='**_test.go' --files-with-matches 'func Fuzz' .)

for file in ${files}
do
        funcs=$(grep -o 'func Fuzz\w*' $file | sed 's/func //')
        for func in ${funcs}
        do
                echo "Fuzzing $func in $file"
                parentDir=$(dirname $file)
                go test $parentDir -run=$func -fuzz=$func -fuzztime=${fuzzTime}s
        done
done
```

## Bypassing test cache

VSCode/Cursor sometimes thinks code has not changed, showing old results. Pass `-count 1` to `go test`.

## Performance profiling slow tests

## Bisecting test errors to find bugs

https://research.swtch.com/diffcover

## Table driven testing

Fuzzying probably helps finding more bugs, but if you're doing unit tests (or example-based tests), `t.Run()` is good to know about: https://go.dev/wiki/TableDrivenTests

In go.1.25, I'm looking forward to `T.Attr()` and `T.Output()` to further help differentiating tests in "tables": https://antonz.org/go-1-25/#test-attributes-and-friends

- t.Helper()
-  https://matttproud.com/blog/posts/debugging-go-table-tests.html


# Formatting test output

https://github.com/vakenbolt/go-test-report

# TODO

- https://www.youtube.com/watch?v=8hQG7QlcLBk
- golden testing (see elly + vcr & whatever we use at matchi)
- google/cmp
