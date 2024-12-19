---
date: "2011-09-19 21:12"
title: "Wrapping lines in git diff"
permalink: "wrapping-lines-git-diff"
tags: [git]
---

    $ git config --global core.pager 'less -r'

overrides the default option (less) by giving it the ability to wrap long lines in text files. That way you can see what's really different on two rows that begin in a similar fashion but differs at the end.
