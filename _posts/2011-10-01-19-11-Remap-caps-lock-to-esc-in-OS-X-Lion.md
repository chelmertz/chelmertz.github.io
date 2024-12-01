---
date: "2011-10-01 19:11"
layout: post
title: "Remap caps-lock to esc in OS X Lion"
permalink: "map-esc-caps-lock-os-lion"
tags: [mac]
---

The reason for removing caps-lock is mainly about being faster in vim at exiting entered modes, but also since I never use it and it's always in the way. Instead of just removing the key, we could make it usable though, and bind it to escape. The process is fairly easy.
<ol>
	<li>First you need to disable the key "controlling case sensitivity" (= caps lock). You find the option in <em>System Preferences &gt; Keyboard &gt; Modifier keys (in the first tab)</em>.</li>
	<li>Next, to bind caps-lock to the functionality of esc, install the <a href="http://pqrs.org/macosx/keyremap4macbook/extra.html">PCKeyboardHack</a></li>
	<li>Follow the instructions on the page in the link and use the keycode 53 for escape.</li>
	<li>You're done.</li>
</ol>
Via <a href="http://stackoverflow.com/questions/127591/using-caps-lock-as-esc-in-mac-os-x">Using Caps Lock as Esc in Mac OS X</a>
