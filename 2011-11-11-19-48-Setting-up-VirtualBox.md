---
layout: post
title: "Setting up VirtualBox"
permalink: "/setting-up-virtualbox"
categories: [virtualbox]
---

In order to access the virtual machine, select your machine's settings, pick network and use Bridged network. Also, don't forget to plugin the cable.

(Vague title in case I add more reminders.)

<hr />

<h2>Installing Guest Additions</h2>
<blockquote>To install Guest Additions, do the following:

Sudo apt-get install dkms

sudo apt-get install build-essential

Reboot the server by entering the command

sudo reboot

Login in to the server.

Go to the Virtualbox Devices entry on the menu bar of the guest OS and select Install Guest Additions… , this will load the Guest Additions ISO CD image.

Change directory to media

cd /media

Create a directory called cdrom, this will become our mount point

sudo mkdir cdrom

Mount the Guest Additions ISO to the mount point

sudo mount /dev/cdrom /media/cdrom

Now change to the cdrom directory

cd /media/cdrom

Display the directory contents i.e. the ISO image

ls

Depending on whether you are running a 32bit or 64bit OS, run the relevant installer. In this case 32bit so enter

sudo ./VBoxLinuxAdditions-x86.run</blockquote>
Citation from <a href="http://mylinuxramblings.wordpress.com/2010/06/03/installing-virtualbox-guest-additions-on-ubuntu-server-10-04/">mylinuxramblings.com</a>
