---
date: "2010-05-12 03:15"
title: "Supply and demand"
permalink: "supply-and-demand"
tags: [thoughts]
---

Everybody knows how to backup but the important thing is to [know how to restore from that backup]("http://www.joelonsoftware.com/items/2009/12/14.html").

I’m thinking this POV often is missing when discussing documentation of software. The [Wikipedia article on Software Documentation](http://en.wikipedia.org/wiki/Software_documentation) is focusing a lot on pigeonholing the different type of documentation that <em>may exist</em> so that no type is left out. That’s sometimes what I feel myself when trying to document a process - have I really made sure that this way of documenting is #1 or should I just go over the alternatives again?

Hold your horses. From now on, every inch taken towards documenting something must be questioned.

### **Who** will read this?

What does the _person_ willing to read this **expect to extract** from this enunciation?

- Find clues _fast_ to write their own, compatible code.
- Understand possible bottlenecks or places where refactoring would help extending the software.

### Is there any way to make this process **automated**?

- PHPDoc is perfect for me. Short comments are generated when typing - doesn’t conflict with a mind set on the job at hand. Supports generation of a more thorough “book” without any effort.

### Can I **write better code** so that the comment becomes superfluous?

- Crystal clear, long names of variables and methods. Few arguments for each method (remember those parameter objects (Fowler)).

### In which **form** does the documentation best fit its purpose?
- UML, wiki, PHPDoc, essays.. I’d say: only document after code is written and tested. This does not include the project plan/spec, but the long time, maintenance documentation.
  - Document in different scopes:
    - Overview. Why does it exist? Which dependencies are there? (OS/compiled modules/ram/libraries..) I often learn better from figures.
	- One simple and one advanced use case.
	- Detailed. This is connected to dynamic typed languages such as PHP & JS: type hinting is needed for method signatures. Easily applied with PHPDoc.

Make documentation as automated as possible. Invalid documentation is worse than spaghetti code because you believe you got everything under control and then… there’s a deadline and you’re toast.
