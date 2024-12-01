---
permalink: "building-packages-on-solaris"
layout: post
title: "Building packages on Solaris"
date: "2012-08-30 22:44"
tags: 
---

This post will act as a rough guide for how you get your package building going on Solaris, from the perspective of someone who is used to Linux and has built an RPM or two.

The goal is to have a pkg-file with files to-be-installed in, installation, uninstallation and upgrade scripts, the dependencies should be handled.


## Making yourself comfortable

The Solaris 10 I've been working with (or rather, the Solaris 10 *zone*), is rather naked at start. Your .bashrc should probably include something like this:

TERM=xterm;export TERM

And set your shell to /opt/csw/bin/bash in /etc/passwd. Also, note that sudo is not installed by default, you might want to install it or use the ACL-like [pfexec](http://developers.sun.com/developer/technicalArticles/opensolaris/pfexec.html).

You need to setup your project so that the build tools will get you. First of all, install the [pkgutil](http://pkgutil.wikidot.com/get-install-and-configure) and [pkgutilplus](http://www.opencsw.org/packages/pkgutilplus/) packages.


## Preparing your project for packaging

Create a folder in your project to hold our metadata. I'll call it just that, "metadata". It should contain at least a [pkginfo file](http://www.garex.net/sun/packaging/pkginfo.html) but can contain several other files.


## Defining your project's dependencies

Since your program might rely on other packages presence, you'd want to know how to specify them. We will start by finding out the relation between the newly created pkginfo-file and the [pkginfo command](http://heirloom.sourceforge.net/pkgtools/pkginfo.1.html). Try typing `pkginfo -l [package name]` to see an already installed package's properties, or add the -d flag: `pkginfo -l -d [path to pkg file]` if you want to check the properties of a local copy of a package.

In the metadata folder


### Finding packages

In order to manage the dependencies, you will have to bundle in a [single, large, pkg-file](http://www.opencsw.org/manual/for-administrators/getting-started.html#creating-a-pkg-file-for-a-host-without-an-internet-connection). The dependencies must either be installed locally or found in any of your mirrors defined in `/etc/opt/csw/pkgutil.conf` (much like apt-get's `/etc/apt/sources.list` or yum's `/etc/yum.repos.d/*.repo`-files).


## Decide what the installation will look like

You can help users install your package by providing an [administration file](http://www.opensolarisforum.org/man/man4/admin.html).


pkgadd

pkginfo

pkgproto

pkgtrans

bldcat --stream

chkcat

prototype

administration file

Further reading:
  - [Creating pkgadd Software Packages under Solaris](http://www.sunfreeware.com/pkgadd.html)
