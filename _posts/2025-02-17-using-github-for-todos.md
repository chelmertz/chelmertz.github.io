---
permalink: "using-github-for-todos"
title: "Using Github issues for TODOs"
summary: "Easily create a lot of TODOs, that ends up in a consistent place in Github"
date: "2025-02-17 14:33"
---

I've got an executable script named `todo` in my `$PATH`, and it looks like this:

```shell
#!/bin/sh

# REPO formatted as user/repo
test -z "$GH_TODO_REPO" && echo missing GH_TODO_REPO && exit 1
# PROJECT formatted in plaintext
# (I would skip the spaces, I think I had issues with that even when quoting)
test -z "$GH_TODO_PROJECT" && echo missing GH_TODO_PROJECT && exit 1

if [ -t 0 ]; then
        # should be within a terminal

        # requires access rights: gh auth refresh -s project
        gh issue create --repo "$GH_TODO_REPO" \
                --project "$GH_TODO_PROJECT" \
                --assignee @me \
                --milestone 2025:1 \
                --editor
        res=$?
        exit $res
fi
echo Launch through terminal
exit 1
```

Complementing that script, I have a clickable "assigned to me" todo-counter always visible in the status bar in my [window manager](https://i3wm.org/), courtesy of [i3blocks](https://github.com/vivien/i3blocks):

```shell
#!/usr/bin/env bash

# --owner responds to the first part of the path in Github URLs,
# in this case, it's my employer
count=$(gh search issues --assignee @me --owner matchiapp --state open --json url --jq 'length')

# glyph from fontawesome, "bug" f188

# long
echo " $count"
# short
echo " $count"
# color
echo "#00ff00"

__open() {
	url=$1
        # open the URL and tell i3 to focus the workspace,
        # and use wmctrl to focus the correct window
	xdg-open $url &>/dev/null && i3-msg workspace number 2 &>/dev/null && wmctrl -a firefox &>/dev/null
}

if [[ "$BLOCK_BUTTON" -eq 1 ]]; then
	# left click
	# "my" board
	__open https://github.com/orgs/matchiapp/projects/32
elif [[ "$BLOCK_BUTTON" -eq 3 ]]; then
	# right click
	# all of my issues
	__open https://github.com/issues/assigned
fi
```

The above then renders like this:

![i3blocks widget, a green bug icon and the number 20](/assets/i3blocks_gh_todos.png)