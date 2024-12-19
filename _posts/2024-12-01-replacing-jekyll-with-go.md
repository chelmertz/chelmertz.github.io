---
permalink: "replacing-Jekyll-with-go"
title: "Replacing Jekyll with Go"
summary: "Jekyll was too annoying and now wouldn't start - replacing it all with Go and a couple of libraries took a couple of hours"
date: "2024-12-01 11:58"
---

My old dockerized [Jekyll](https://jekyllrb.com/) instance started breaking, and
keeping the blog boring but built in a totally different way seemed like a good
way to spend some quality time.

The end result became [a Go
file](https://github.com/chelmertz/chelmertz.github.io/blob/2b54b9fd13b68cba9098929535eb878e1c974056/cmd/blog/blog.go)
that does things in a simple way - a saving clause from said file: "hardcoded,
idempotent, blunt error handling, not incremental, not parallel, duplicated data
in memory".

This is a small devlog:

- Github pages "just works" with Jekyll, when moving to non-Jekyll I:
  - "Deploy from the *main branch*"; maintaining separate branches is a personal
    no-no, for a small personal project, [trunk-based
    development](https://trunkbaseddevelopment.com/) is king.
  - "Deploy from a *folder*"; pre-generating & checking in things lets you know
    what you get ([just spin up a local HTTP
    server](https://github.com/chelmertz/serve)), compared to crossing your
    fingers everytime Github's Jekyll integration takes over.
    - I read the docs on [DNS
      configuration](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site)
      in the middle of the night, while making changes & waiting for TTLs to
      expire, but what finally worked: *have a CNAME file in the `docs/` folder,
      even if you're using A records*.
- [The new blog "script"](https://github.com/chelmertz/chelmertz.github.io/blob/2b54b9fd13b68cba9098929535eb878e1c974056/cmd/blog/blog.go) was written iteratively, fixing lots of papercuts along the way.
  - As an example: sometime around 2012, my version of Jekyll at the time,
    started to add a `date:` in the frontmatter, rather than having it as part
    of the filename. There were several such cutoffs (`permalink:` not always
    being there, ... etc.).
    - If I were developing Jekyll, I would instead use a migration strategy to
      bring all files up to date, and then never having to worry about the old
      format again (like.. every RDBMS migration library), but instead Jekyll
      seemed to have kept the complexity of "keep every old version working"
      within their tool. What Jekyll's codebase looks like with that kind of
      approach, I don't want to know.
  - Fixing one issue at a time, such as "parse publication date from filename
    _or_ frontmatter `date:`", let me:
    1. View a successfully rendered blog quicker.
    1. Amend the markdown posts with a more complete version of frontmatter
      properties.
    2. [Throw out the "reparation" code branches and replace them with a
      validation step](https://github.com/chelmertz/chelmertz.github.io/commit/9908552b27e5ad4b1005afaf514af59a994651d5).
- Go's backwards compatibility story will hopefully prove this right, in a
  couple of years from now.
- I am missing a single feature of jekyll: [extracting
  excerpts](https://github.com/jekyll/jekyll/blob/0e4182aefad27c72c6b1c0f0e300e57edefaa0ba/lib/jekyll/excerpt.rb#L98-L145),
  that are put in the generated HTML's `<meta name="description"/>` and the items
  of the Atom feed's `<summary>`.
  - I just put this as a manual thing to write as an optional `summary:` front
    matter property, which would probably create a much better result in the
    end.

Here are some sed lines for working with frontmatter, notably the "only work
within the delimiter lines", which was news to me.

**Removing slashes in `permalink:`**

```bash
sed -i -e '/^---$/,/^---$/ {/^\s*permalink: /s/\///g}' _posts/*md
```

**Removing empty lines in frontmatter:**

```bash
sed -i '/^---$/,/^---$/ {/^\s*$/d}' _posts/*md
```

**Removing a frontmatter property:**

```bash
sed -i '/^---$/,/^---$/ {/^\s*comments:/d}' _posts/*md
```

And here's a paper trail of what broke the camel's back:

```
...
Error reading file /srv/Jekyll/_layouts/post.html: No such file or directory @ rb_sysopen - /srv/jekyll/my_site/srv/jekyll/_layouts/post.html 
Error reading file /srv/Jekyll/_layouts/default.html: No such file or directory @ rb_sysopen - /srv/jekyll/my_site/srv/jekyll/_layouts/default.html 
Error reading file /srv/Jekyll/_layouts/page.html: No such file or directory @ rb_sysopen - /srv/jekyll/my_site/srv/jekyll/_layouts/page.html 
Error reading file /srv/Jekyll/_layouts/home.html: No such file or directory @ rb_sysopen - /srv/jekyll/my_site/srv/jekyll/_layouts/home.html
```
