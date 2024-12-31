---
date: "2011-02-02 16:00"
title: "A grep-quickie for the road (or svn, really)"
permalink: "a-grep-quickie-for-the-road-or-svn-really"
tags: [bash, svn]
---

This post is outdated, you should read ["sed quickie, search and replace in files"](/sed-quickie-search-replace-files) instead!

-----

When working with svn instead of git, there are a lot more hits when performing a simple

```shell
grep -Ri "search for me" .
```

which, when using svn, returns in lots more hits than wanted. To filter them out, add another grep with the -v flag which inverts the results, i.e. passes along everything that does <em>not</em> match the following pattern. Like this:

```shell
grep -Ri "search for me" . | grep -v\.svn
```
