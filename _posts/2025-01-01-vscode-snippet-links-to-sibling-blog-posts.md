---
permalink: "vscode-snippets-for-internal-blog-links"
title: "VS Code snippets for internal blog links"
summary: "Linking to other blog posts is only a few keys away"
date: "2025-01-01 23:33"
---

[Visual Studio Code supports user defined snippets](https://code.visualstudio.com/docs/editor/userdefinedsnippets)

It lets us spit out one snippet per blog post, by hooking into the generation of our blog. This is the end result:

<video controls>
  <source src="https://github.com/user-attachments/assets/02d53ef2-d89e-4afd-a720-9662e25ab230" type="video/mp4" />
</video>

## VS Code's JSON format

After some experimentation, this is the target format:

```json
{
  "A few system design \u0026 architecture experiences, low-level debugging stories and post-mortems": {
    "scope": "markdown",
    "prefix": "@@interesting-reads - A few system design \u0026 architecture experiences, low-level debugging stories and post-mortems",
    "body": "[${1:A few system design \u0026 architecture experiences, low-level debugging stories and post-mortems}](/interesting-reads) ",
    "description": "A few system design \u0026 architecture experiences, low-level debugging stories and post-mortems"
  },
  "A grep-quickie for the road (or svn, really)": {
    "scope": "markdown",
    "prefix": "@@a-grep-quickie-for-the-road-or-svn-really - A grep-quickie for the road (or svn, really)",
    "body": "[${1:A grep-quickie for the road (or svn, really)}](/a-grep-quickie-for-the-road-or-svn-really) ",
    "description": "A grep-quickie for the road (or svn, really)"
  },
  // ...
}
```

Good ideas:

- Use a common starting string, `@@`, for all snippets' `"prefix"`.
  - This should limit suggestions to only blog posts.
- Spam keywords in `"prefix"`.
  - This is the only part that VS Code's fuzzy auto completion matches against, `"description"` is decoration only.
- Provide a default for the link text (to the blog post title).
  - In body text, the link text will almost always be removed, but having a writing prompt is not too shabby.
  - VS Code smartly pre-selects the default, so there's no cost of just starting to write something completely different.
- Put a trailing space in `"body"`.
  - Allows for quickly continuing to write. If the link is at the end of a sentence, you need to follow up with a backspace. So be it.

## Generating the snippets

Thanks to VS Code's format, [all it takes is a loop to populate the snippets](https://github.com/chelmertz/chelmertz.github.io/commit/9bf5b41ccba352ed91450316b638d2f6a3bda6b9#diff-94984348a77639388f82ce61d30baf29dd5177f404e5a7a22dc3e2945d9ec5e2R360-R382).

By regenerating the snippets for each build of the blog contents, we make sure
that the snippets' URLs & titles are up to date.

## Postscript, on motivation

I'm trying to get over my private org mode file usage and "blog" more.  `wc`
reports 240k words, in comparison to only three blog posts between 2013 and
2024.

One part of this effort was to make the blog engine a bit more lightweight and
custom, [so I wrote a new one](/replacing-Jekyll-with-go).

Another part was to find a quick way to refer to sibling blog posts here, when
writing markdown.

Zettelkasten (small notes and a giant web of internal links) seems to be super
popular. There are many plugins supporting this, it sounds a bit like
[`gf`](https://vimdoc.sourceforge.net/htmldoc/editing.html#gf) in vim), but I
figured that each markdown post's `permalink` would be unique & stable enough to
use as a target.





