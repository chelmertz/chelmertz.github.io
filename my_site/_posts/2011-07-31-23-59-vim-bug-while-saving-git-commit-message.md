---
layout: post
title: "vim \"bug\" while saving git commit message"
permalink: "/vim-bug-while-saving-git-commit-message"
tags: [git, mac, vim]
---

Reminder for the issue <pre><code lang=""bash"">error: There was a problem with the editor 'vi'</code></pre> while trying to save a commit message.

This can be solved by
<pre><code lang=""bash"">$ git config --global core.editor `which vim`</code></pre> (via <a href="http://tooky.github.com/2010/04/08/there-was-a-problem-with-the-editor-vi-git-on-mac-os-x.html">http://tooky.github.com/2010/04/08/there-was-a-problem-with-the-editor-vi-git-on-mac-os-x.html</a>)

Some recent changes that might have introduced this "bug" in my setup:
<ul>
	<li>Installing a brand new config: <a href="https://github.com/spf13/spf13-vim">spf13-vim</a></li>
	<li>Started to use gVim, more especially <a href="https://github.com/b4winckler/macvim">MacVim</a></li>
</ul>
