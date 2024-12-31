---
date: "2010-07-14 06:20"
title: "Set enviroment variables in Apache"
permalink: "set-enviroment-variables-in-apache"
tags: [server, zend-framework]
---

In .htaccess:

```
SetEnv APPLICATION_ENV development
```

In php (from ZF’s bootstrapped setup):

```php
// Define application environment
defined('APPLICATION_ENV')
|| define('APPLICATION_ENV', (getenv('APPLICATION_ENV') ? getenv('APPLICATION_ENV') : 'production'));
```
