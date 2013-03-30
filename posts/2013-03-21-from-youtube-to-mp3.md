---
layout: post
title: "From Youtube to mp3"
date: 2013-03-21 21:28
categories: tools
---

Because I always forget the exact commands:

{% codeblock lang:bash %}
youtube-dl https://youtube.com/?v=id_here
ffmpeg -i id_here.flv -acodec mp3 output.mp3
{% endcodeblock %}

On OS X, you'll find both youtube-dl and ffmpeg through homebrew.
