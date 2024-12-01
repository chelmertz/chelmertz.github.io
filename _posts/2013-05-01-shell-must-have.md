---
permalink: "shell-must-have"
layout: post
title: "A shell's must-have features"
date: "2013-05-01 20:28"
tags: bash
---

I've been using zsh with oh-my-zsh for a long time now. If I were to replace it,
this is what I'd look for in the future:

  - Case insensitive completion/globbing
  - Auto completion of filenames via tab, left + right, enter
  - Fuzzy matching/spelling correction of targets when `cd`-ing
  - Have sensible prompt looks
    - (aliased) directory
    - git status
      - branch name
      - dirty-index indicator
    - no username
  - <kbd>ctrl+r</kbd> history
  - `!$` and `!!` expansion before execution
  - up + down to traverse history
  - globbing with `**`
