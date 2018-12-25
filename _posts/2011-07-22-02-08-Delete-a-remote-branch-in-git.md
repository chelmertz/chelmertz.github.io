---
layout: post
title: "Delete a remote branch in git"
permalink: "/delete-remote-branch-git/"
tags: [git]
---

Prefix the branch's name with a colon to delete it from the remote repository. I.e. `git push origin :awesome-feature` if the deleted branch was called "awesome-feature" and lived in the remote repository called "origin".

Of course you need to have pushed the branch to the remote repository on beforehand, since only your master branch is pushed by default.
