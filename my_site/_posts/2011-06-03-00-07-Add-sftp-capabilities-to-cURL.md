---
layout: post
title: "Add sftp-capabilities to cURL"
permalink: "/add-sftp-capabilities-to-curl/"
tags: [server]
---

I just ran into a problem while importing some cronjobs to a new server - apparently the sftp protocol is not supported by a newly installed cURL's default settings, which in turn made the cronjob pretty useless. The good thing was the error message, cURL is usually really good at helping users get going again:

>curl: (1) Protocol sftp not supported or disabled in libcurl

I went on checking features, following someone else's steps:

    $ curl -V
    curl 7.19.7 (x86_64-pc-linux-gnu) libcurl/7.19.7 OpenSSL/0.9.8k zlib/1.2.3.3 libidn/1.15
    Protocols: tftp ftp telnet dict ldap ldaps http file https ftps
    Features: GSS-Negotiate IDN IPv6 Largefile NTLM SSL libz

There it was, right out in the open: no sftp support. Onwards, googling led me to the third post in <a href="https://bugs.launchpad.net/ubuntu/+source/curl/+bug/311029">this thread</a> which I'm quoting:

To generate a curl with ssh support on Ubuntu 9.04:

    sudo apt-get install build-essential debhelper libssh2-1-dev
    sudo apt-get source libcurl3
    sudo apt-get build-dep libcurl3

    cd curl-7.18.2/debian

    gedit rules

    find and replace "--without-libssh2" with "--with-libssh2"

    cd ..

    sudo dpkg-buildpackage

    cd ..

    sudo dpkg -i curl_7.18.2-8ubuntu4.1_amd64.deb
    sudo dpkg -i libcurl3_7.18.2-8ubuntu4.1_amd64.deb
    sudo dpkg -i libcurl3-gnutls_7.18.2-8ubuntu4.1_amd64.deb

Those steps got it working for me, even though I'm on Ubuntu 10.10. Worth mentioning: I did not uninstall cURL before starting out. Some of the version numbers are newer but the process I went through was identical. Checking cURL's features, again, gave me:

    $ curl -V
    curl 7.19.7 (x86_64-pc-linux-gnu) libcurl/7.19.7 OpenSSL/0.9.8k zlib/1.2.3.3 libidn/1.15 libssh2/1.2.2
    Protocols: tftp ftp telnet dict ldap ldaps http file https ftps scp sftp
    Features: GSS-Negotiate IDN IPv6 Largefile NTLM SSL libz

alas, there is support for scp as well as sftp - nice!
