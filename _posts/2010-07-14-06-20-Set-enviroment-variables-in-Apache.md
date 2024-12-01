---
layout: post
title: "Set enviroment variables in Apache"
permalink: "set-enviroment-variables-in-apache"
tags: [server, zend-framework]
---

In .htaccess:

<code>SetEnv APPLICATION_ENV development</code>

In php (from ZF’s bootstrapped setup):
<code>// Define application environment</code>
<code>defined('APPLICATION_ENV')</code>
<code>|| define('APPLICATION_ENV', (getenv('APPLICATION_ENV') ? getenv('APPLICATION_ENV') : 'production'));</code>
