---
date: "2011-02-15 20:42"
title: "Automatically insert acronyms with Hz_Filter_Acronym"
permalink: "automatically-insert-acronyms-with-hz_filter_acronym"
tags: [php, project, zend-framework]
---

I hope this demo speaks for itself:

<script src="https://gist.github.com/828226.js?file=acronym_demo.php"></script>

This class is only good for parsing. If you’d like to use it in your application, here are some tips using the variable names from the gist’s example:

- initialize the class once, so you won’t have to pass on `$words` on every initiation
- cache the output, it’s an expensive round trip you do not want to do
- invalidate (delete) cache only when `$words` get bigger or when `$sentence` changes

Requirements:

- Zend Framework, or at least the `Zend_Filter_Interface`-class on your include path. Even this single requirement is optional, just comment out `implements Zend_Filter_Interface` if you want to use it as a standalone component.

You find [Hz_Filter_Acronym at github](https://github.com/chelmertz/Hz/blob/master/Hz/Filter/Acronym.php)
