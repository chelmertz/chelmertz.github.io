---
layout: post
title: "Wrapping lines in git diff"
permalink: "/wrapping-lines-git-diff"
tags: [git]
---

<pre><code lang=""bash"">$ git config --global core.pager 'less -r'</code></pre> overrides the default option (less) by giving it the ability to wrap long lines in text files. That way you can see what's really different on two rows that begin in a similar fashion but differs at the end.
