---
date: "2010-12-02 07:27"
title: "Using git init --bare for a slim repository look"
permalink: "using-git-init-bare-for-a-slim-repository-look"
tags: [git, server]
---

Until now, I’ve mainly worked alone with my git repositories. Sometime pushing to [my github-account](https://github.com/chelmertz), but always using the service as a kind of backup solution. Until now. At work, we’re setting up our new environment with git and I’m lucky enough to learn it while doing it.

This setting will differ quite a lot from what I’m used to:

- for private projects, the git repository is my project folder;
- for work, the repository has to be _synced across multiple users_, a base for _deployment_ and _act as a more important repository_ than our local ones, kind of what SVN was doing and what, I’m told, git isn’t really made for.

Also, exercising a plain `git init` routine creates a complete repository with all source files, branches & stuff, which really are ment to keep track off and develop in. Let’s not, let’s just keep the repository there and only `git push` there so _all development will be done in each user’s local repository_.

Let’s get to the good stuff:

```
# create a user named 'git' and make that user the owner of your repositories, and nothing else
$ mkdir -p /home/git/repositories
$ cd /home/git/repositories
$ mkdir your-project.git && cd your-project.git
$ git init --bare
$ git config receive.denyCurrentBranch "ignore"
$ mkdir -p ~/www/your-project && cd ~/www/your-project
$ echo "lalala, random file" > random.txt && git add . && git commit -m "First commit"
$ git remote add origin /home/git/repositories/your-project.git
$ git push origin master
```

There you go, a fully capable external repository, local on your dev machine, ready to export to a staging/testing/Q&A-setting. For an easy setup of this, go read _[Using Git to manage a web site](http://toroid.org/ams/git-website-howto)_ which _worked out of the box_ for us.

If you get any errors whilst doing this, google it or just try to parse it (for us, it mostly came down to `chown`- and `chmod`-ing the `-R /home/git/repositories/` folders, so using `sudo` helps a lot).
