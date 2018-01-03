---
layout: post
title: "From Wordpress on Apache to mynt on nginx"
permalink: "/wordpress-apache-to-mynt-nginx"
tags: "project"
---

Going from Wordpress to [mynt](http://mynt.mirroredwhite.com/).

## Purpose:

- smaller footprint, tune down vps (mysql & apache + my bad tuning skills = ram gets eaten up)
- write posts in markdown
- use CLI for faster editing
- posts in git
- local preview for free

## Bonuses:

- getting to know a python project (mynt)
- trying out nginx which is recommended for static content
  - before: mysql & apache processes hanging around, using 500 mb ram with near 0 load
  - response time (with wordpress cache plugin) was 2.49 s for 100.55 KB split on 17 requests (front page)
  - $ ab -n 1000 -c 5 http://iamnearlythere.com/ crashes
- take the opportunity to learn less for a redesign, fits with the pregeneration. also, include responsive design (octopress = good influence)

## Good to have:

- rsync for deploying
- log 404's
- nice source syntax (see examples on mynt's web)


## Procedure:

- Install mynt, requires xcode on osx (for "watchdog" -- an, in my eyes, unnecessary dependency)
- Export all posts from wordpress (I got an XML file out of 3.3.1)
- Run `$ ./wp2mynt.php export-file.xml` (<a href="https://github.com/chelmertz/mynt-tools/blob/master/wp2mynt.php">wp2mynt.php is on github</a>)
- Modify the theme a bit (far too many h1 used, for example)
  - Get a pattern from subtlepatterns.com
- Upload it to a new /var/www/directory-here, add a virtual host section to apache's config with ServerName test.iamnearlythere.com. Forgot to register test.iamnearlythere.com as an A record at my DNS registrar, 30 min wasted.
- Install nginx and set the default mimetype to text/html since I was silly enough to use the URI form of url.com/post-title which made mynt generate posts as files without extensions which by default are served as application/octet-stream by nginx ("want to download this?") and served as text/plain by apache.
  - The request dropped to 6 requests at 190 kb, totalling 1.47 s.
- Since the machine now no longer needs apache or mysql:
  - Make sure apache doesn't start on reload: $ sudo update-rc.d -f apache2 remove
  - Make sure mysqld doesn't start on boot: http://superuser.com/a/139059/9539
- Make sure nginx does start on boot: http://articles.slicehost.com/2007/10/17/ubuntu-lts-adding-an-nginx-init-script (I had to replace the nginx path with /usr/local/nginx/sbin/nginx)
- Allright, I'm done.. ish. When using `$ tail -f /usr/local/nginx/logs/error.log` I notice a bunch of stuff I missed. Redirections here we go
- Also, <a href="https://github.com/chelmertz/mynt-tools/blob/master/Makefile">create a Makefile</a> for generating, testing and deploying the blog (mynt generate; rsync)
  - It contains [multiline strings](http://stackoverflow.com/questions/649246/is-it-possible-to-create-a-multi-line-string-variable-in-a-makefile) and [feeding vim with input](http://superuser.com/questions/421367/read-from-stdin-to-new-named-file-in-vim)

## Result:

- Memory on vps could be turned down from 768 MB to 128, without tweaking nginx. Currently using about 20%.
- `$ ab -n 10000 -c 25` can be run without any problems what so ever

## Conclusion:

I'm pretty sure I suck at micromanaging apache & mysql. And that nginx is pretty fly with static files.

Also, a hiccup: I issued a 301 redirect in one case, and it went badly -- my nginx non-ninja-config-skils made me issue a recursive redirect loop. 301 responses are heavily cached by browsers, I had to delete the cache after removing the direction out of the config file. Do not make this mistake during 'live hours' or with a heavily trafficked page.

