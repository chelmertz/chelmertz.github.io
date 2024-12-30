---
date: "2011-03-03 18:00"
title: "Fastest way to auto-login via ssh to your remote server"
permalink: "fastest-auto-login-ssh-remote-server"
tags: [bash, server]
---

Via [Remote Deployment with Ant and Rsync](http://topecoders.blogspot.com/2010/04/remote-deployment-with-ant-and-rsync.html):

```
$ scp ~/.ssh/id_rsa.pub remote_user_name@remoteserver:remote_user_name.pub

remoteserver$ cat ~/remote_user_name.pub >> ~/.ssh/authorized_keys
```

That's an excerpt but it was enough for me, having a public ssh key generated since before.
