---
date: "2010-11-25 07:06"
title: "Updating a forked project on Github"
permalink: "updating-a-forked-project-on-github"
tags: [git]
---

Ok, just noticed that a forked repository of mine wasn’t actually automatically updated by Github. Luckily enough, http://help.github.com/forking/ made the process pretty painless:

```shell
git clone my-repo-with-read+write-rights
git remote add upstream original-repo-url
git fetch upstream
git merge upstream/master
git push origin master
```

Next time I’ll just repeat

```shell
git fetch upstream
git merge upstream/master
git push origin master
```
