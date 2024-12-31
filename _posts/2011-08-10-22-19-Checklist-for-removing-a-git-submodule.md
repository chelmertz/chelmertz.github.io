---
date: "2011-08-10 22:19"
title: "Checklist for removing a git submodule"
permalink: "checklist-removing-git-submodule"
tags: [git]
---


>- Delete the relevant line from the _.gitmodules_ file.
>- Delete the relevant section from _.git/config_.
>- Run _git rm --cached path_to_submodule_ (no trailing slash).
>- Commit the superproject.
>- Delete the now untracked submodule files.


From https://git.wiki.kernel.org/index.php/GitSubmoduleTutorial
