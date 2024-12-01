---
permalink: "replacing-jekyll-with-200-lines-of-go"
layout: post
title: "Replacing jekyll with 200 lines of go"
date: "2024-12-01 11:58"
---

triggered by some breaking changes in my old dockerized jekyll instance (folder mismatch.. something)

- dns (docs/CNAME required even though I have A records)
- all sed things for dealing with frontmatter
- jekyll backwards compatible strategy: keep accepting old input instead of migrating it
- checking in generated code is nice, let's me play with the generator and see the output
- no separate branches etc, but still a "www root" (in docs/) keeps everything quite simple
- go's bc story is friendlier
- missing: automatic meta-description/summary in atom feed
