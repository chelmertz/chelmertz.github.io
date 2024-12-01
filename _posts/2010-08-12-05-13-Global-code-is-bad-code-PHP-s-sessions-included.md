---
layout: post
title: "Global code is bad code - PHP's sessions included"
permalink: "global-code-is-bad-code-phps-sessions-included"
tags: [php, thoughts]
---

If you rely on sessions in your application, just one call to <code>session_destroy()</code>, like that without any arguments, all your storage is reset. Doesn’t matter if it’s a shopping cart, logged in user’s surfing session, message alert, et cetera.

The only viable solution is to user namespaces like<code>$_SESSION['userSession']</code> (with <code>loggedInSince</code> as possible key), <code>$_SESSION['shoppingCart']</code> like such:
<ul>
	<li>product
<ul>
	<li>id =&gt; 4</li>
	<li>quantity =&gt; 5</li>
</ul>
</li>
</ul>
Thus, <strong><code>session_destroy()</code> becomes <code>unset($_SESSION['namespace'])</code>.</strong>
