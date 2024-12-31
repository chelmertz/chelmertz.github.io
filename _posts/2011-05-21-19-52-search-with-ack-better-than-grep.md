---
date: "2011-05-21 19:52"
title: "search with ack - \"better than grep\""
permalink: "search-ack-better-grep"
tags: [bash]
---

[ack](http://betterthangrep.com/) is a convenient alternative to searching with `grep`. It feels faster and has the stuff you use most often, searching recursively as a default, for example. And leaving out .svn-folders. For example:

```shell
grep -Ri "my_function" . | grep -v \.svn
```

becomes

```shell
ack -i "my_function"
```

Of course, after you install it, you'd want to add

```shell
alias "ack=ack-grep"
```

in your ~/.bash_aliases, to write even less.
