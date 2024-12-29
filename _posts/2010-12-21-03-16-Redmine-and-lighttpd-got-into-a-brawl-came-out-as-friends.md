---
date: "2010-12-21 03:16"
title: "Redmine and lighttpd got into a brawl, came out as friends"
permalink: "redmine-and-lighttpd-got-into-a-brawl-came-out-as-friends"
tags: [server]
---

So much for the [easy setup of Redmine](/installing-redmine-on-ubuntu-and-lighttpd), our server gazed at me with a 500 Internal Server error the following monday morning.

To track down the error, I started out by tailing Redmine’s error log that I previously defined in <em>/etc/lighttpd/lighttpd.conf</em>, showing:

```
$ tail -f /var/log/lighttpd/redmine_error.log
2010-12-20 12:09:04: (mod_fastcgi.c.2582) unexpected end-of-file (perhaps the fastcgi process died): pid: 6425 socket: unix:/tmp/redmine.socket-0
2010-12-20 12:09:04: (mod_fastcgi.c.3367) response not received, request sent: 909 on socket: unix:/tmp/redmine.socket-0 for /dispatch.fcgi?, closing connection                                                                          
2010-12-20 12:09:04: (mod_fastcgi.c.1734) connect failed: Connection refused on unix:/tmp/redmine.socket-3
2010-12-20 12:09:04: (mod_fastcgi.c.3037) backend died; we'll disable it for 1 seconds and send the request to another backend instead: reconnects: 0 load: 1
```

What I did to solve it was (after some [googling](http://www.webhostingtalk.com/archive/index.php/t-662362.html)) to change the lighttpd conf setting for Redmine’s socket from <em>/tmp/redmine.socket</em> to <em>/tmp/lighttpd/redmine.socket</em>. Before I could do that, I had to

```
$ sudo mkdir /tmp/lighttpd
$ sudo chown -R www-data:www-data /tmp/lighttpd
```

to make sure it was going to be writable. Restarting lighttpd after that seems to have solved the issue, permanently I hope.
