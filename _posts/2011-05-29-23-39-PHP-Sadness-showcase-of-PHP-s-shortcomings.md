---
date: "2011-05-29 23:39"
layout: post
title: "PHP Sadness - showcase of PHP's shortcomings"
permalink: "php-sadness-showcase-phps-shortcomings"
tags: [php, thoughts]
---

<a href="http://phpsadness.com/">PHP Sadness</a> is a listing of PHP's roughest edges - things that could have been clearer, weird implementations and plain errors in design. I agree with most of the list's items, especially the ones concerning inconsistent behavior/signatures of native functions. Most of those annoyances would be impossible to get wrong if everything was an object and functions were instead methods (yes, I'm jealously looking at Ruby). For example:

    // what's correct again?
    in_array('blue', $array);
    // or, perhaps:
    in_array($array, 'blue');

    //wouldn't this be better?
    $array->contains('blue')
    // or perhaps
    'blue'->in_array($array)

The second most important part is the often really badly expressed error messages, sometimes even the stack trace is corrupted and then the developer's guessing game begins.

I really hope that some of the issues gets fixed since the article has been by wide audiences already, perhaps it will put some pressure on fixing what needs to be fixed.
