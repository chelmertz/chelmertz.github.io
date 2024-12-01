---
permalink: "elly-keeps-track-of-prs"
layout: post
title: "elly - PRs you can act on"
date: "2023-08-02 16:34"

tags: project, go
comments: true
---

[elly](https://github.com/chelmertz/elly) shows you all of the pull requests you
can take action on, by periodically pulling data from Github.

**You should deal with the PRs in order, they're ranked by actionability.**

For example, if a PR gets approved, _elly_ shows the PR in the top - just merge it.

If there's a new comment, _elly_ bumps that PR up.

If you just responded to someone's comment, but kept the comment thread open instead of resolving it, _elly_ knows you're waiting for a reply and places the PR lower in the list.

This is what it looks like:

![gui](/assets/elly_gui.png)

![about](/assets/elly_about.png)

## Installation instructions

You'll find up to date [installation instructions in the project's README](https://github.com/chelmertz/elly).

## System design

![architecture](/assets/elly_architecture.png)

## A small retrospective

- [Go fuzz testing](https://go.dev/security/fuzz/)
  - A [single test found 3 bugs](https://github.com/chelmertz/elly/commit/4bd771bb32ded27cd048d168034d860ae2bf77ba)
  - Inspired by re-reading [Dan Luu on testing](https://danluu.com/testing/) for
    the fiftieth time. Finally something that unblocked me, after having read
    countless "reversing things twice" or "encoding/decoding is commutative"
    example. Writing longer test method bodies are fine.
  - Going from one-off unit tests, to [table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests),
    to fuzzying, seems like a path I want to go down.
- Writing [ADR](https://adr.github.io/)s (architectural decision record) helped
  even for such a small project (see the `/decisions` folder)
  - Feels nice to not having to go back on certain decisions, since they're
    spelled out and reasoned about. For example: KISS, with a JSON file rather
    than an SQLite DB, felt better after articulating it in text.
  - Went back and edited them a lot. They should probably have the _draft_ or
    _suggestion_ status for a while.
  - The relation between retrospective items and ADRs almost feel 1:1. Using
    these methods, or something else, that triggers thinking about design
    before/during/after implementation is gold.
- [GraphQL](https://graphql.org/)
  - First time querying against it. Using headers for auth, getting proper
    response codes back, and having the query in plain text is pleasant.
  - Github specific:
    - Good [API explorer](https://docs.github.com/en/graphql/overview/explorer),
      especially with the search box. I guess that the explorer is standard,
      through some framework.
    - Trying to extract enough comments (below pull requests, below repositories,
      below search results) is weird - should I traverse that paginated
      sub-sub-sub resource? The "a single query that reaches everywhere"
      paradigm breaks down a bit, to me.
- Front end
  - Making a [footer really sticky](https://stackoverflow.com/questions/4575826/how-to-push-a-footer-to-the-bottom-of-page-when-content-is-short-or-missing)
  - [`<dialog/>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog).
    Amazing that this now exists. No deps/configuration needed, and `::backdrop`
    is a perfect partner.
  - [Chrome dev tool's flexbox editor](https://developer.chrome.com/docs/devtools/css/flexbox/)
    - Never noticed this before, it's makes for a very quick feedback cycle.
  - Color scheme inspired by (and butchered, I'm no designer) [Wolfenstein](https://www.gameuidatabase.com/gameData.php?id=441)
- [i3](https://i3wm.org/) status bar integration (via [i3blocks](https://github.com/vivien/i3blocks))
- [systemd](https://wiki.archlinux.org/title/systemd) integration

Prior art: Gitlab variant ([gitlab-mr-bot](https://gitlab.com/chelmertz/gitlab-mr-bot/))
  - I made this while at my previous employer who self-hosted Gitlab, where it worked fine.
  - Gitlab's API for threads/comments on PRs is much easier to deal with
    (_elly_ needs to guess if there are unanswered comments, since that
    schema is paginated in a nested way).
  - `gitlab-mr-bot` was used as a one-off script, storing state in SQLite, for
    others to query.
    - I'm leaning towards "hiding things behind an API is friendler", especially
      if one wants to host things remotely. The transparency/hackability of
      having SQLite is very nice though.







