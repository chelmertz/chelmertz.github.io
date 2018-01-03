---
layout: post
title: "site2pdf - save all pages from a domain as a single PDF"
permalink: "/site2pdf-save-pages-domain-single-pdf"
tags: [project, ruby]
---

Point <strong>site2pdf</strong> to a domain you want and it will save all of its pages to a single PDF:

<pre><code lang=""bash"">ruby site2pdf http://12factor.net</code></pre>

This detects all "internal links" (links to pages on the same domain), crawls these and joins them together. The target for my development process was <a href="http://12factor.net">12factor.net</a> which is well written checklist and guide for your development and deployment.

There are a couple of dependencies: you need rubygems and wkhtmltopdf (compiled with qt in order to get all URI:s into the same PDF document).

<strong>Download site2pdf on github: <a title="site2pdf - save all pages from a domain as a single PDF" href="https://github.com/chelmertz/site2pdf">https://github.com/chelmertz/site2pdf</a> </strong>(code is subject for a lot of refactoring when I'm completely out of other stuff to do).
