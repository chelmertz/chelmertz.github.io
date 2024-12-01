---
date: "2010-06-03 07:17"
layout: post
title: "Read headers with curl"
permalink: "read-headers-with-curl"
tags: [server]
---

cURL is great for zooming in on a host, example:

<code>curl -I <a href="http://www.google.com">http://www.google.com</a></code>
<strong>HTTP/1.1 302 Found</strong>
<strong>Location: <a href="http://www.google.se/">http://www.google.se/</a></strong>

Cache-Control: private

Content-Type: text/html; charset=UTF-8

Set-Cookie: PREF=ID=4b82b41ad9278a4f:TM=1275516833:LM=1275516833:S=Qb2w71XyZoBvyONx; expires=Fri, 01-Jun-2012 22:13:53 GMT; path=/; domain=.google.com

Date: Wed, 02 Jun 2010 22:13:53 GMT

Server: gws

Content-Length: 218

X-XSS-Protection: 1; mode=block

<code>curl -I <a href="http://www.google.se">http://www.google.se</a></code>
<strong>HTTP/1.1 200 OK</strong>

Date: Wed, 02 Jun 2010 22:14:04 GMT

Expires: -1

Cache-Control: private, max-age=0

Content-Type: text/html; charset=ISO-8859-1

Set-Cookie: PREF=ID=c90ae10c41b80229:TM=1275516844:LM=1275516844:S=yKvAIaXwqmsaWD-H; expires=Fri, 01-Jun-2012 22:14:04 GMT; path=/; domain=.google.se

Server: gws

X-XSS-Protection: 1; mode=block

Transfer-Encoding: chunked

Nice!
