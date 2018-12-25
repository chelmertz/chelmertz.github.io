---
layout: post
title: "Always use -u with first git push"
permalink: "/use-u-with-first-git-push/"
tags: [git]
---

Sometimes (erhm? sorry for being vague) the default upstream branch is not being tracked by default when you `git clone` a repository.

When that happens, use the builtin `-`-switch:

    chelmertz@chelmertz ~/Sites/picture-this> git push -u origin gh-pages
    Branch gh-pages set up to track remote branch gh-pages from origin.
    Everything up-to-date
