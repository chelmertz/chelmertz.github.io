---
date: "2010-05-07 02:13"
title: "Installing Vim = discovering Mac's internals"
permalink: "installing-vim-discovering-macs-internals"
tags: [mac, vim]
---

> Installation should be easy:  just drag Vim.app to where you want to keep it.  I put it in /Applications/ .
> 
> Unless you want a maximally vi-compatible version of vim, it is recommended that you copy the standard startup files to your home directory, and name them .vimrc and .gvimrc .  From a shell,
> 
> ```
> $ cd Vim.app/Contents/Resources/vim/runtime
> $ cp vimrc_example.vim ~/.vimrc
> $ cp gvimrc_example.vim ~/.gvimrc
> ```

So.. that means that everything in /Applications/ is some kind of archive/directory. Interesting.
