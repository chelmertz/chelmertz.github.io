---
layout: post
title: "Fixing VLC when it doesn't want to play AVI-files"
permalink: "/fixing-vlc-when-it-doesnt-want-to-play-avi-files"
categories: [mac]
---

<a href="http://forum.videolan.org/viewtopic.php?f=12&amp;t=68748">http://forum.videolan.org/viewtopic.php?f=12&amp;t=68748</a>

Problem:
<blockquote class="posterous_short_quote">VLC cannot open files in the “AVI container” format</blockquote>
Solution:
<blockquote>
<ol>
	<li>Uninstall VLC</li>
	<li>Find and delete every file containing VLC in the name (“find ~ -iname *vlc*” in Terminal)</li>
	<li>Reinstall VLC</li>
</ol>
</blockquote>
