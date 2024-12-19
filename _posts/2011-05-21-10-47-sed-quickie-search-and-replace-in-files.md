---
date: "2011-05-21 10:47"
title: "sed quickie, search and replace in files"
permalink: "sed-quickie-search-replace-files"
tags: [bash]
---

There are many recipes for search and replacing via the command line but this one caught my attention because you can easily att more options to <strong>find</strong>.

    find containingFolder/ -type f -exec sed -i 's/oldText/newText/g' {} \;
