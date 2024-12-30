---
date: "2010-07-15 07:28"
title: "Bug when using `Zf create config`, wrong syntax for include path"
permalink: "bug-when-using-zf-create-config-wrong-syntax-for-include-path"
tags: [netbeans, zend-framework]
---

https://netbeans.org/bugzilla/show_bug.cgi?format=multiple&id=187234

The NetBeans - through `zf create config` - generated file _~/.zf.ini_ contains the following error:
`php.includepath` should be `php.include_path`
