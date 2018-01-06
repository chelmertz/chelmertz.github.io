---
layout: post
title: "Debugging some php.."
permalink: "/debugging-some-php/"
tags: [php, server]
---

<code>apt-get install php5-dev</code>
<code>pecl install xdebug</code>

edit php/conf.d/xdebug.ini - xdebug autoloads if correct, check phpinfo(). Set path to [..].so file for xdebug, set <code>xdebug.profiler_enable_trigger = 1</code> and <code>xdebug.profiler_enable = 0</code>

install kcachegrind (or webgrind from google code)
