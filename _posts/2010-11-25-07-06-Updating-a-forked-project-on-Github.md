---
layout: post
title: "Updating a forked project on Github"
permalink: "/updating-a-forked-project-on-github/"
tags: [git]
---

Ok, just noticed that a forked repository of mine wasn’t actually automatically updated by Github. Luckily enough, <a href="http://help.github.com/forking/">http://help.github.com/forking/</a> made the process pretty painless:

<code lang="bash">
git clone my-repo-with-read+write-rights
git remote add upstream original-repo-url
git fetch upstream
git merge upstream/master
git push origin master
</code>

Next time I’ll just repeat
<code lang="bash">
git fetch upstream
git merge upstream/master
git push origin master
</code>
