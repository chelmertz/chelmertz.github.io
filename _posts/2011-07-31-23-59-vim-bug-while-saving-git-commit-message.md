---
date: "2011-07-31 23:59"
title: "vim \"bug\" while saving git commit message"
permalink: "vim-bug-while-saving-git-commit-message"
tags: [git, mac, vim]
---

Reminder for the issue `error: There was a problem with the editor 'vi'` while trying to save a commit message.

This can be solved by

```shell
git config --global core.editor `which vim`
```

via http://tooky.github.com/2010/04/08/there-was-a-problem-with-the-editor-vi-git-on-mac-os-x.html

Some recent changes that might have introduced this "bug" in my setup:

- Installing a brand new config: [spf13-vim](https://github.com/spf13/spf13-vim)
- Started to use gVim, more especially [MacVim](https://github.com/b4winckler/macvim)
