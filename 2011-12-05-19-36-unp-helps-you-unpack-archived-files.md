---
layout: post
title: "unp - helps you unpack archived files"
permalink: "/unp-helps-unpack-archived-files"
categories: [app]
---

I used to always forget how to unpack files properly, with <code lang=""bash"">tar -zxf</code> for example. Still, it's a hassle when you manage tar, tar.gz, zip and rar files with their separate commands for unarchiving.

unp is, like <a href="http://iamnearlythere.com/search-ack-better-grep">ack</a>, a utility written in perl which unifies your ways of unpacking archives:

<pre><code lang=""bash"">unp archive.zip
unp package.rpm
unp -U *.rar
unp -U *</code></pre>

As usual, the easiest way to get hold of this package on OSX is by using homebrew: <code lang=""bash"">brew install unp</code>

Via <a href="http://glesys.se">glesys.se</a>
