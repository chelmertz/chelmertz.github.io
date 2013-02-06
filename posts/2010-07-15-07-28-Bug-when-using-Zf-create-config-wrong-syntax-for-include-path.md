---
layout: post
title: "Bug when using `Zf create config`, wrong syntax for include path"
permalink: "/bug-when-using-zf-create-config-wrong-syntax-for-include-path"
categories: [netbeans, zend-framework]
---

<a href="https://netbeans.org/bugzilla/show_bug.cgi?format=multiple&amp;id=187234">https://netbeans.org/bugzilla/show_bug.cgi?format=multiple&amp;id=187234</a>

The NetBeans - through <code>zf create config</code> - generated file <em>~/.zf.ini</em> contains the following error:
<code>php.includepath</code> should be <strong><code>php.include_path</code></strong>
