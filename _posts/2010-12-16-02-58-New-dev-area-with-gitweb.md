---
date: "2010-12-16 02:58"
title: "New dev area with gitweb"
permalink: "new-dev-area-with-gitweb"
tags: [git, server]
---

I’m currently setting up a new dev server at work. This includes hands-on work such as writing PHP organized in namespaces, but at a more technical level, modifying lighttpd and installing lots of packages on a Ubuntu server.

The dev server which is being deprecated is equipped with svn & trac, which has been really nice, especially for browsing the source code, code reviews and managing tickets, so the goals are set for this new server.

Gitweb will correspond to the demands for viewing source code, diffs and branches. Installing it and getting it up on lighttpd was a breeze, thanks to the guide [How to setup gitweb with lighttpd on Ubuntu - jonathanrobson.me](http://jonathanrobson.me/2009/11/how-to-setup-gitweb-with-lighttpd-on-ubuntu).

I used his config with some tweaks. The `$projectroot` variable in _/etc/gitweb.conf_ now points at a directory which has a clone of each of our reposes. I.e.

```
$ cd /home/git/projects
```

if _/home/git/projects/_ is where you store all of your repositories, in a common place as a base for your dev or stage area. Read my entry [Using git init —bare for a slim repository look](/using-git-init-bare-for-a-slim-repository-look) for a directory layout that works.

If you followed my approach,

```
$ ls /home/git/projects
```

should only consist of directories called `project-1.git`, `project-2.git` and so forth. This is ideal for gitweb, which prefers to have a single directory to check for git repositories. (Check jonathanrobson’s link above for the whole shebang.)

## Bonus feature: styling

There’s a github repository on [https://github.com/kogakure/gitweb-theme](https://github.com/kogakure/gitweb-theme) which has an easily installed theme that looks much better than the original bundled with gitweb.
