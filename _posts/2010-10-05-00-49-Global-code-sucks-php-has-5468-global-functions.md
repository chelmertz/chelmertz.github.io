---
date: "2010-10-05 00:49"
title: "Global code sucks - php has 5468 global functions"
permalink: "global-code-sucks-php-has-5468-global-functions"
tags: [php, thoughts]
---

Background: I’m thinking of studying for a Zend Certified Engineer (ZCE) exam, for php. A lot of it seems to understand and remember many different types of functions and their parameters (and their order).

This sucks. [I made a simple test](http://gist.github.com/609898), realizing that there are well over 5000 globally defined php functions. Many of them may not be bundled, many of them will never appear in the ZCE test; but many of them affect the behavior of coders thinking that `hopefully_this_might_be_a_unique_namespace_my_function()` is a good way to organize code. Maybe, when you’ve reached around 4000 functions, maybe just then you begin to realize that some functions should be grouped in real classes and namespaces.

Who volunteer to help out in the migration process?
