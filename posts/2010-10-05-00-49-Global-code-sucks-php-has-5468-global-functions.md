---
layout: post
title: "Global code sucks - php has 5468 global functions"
permalink: "/global-code-sucks-php-has-5468-global-functions"
categories: [php, thoughts]
---

<em>Background: I’m thinking of studying for a Zend Certified Engineer (ZCE) exam, for php. A lot of it seems to understand and remember many different types of functions and their parameters (and their order).</em>

This sucks. <a href="http://gist.github.com/609898">I made a simple test</a>, realizing that there are well over 5000 globally defined php functions. Many of them may not be bundled, many of them will never appear in the ZCE test; but many of them affect the behavior of coders thinking that <code>hopefully_this_might_be_a_unique_namespace_my_function()</code> is a good way to organize code. Maybe, when you’ve reached around 4000 functions, maybe just then you begin to realize that some functions should be grouped in real classes and namespaces.

Who volunteer to help out in the migration process?
