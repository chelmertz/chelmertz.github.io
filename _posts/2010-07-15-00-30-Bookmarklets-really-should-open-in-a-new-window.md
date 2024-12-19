---
date: "2010-07-15 00:30"
title: "Bookmarklets really should open in a new window"
permalink: "bookmarklet-really-should-be-in-a-new-window"
tags: [bookmarklet, javascript, thoughts, usability]
---

Until today (actually, tomorrow) I’ve made bookmarklets as alert/confirm-boxes with plain JS. Why on earth did I not use <code>window.open</code>?
<ul>
	<li>it will be updated upon each request, no need for clients to update the actual bookmark</li>
	<li>it’s styled</li>
	<li>it could be used with external JS-libraries</li>
</ul>
Doh… at least someone will be happy tomorrow at work :)

Inspiration for this post: the Delicious plugin for Firefox; but first and foremost, the tumblr-bookmarklet which is REALLY elegant.
