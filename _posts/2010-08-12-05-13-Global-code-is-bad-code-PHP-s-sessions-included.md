---
date: "2010-08-12 05:13"
title: "Global code is bad code - PHP's sessions included"
permalink: "global-code-is-bad-code-phps-sessions-included"
tags: [php, thoughts]
---

If you rely on sessions in your application, just one call to `session_destroy()`, like that without any arguments, all your storage is reset. Doesn’t matter if it’s a shopping cart, logged in user’s surfing session, message alert, et cetera.

The only viable solution is to user namespaces like`$_SESSION['userSession']` (with `loggedInSince` as possible key), `$_SESSION['shoppingCart']` like such:

```
<ul>
	<li>product
	<ul>
		<li>id => 4</li>
		<li>quantity => 5</li>
	</ul>
</li>
</ul>

```

Thus, `session_destroy()` becomes `unset($_SESSION['namespace'])`.
