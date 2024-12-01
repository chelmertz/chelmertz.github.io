---
date: "2011-11-28 21:49"
layout: post
title: "Handling segmentation faults with gdb"
permalink: "handling-segmentation-faults-gdb"
tags: [c, usability]
---

As a C newbie, I've never gotten a segfault until tonight (yes, my homework's due in a couple of hours..).

gdb helped out a lot by letting me get from "segfault happened" to "function x and y were called but z failed at my_file.c:30, here's that line for you: ...".

The slide <a href="http://www.slideshare.net/noobyahoo/introduction-to-segmentation-fault-handling-5563036">Introduction to segmentation fault handling</a> rocks; based on that article, this is kinda what my process & debugging step in the makefile looks like now:

    gcc -Wall -c -g *.c
    gcc -o my_program *.o
    gdb my_program
    (inside gdb) run
    (inside gdb) <..make your program segfault..>
    (inside gdb) backtrace

Sweet!
