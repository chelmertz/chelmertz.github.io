---
date: "2011-11-28 22:10"
title: "site2pdf - save all pages from a domain as a single PDF"
permalink: "site2pdf-save-pages-domain-single-pdf"
tags: [project, ruby]
---

Point _site2pdf_ to a domain you want and it will save all of its pages to a single PDF:

    ruby site2pdf http://12factor.net

This detects all "internal links" (links to pages on the same domain), crawls these and joins them together. The target for my development process was [12factor.net](http://12factor.net) which is well written checklist and guide for your development and deployment.

There are a couple of dependencies: you need rubygems and wkhtmltopdf (compiled with qt in order to get all URI:s into the same PDF document).

Download site2pdf on github: https://github.com/chelmertz/site2pdf

(code is subject to a lot of refactoring when I'm completely out of other stuff to do).
