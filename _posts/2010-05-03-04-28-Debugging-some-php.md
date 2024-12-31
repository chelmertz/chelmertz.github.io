---
date: "2010-05-03 04:28"
title: "Debugging some php.."
permalink: "debugging-some-php"
tags: [php, server]
---

```shell
apt-get install php5-dev
pecl install xdebug
```

edit _php/conf.d/xdebug.ini_ - xdebug autoloads if correct, check `phpinfo()`.

Set path to ....so file for xdebug, set `xdebug.profiler_enable_trigger = 1` and `xdebug.profiler_enable = 0`

install kcachegrind (or webgrind from google code)
