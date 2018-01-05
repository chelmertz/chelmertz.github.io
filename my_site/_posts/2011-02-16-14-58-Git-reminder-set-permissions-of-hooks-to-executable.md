---
layout: post
title: "Git reminder: set permissions of hooks to executable"
permalink: "/git-reminder-set-permissions-of-hooks-to-executable"
tags: [bash, git]
---

When you’ve written a git hook, don’t forget to make it executable. Git will not tell you why it’s skipped and go ahead and skip the hook all together.

Solution:

    chmod ug+x .git/hooks/*
