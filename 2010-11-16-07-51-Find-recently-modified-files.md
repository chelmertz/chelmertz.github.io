---
layout: post
title: "Find recently modified files"
permalink: "/untitled-3"
categories: [bash]
---

    sudo find / -type f -ctime 0 | more

last 24 hours or change -ctime 0 to -cmin 5 for last five minutes
