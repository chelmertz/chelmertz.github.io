---
layout: post
title: "Properly display MySQL result in console"
permalink: "/properly-display-mysql-result-console"
tags: [mysql]
---

If your query results in a result set with many columns, you might end up with totally unreadable jibberish since the console defaults to print the columns horizontally. To display the data vertically instead, with support for really many columns, end the query with \G:

<pre><code lang="sql">select * from users \G</code></pre>

... instead of

<pre><code lang="sql">select * from users;</code></pre> 
