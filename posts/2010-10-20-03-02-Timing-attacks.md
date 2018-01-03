---
layout: post
title: "Timing attacks"
permalink: "/timing-attacks"
tags: [security]
---

<a href="http://devzone.zend.com/article/12654-Zend-Framework-1.11.0BETA1-Released">http://devzone.zend.com/article/12654-Zend-Framework-1.11.0BETA1-Released</a>
<blockquote class="posterous_medium_quote">The nature of the leaks is that strings are often compared byte by byte, with a negative result being returned early as soon as any set of non-matching bytes is detected. The more bytes that are equal (starting from the first byte) between both sides of the comparison, the longer it takes for a final result to be returned. Based on the time it takes to return a negative or positive result, it is possible that an attacker could, over many samples of requests, craft a string that compares positively to another secret string value known only to a target server simply by guessing the string one byte at a time and measuring each guess’ execution time.</blockquote>
Perfect scope for a problem that is easily adopted on framework level in an application. It’s probably a very unusual type of attack to defend against, building the ordinary out-of-the-box-application.

<strong>I wonder why not all companies producing serious code, hire hackers for a lot of $$$ to find bugs in their software.</strong>
