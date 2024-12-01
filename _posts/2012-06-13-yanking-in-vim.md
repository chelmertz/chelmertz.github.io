---
permalink: "yanking-in-vim"
layout: post
title: "Yanking in vim"
date: "2012-06-13 20:26"
tags: vim
---

A common pattern while editing is to yank anything (i.e. <kbd>yy</kbd>), visually select what you want to remove (i.e. <kbd>vj</kbd>) and paste (<kbd>p</kbd>). This will in turn place the overwritten test in the yank register; but we wanted to paste that same line again!

Type `:reg` to see all currently populated registers, they're all preceded by `"`. This means that you can paste by visually selecting and typing <kbd>"0p</kbd> if "0 is the register's content you want to paste.
