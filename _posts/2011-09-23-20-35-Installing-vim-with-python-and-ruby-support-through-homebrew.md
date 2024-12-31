---
date: "2011-09-23 20:35"
title: "Installing vim with python and ruby support through homebrew"
permalink: "installing-vim-python-ruby-support-homebrew"
tags: [mac, vim]
---

My new macbook came builtin with an "old"(er) version of vim which apparently was compiled without python and ruby. I want that.

## Prerequisites

Start by [installing xcode](http://developer.apple.com/xcode/), [homebrew](http://mxcl.github.com/homebrew/) and git via `$ brew install git`

## Installing custom formula

Now, since I want vim compiled with python and ruby support (for some plugins), I can not use

```shell
brew install vim
```

so I need a [custom brew formula](https://raw.github.com/Homebrew/homebrew-dupes/master/vim.rb):

```shell
brew install https://raw.github.com/Homebrew/homebrew-dupes/master/vim.rb
```

## Tell shell to use correctly compiled vim

(By default) homebrew installs its programs in /usr/local/bin. But after installing vim via brew in the previous step, I still get this:

```shell
which vim
/usr/bin/vim
```

My solution to using the correct vim, was to prepend /usr/local/bin to $PATH instead of having it somewhere in the middle. Now my $PATH looks like this in my .zshrc:

```shell
export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/X11/bin:/opt/local/bin:/usr/local/git/bin
```

and, as evidence:

```shell
which vim
/usr/local/bin/vim
vim --version | grep +ruby | echo $?
0
vim --version | grep +python | echo $?
0
```

Allright, we're good to go!
