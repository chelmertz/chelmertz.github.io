---
layout: post
title: "vsftpd - enabling FTP access on your Ubuntu server"
permalink: "/vsftpd-enabling-ftp-access-ubuntu-server/"
tags: [server]
---

I found a great but old guide to installing <a title="vsftpd" href="https://help.ubuntu.com/6.06/ubuntu/serverguide/C/ftp-server.html">FTP on Ubuntu</a>, which I've never done before.

It was a really smooth procedure. The only error I received was this:

<code>
Unable to connect to system bus: Failed to connect to socket /var/run/dbus/system_bus_socket: No such file or directory
</code>

which was easy to solve since that file only needs to exist and be writable/readable:

<pre>
<code>
sudo mkdir -p /var/run/dbus
sudo touch /var/run/dbus/system_bus_socket
</code>
</pre>

By design, the default settings allow you to log onto your machine with your common user credentials, and land in that folder (/home/<user>)
