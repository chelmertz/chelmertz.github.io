---
layout: post
title: "Retrieve removed file from SVN"
permalink: "/retrieve-removed-file-from-svn"
tags: [svn]
---

<code>svn copy -r &lt;revision&gt; http://&lt;repos&gt;/&lt;file&gt; .</code>

Pretty obvious really, but still… Use the remote path ( &amp; ) to the file that existed in  and place it in the current dir: . (or anywhere really).

PS: this retrieved the contents of  but you would also need to <code>svn commit</code> any changes to reinsert the file into the repos.
