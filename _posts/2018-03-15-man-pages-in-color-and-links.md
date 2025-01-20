---
permalink: "man-pages-in-color-and-links"
title: "Man pages in color, with links"
date: "2018-03-15 00:24"
tags: bash
---

Update 2024-12-22:

I am using regular less now, with these lines in my zshrc:

```shell
export LESS_TERMCAP_mb=$'\e[1;32m'
export LESS_TERMCAP_md=$'\e[1;32m'
export LESS_TERMCAP_me=$'\e[0m'
export LESS_TERMCAP_se=$'\e[0m'
export LESS_TERMCAP_so=$'\e[01;33m'
export LESS_TERMCAP_ue=$'\e[0m'
export LESS_TERMCAP_us=$'\e[1;4;31m'
# the above worked in Ubuntu 22.04 but something broke "in" 24.04 and this workaround does it for me
# also see https://github.com/jedsoft/most/issues/18
export GROFF_NO_SGR=1
```

----

Old post:

[Neovim](https://neovim.io/) has builtin support for colorizing man pages, making them easier to read. It also lets you follow links to other man pages, by pressing <kbd>K</kbd> while hovering any term that has a man page.

To set up Neovim as the default reader for when you invoke `man apropos` etc., place the following code in your ~/.bashrc or such:

    export MANPAGER="nvim -c 'set ft=man' -"

There are some more pointers in the short and readable section of the [Neovim manual on viewing man pages](https://github.com/neovim/neovim/blob/5ce8158a5d462043306ee67a3261794f169bdb17/runtime/doc/filetype.txt#L509).
