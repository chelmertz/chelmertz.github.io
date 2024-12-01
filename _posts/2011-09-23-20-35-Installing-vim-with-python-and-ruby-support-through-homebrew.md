---
layout: post
title: "Installing vim with python and ruby support through homebrew"
permalink: "installing-vim-python-ruby-support-homebrew"
tags: [mac, vim]
---

<h2>Background</h2>
My new macbook came builtin with an "old"(er) version of vim which apparently was compiled without python and ruby. I want that.

<h2>Prerequisites</h2>
Start by <a href="http://developer.apple.com/xcode/">installing xcode</a>, <a href="http://mxcl.github.com/homebrew/">homebrew</a> and git via <pre><code lang=""bash"">$ brew install git</code></pre>
<h2>Installing custom formula</h2>
Now, since I want vim compiled with python and ruby support (for some plugins), I can not use

    $ brew install vim

so I need a <a href="https://raw.github.com/Homebrew/homebrew-dupes/master/vim.rb">custom brew formula</a>:

    $ brew install https://raw.github.com/Homebrew/homebrew-dupes/master/vim.rb

<h2>Tell shell to use correctly compiled vim</h2>
(By default) homebrew installs its programs in /usr/local/bin. But after installing vim via brew in the previous step, I still get this:

    $ which vim
    /usr/bin/vim

My solution to using the correct vim, was to prepend /usr/local/bin to $PATH instead of having it somewhere in the middle. Now my $PATH looks like this in my .zshrc:

    export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/X11/bin:/opt/local/bin:/usr/local/git/bin

and, as evidence:

    $ which vim
    /usr/local/bin/vim
    $ vim --version | grep +ruby | echo $?
    0
    $ vim --version | grep +python | echo $?
    0

Allright, we're good to go!
