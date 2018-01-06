---
layout: post
title: "A grep-quickie for the road (or svn, really)"
permalink: "/a-grep-quickie-for-the-road-or-svn-really/"
tags: [bash, svn]
---

This post is outdated, you should read <a title="use ack instead of grep" href="http://iamnearlythere.com/sed-quickie-search-replace-files">"sed quickie, search and replace in files"</a> instead!

<hr />

When working with <a href="http://iamnearlythere.com/tag/svn">svn</a> instead of <a href="http://iamnearlythere.com/tag/git">git</a>, there are a lot more hits when performing a simple

<code lang="bash">grep -Ri "search for me" .</code>

which, when using svn, returns in lots more hits than wanted. To filter them out, add another grep with the -v flag which inverts the results, i.e. passes along everything that does <em>not</em> match the following pattern. Like this:

<code lang="bash">grep -Ri "search for me" . | grep -v\.svn</code>
