---
date: "2011-10-01 19:11"
title: "Remap caps-lock to esc in OS X Lion"
permalink: "map-esc-caps-lock-os-lion"
tags: [mac]
---

The reason for removing caps-lock is mainly about being faster in vim at exiting entered modes, but also since I never use it and it's always in the way. Instead of just removing the key, we could make it usable though, and bind it to escape. The process is fairly easy.

- First you need to disable the key "controlling case sensitivity" (= caps lock). You find the option in _System Preferences > Keyboard > Modifier keys (in the first tab)_.
- Next, to bind caps-lock to the functionality of esc, install the [PCKeyboardHack](http://pqrs.org/macosx/keyremap4macbook/extra.html)
- Follow the instructions on the page in the link and use the keycode `53` for escape.
- You're done.

Via [Using Caps Lock as Esc in Mac OS X](http://stackoverflow.com/questions/127591/using-caps-lock-as-esc-in-mac-os-x)
