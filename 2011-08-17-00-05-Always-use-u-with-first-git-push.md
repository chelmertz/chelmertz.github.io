---
layout: post
title: "Always use -u with first git push"
permalink: "/use-u-with-first-git-push"
categories: [git]
---

Sometimes (erhm? sorry for being vague) the default upstream branch is not being tracked by default when you <code lang=""bash"">git clone</code> a repository.

When that happens, use the builtin <code lang=""bash"">-u</code>-switch:

<pre><code lang=""bash"">chelmertz@chelmertz ~/Sites/picture-this> git push -u origin gh-pages
Branch gh-pages set up to track remote branch gh-pages from origin.
Everything up-to-date</code></pre> 
