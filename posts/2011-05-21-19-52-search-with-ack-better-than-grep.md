---
layout: post
title: "search with ack - \"better than grep\""
permalink: "/search-ack-better-grep"
tags: [bash]
---

<a title="ack" href="http://betterthangrep.com/">ack</a> is a convenient alternative to searching with <strong>grep</strong>. It feels faster and has the stuff you use most often, searching recursively as a default, for example. And leaving out .svn-folders. For example:

<pre><code lang=""bash"">grep -Ri "my_function" . | grep -v \.svn</code></pre>

becomes

<pre><code lang=""bash"">ack -i "my_function"</code></pre>

Of course, after you install it, you'd want to add

<pre><code lang=""bash"">alias "ack=ack-grep"</code></pre>

in your ~/.bash_aliases, to write even less.
