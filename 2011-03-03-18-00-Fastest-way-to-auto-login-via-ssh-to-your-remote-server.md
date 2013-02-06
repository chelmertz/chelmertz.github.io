---
layout: post
title: "Fastest way to auto-login via ssh to your remote server"
permalink: "/fastest-auto-login-ssh-remote-server"
categories: [bash, server]
---

Via <a href="http://topecoders.blogspot.com/2010/04/remote-deployment-with-ant-and-rsync.html">Remote Deployment with Ant and Rsync</a>:
<pre>$ scp ~/.ssh/id_rsa.pub remote_user_name@remoteserver:remote_user_name.pub

remoteserver$ cat ~/remote_user_name.pub &gt;&gt; ~/.ssh/authorized_keys</pre>
That's an excerpt but it was enough for me, having a public ssh key generated since before.
