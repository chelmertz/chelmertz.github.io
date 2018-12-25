---
layout: post
title: "Installing Vim = discovering Mac's internals"
permalink: "/installing-vim-discovering-macs-internals/"
tags: [mac, vim]
---

<blockquote>Installation should be easy:  just drag Vim.app to where you want to keep it.  I put it in /Applications/ .

Unless you want a maximally vi-compatible version of vim, it is recommended that you copy the standard startup files to your home directory, and name them .vimrc and .gvimrc .  From a shell,
<code> </code>
<pre><code>$ cd Vim.app/Contents/Resources/vim/runtime
$ cp vimrc_example.vim ~/.vimrc
$ cp gvimrc_example.vim ~/.gvimrc
</code></pre>
<code> </code>

&nbsp;</blockquote>
So.. that means that everything in /Applications/ is some kind of archive/directory. Interesting.
