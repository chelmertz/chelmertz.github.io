---
date: "2011-07-31 23:59"
layout: post
title: "vim \"bug\" while saving git commit message"
permalink: "vim-bug-while-saving-git-commit-message"
tags: [git, mac, vim]
---

Reminder for the issue `error: There was a problem with the editor 'vi'` while trying to save a commit message.

This can be solved by

    $ git config --global core.editor `which vim`

(via <a href="http://tooky.github.com/2010/04/08/there-was-a-problem-with-the-editor-vi-git-on-mac-os-x.html">http://tooky.github.com/2010/04/08/there-was-a-problem-with-the-editor-vi-git-on-mac-os-x.html</a>)

Some recent changes that might have introduced this "bug" in my setup:

- Installing a brand new config: <a href="https://github.com/spf13/spf13-vim">spf13-vim</a>
- Started to use gVim, more especially <a href="https://github.com/b4winckler/macvim">MacVim</a>
