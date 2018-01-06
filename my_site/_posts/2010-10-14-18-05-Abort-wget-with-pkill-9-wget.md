---
layout: post
title: "Abort wget with pkill -9 wget"
permalink: "/abort-wget-with-pkill-9-wget/"
tags: [bash]
---

<a href="http://stackoverflow.com/questions/3410730">http://stackoverflow.com/questions/3410730</a>

Sometimes, <code>wget</code> plays a trick and refuses to switch off if a URL is redirecting you around and around. CTRL+C did not cancel the process for me but <code>pkill -9 wget</code> did.
