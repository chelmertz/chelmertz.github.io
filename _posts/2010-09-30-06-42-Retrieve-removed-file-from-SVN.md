---
date: "2010-09-30 06:42"
title: "Retrieve removed file from SVN"
permalink: "retrieve-removed-file-from-svn"
tags: [svn]
---

    svn copy -r a-revision http://some-repo/some-file .

Pretty obvious really, but still. Use the remote path to the file that existed in that revision and place it in the current dir: . (or anywhere really).

PS: this retrieved the contents of  but you would also need to `svn commit` any changes to reinsert the file into the repos.
