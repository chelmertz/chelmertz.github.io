---
date: "2010-09-29 19:59"
title: "Put a tail on that error log"
permalink: "put-a-tail-on-that-error-log"
tags: [server]
---

Place the following in a file called _/usr/local/bin/tail_log_:

```shell
tail -f /var/log/lighttpd/php_error_log|egrep -i "fatal|stack|#|thrown"
```

That enables a user to issue _tail_log_ to listen for fatals in the php error log, usable for dev enviroments when developing in new realms.

Note that the location of _php_error_log_ might differ, especially if the server isn’t lighttpd…
