---
permalink: "man-pages-in-color-and-links"
layout: post
title: "Man pages in color, with links"
date: "2018-03-15 00:24"
tags: bash
---

[Neovim](https://neovim.io/) has builtin support for colorizing man pages, making them easier to read. It also lets you follow links to other man pages, by pressing <kbd>K</kbd> while hovering any term that has a man page.

To set up Neovim as the default reader for when you invoke `man apropos` etc., place the following code in your ~/.bashrc or such:

    export MANPAGER="nvim -c 'set ft=man' -"

There are some more pointers in the short and readable section of the [Neovim manual on viewing man pages](https://github.com/neovim/neovim/blob/5ce8158a5d462043306ee67a3261794f169bdb17/runtime/doc/filetype.txt#L509).
