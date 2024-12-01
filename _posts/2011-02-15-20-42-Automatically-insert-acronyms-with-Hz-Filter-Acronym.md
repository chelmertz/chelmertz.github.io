---
date: "2011-02-15 20:42"
layout: post
title: "Automatically insert acronyms with Hz_Filter_Acronym"
permalink: "automatically-insert-acronyms-with-hz_filter_acronym"
tags: [php, project, zend-framework]
---

I hope this demo speaks for itself:

<script src="https://gist.github.com/828226.js?file=acronym_demo.php"></script>

This class is only good for parsing. If you’d like to use it in your application, here are some tips using the variable names from the gist’s example:
<ul>
	<li>initialize the class once, so you won’t have to pass on <code>$words</code> on every initiation</li>
	<li>cache the output, it’s an expensive round trip you do not want to do</li>
	<li>invalidate (delete) cache only when <code>$words</code> get bigger or when <code>$sentence</code> changes</li>
</ul>
Requirements:
<ul>
	<li>Zend Framework, or at least the <code>Zend_Filter_Interface</code>-class on your include path. Even this single requirement is optional, just comment out <code>implements Zend_Filter_Interface</code> if you want to use it as a standalone component.</li>
</ul>
You find <a href="https://github.com/chelmertz/Hz/blob/master/Hz/Filter/Acronym.php">Hz_Filter_Acronym at github</a>
