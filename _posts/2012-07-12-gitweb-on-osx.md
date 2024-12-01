---
permalink: "gitweb-on-osx"
layout: post
title: "Gitweb on OSX"
date: "2012-07-12 15:10"
tags: git mac
---

I have all my code stored in ~/Sites and every sub folder represents a project. Each folder is also a git repository. This means there's history attached to it, which could be much more easily navigated than through the finder or a terminal. Some of the projects are mirrored on [github](https://github.com/chelmertz) but it's too messy to keep them all there, although -- that kind of overview is nice to have.

My solution to getting an overview was to setup gitweb locally, which makes for a very fast solution. This is how you do it on your OS X-machine:

- Find where git's installed at. I use [homebrew](http://mxcl.github.com/homebrew/), which means git (by default) resides in /usr/local/Cellar/git.

- Grab the gitweb cgi, copy it to where your local Apache server can find it:

{% highlight bash %}
# cp /usr/local/Cellar/git/1.7.11.1/share/gitweb/gitweb.cgi /Library/WebServer/CGI-Executables
# cp /usr/local/Cellar/git/1.7.11.1/share/gitweb/static /Library/WebServer/Documents/gitweb
{% endhighlight %}

- Modify gitweb.cgi to let it find the static files, to find the git projects and to not choke on really large directories that I'm sure does not contain anything git' anyway:

{% highlight diff %}
--- /usr/local/Cellar/git/1.7.11.1/share/gitweb/gitweb.cgi	2012-07-11 20:23:10.000000000 +0200
+++ /Library/WebServer/CGI-Executables/gitweb.cgi	2012-07-11 20:50:36.000000000 +0200
@@ -71,11 +71,11 @@
 
 # absolute fs-path which will be prepended to the project path
 #our $projectroot = "/pub/scm";
-our $projectroot = "/pub/git";
+our $projectroot = "/Users/my-user/Sites";
 
 # fs traversing limit for getting project list
 # the number is relative to the projectroot
-our $project_maxdepth = 2007;
+our $project_maxdepth = 5;
 
 # string of the home link on top of all pages
 our $home_link_str = "projects";
@@ -95,15 +95,15 @@
 our $site_footer = "";
 
 # URI of stylesheets
-our @stylesheets = ("static/gitweb.css");
+our @stylesheets = ("/gitweb/static/gitweb.css");
 # URI of a single stylesheet, which can be overridden in GITWEB_CONFIG.
 our $stylesheet = undef;
 # URI of GIT logo (72x27 size)
-our $logo = "static/git-logo.png";
+our $logo = "/gitweb/static/git-logo.png";
 # URI of GIT favicon, assumed to be image/png type
-our $favicon = "static/git-favicon.png";
+our $favicon = "/gitweb/static/git-favicon.png";
 # URI of gitweb.js (JavaScript code for gitweb)
-our $javascript = "static/gitweb.js";
+our $javascript = "/gitweb/static/gitweb.js";
 
 # URI and label (title) of GIT logo link
 #our $logo_url = "http://www.kernel.org/pub/software/scm/git/docs/";
{% endhighlight %}

Now, visit & bookmark http://localhost/cgi-bin/gitweb.cgi in your browser and you should be good to go.
