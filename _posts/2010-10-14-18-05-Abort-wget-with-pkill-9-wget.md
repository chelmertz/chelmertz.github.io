---
date: "2010-10-14 18:05"
title: "Abort wget with pkill -9 wget"
permalink: "abort-wget-with-pkill-9-wget"
tags: [bash]
---

http://stackoverflow.com/questions/3410730

Sometimes, `wget` plays a trick and refuses to switch off if a URL is redirecting you around and around. CTRL+C did not cancel the process for me but `pkill -9 wget` did.
