---
layout: post
title: "Supply and demand"
permalink: "/supply-and-demand/"
tags: [thoughts]
---

Everybody knows how to backup but the important thing is to <a href="http://www.joelonsoftware.com/items/2009/12/14.html">know how to restore from that backup</a>.

I’m thinking this POV often is missing when discussing documentation of software. The <a href="http://en.wikipedia.org/wiki/Software_documentation">Wikipedia article on Software Documentation</a> is focusing a lot on pigeonholing the different type of documentation that <em>may exist</em> so that no type is left out. That’s sometimes what I feel myself when trying to document a process - have I really made sure that this way of documenting is #1 or should I just go over the alternatives again?

Hold your horses. From now on, every inch taken towards documenting something must be questioned.
<h3><strong>Who</strong> will read this?</h3>
<h3>What does the <em>person</em> willing to read this <strong>expect to extract</strong> from this enunciation?</h3>
<ul>
	<li>Find clues <em>fast</em> to write their own, compatible code.</li>
	<li>Understand possible bottlenecks or places where refactoring would help extending the software.</li>
</ul>
<h3>Is there any way to make this process <strong>automated</strong>?</h3>
<ul>
	<li>PHPDoc is perfect for me. Short comments are generated when typing - doesn’t conflict with a mind set on the job at hand. Supports generation of a more thorough “book” without any effort.</li>
</ul>
<h3>Can I <strong>write better code</strong> so that the comment becomes superfluous?</h3>
<ul>
	<li>Crystal clear, long names of variables and methods. Few arguments for each method (remember those parameter objects (Fowler)).</li>
</ul>
<h3>In which <strong>form</strong> does the documentation best fit its purpose?</h3>
<ul>
	<li>UML, wiki, PHPDoc, essays.. I’d say: only document after code is written and tested. This does not include the project plan/spec, but the long time, maintenance documentation.
<ul>
	<li>Document in different scopes:
<ul>
	<li>Overview. Why does it exist? Which dependencies are there? (OS/compiled modules/ram/libraries..) I often learn better from figures.</li>
	<li>One simple and one advanced use case.</li>
	<li>Detailed. This is connected to dynamic typed languages such as PHP &amp; JS: type hinting is needed for method signatures. Easily applied with PHPDoc.</li>
</ul>
</li>
</ul>
</li>
</ul>
Make documentation as automated as possible. Invalid documentation is worse than spaghetti code because you believe you got everything under control and then… there’s a deadline and you’re toast.
