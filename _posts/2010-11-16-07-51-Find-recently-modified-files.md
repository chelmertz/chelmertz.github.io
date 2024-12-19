---
date: "2010-11-16 07:51"
title: "Find recently modified files"
permalink: "find-recently-modified-files"
tags: [bash]
---

    sudo find / -type f -ctime 0 | more

last 24 hours or change -ctime 0 to -cmin 5 for last five minutes
