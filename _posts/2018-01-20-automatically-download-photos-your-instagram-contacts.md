---
permalink: "automatically-download-photos-your-instagram-contacts"
layout: post
title: "Automatically download new photos from your Instagram contacts"
date: "2018-01-20 21:24"
tags: bash
comments: true
---

Let's say you want to downloaded the images uploaded to Instagram by you and your contacts, for backup reasons.

You can login at the regular web application at <https://instagram.com> and [right-click/inspect your way to the actual URI of the image file](https://honeypotmarketing.com/save-a-photo-from-instagram/). Manually downloading one photo at a time soon becomes boring though, let's look for other alternatives.

There are a couple of sites that promote applications that does the job for you. The ones I saw wanted my username and password, which is out of the question.

Let's try to do download the images ourselves instead, from an environment we control, our desktop.

## Installing a scraper on your own computer to download photos

*The rest of this blog post assumes that you are using **Linux or macOS** (for `cron` and regular CLI tools), and have **python** installed (for the scraping program).*

There's an easy to use [scraper](https://en.wikipedia.org/wiki/Web_scraping) by Richard Arcega called [Instagram Scraper](https://github.com/rarcega/instagram-scraper/) that we will use (version 1.5.18 at the time of writing). The scraper is available on [pypi](https://pypi.python.org/pypi/instagram-scraper), and you install it with:


{% highlight bash %}
sudo pip install instagram-scraper
{% endhighlight %}

The next thing we are going to do is to configure who you are and which profiles you want to download. You do that by creating a text file:

{% highlight bash %}
$ cat /home/ch/insta.txt
-u=username
-p=passwordinplaintext
--latest
--destination=/home/ch/Pictures/instagram
--retain_username
a_user_i_follow
another_user_i_follow
even_a_third_user_i_follow
{% endhighlight %}

Let's try it out:

{% highlight bash %}
$ instagram-scraper @/home/ch/insta.txt
{% endhighlight %}

If you've got image files in the folder referred to by `--destination`, you are home free.

## Automating the scraping

To avoid executing the above command manually, we use [cron](https://en.wikipedia.org/wiki/Cron) to fetch new images for us.

This is what a my crontab looks like:

{% highlight bash %}
$ crontab -l
1 * * * * /usr/bin/instagram-scraper -q @/home/ch/insta.txt
{% endhighlight %}

To edit your own crontab, use the command `crontab -e`.

If you need to learn the syntax of cron, there's a handy form at <https://crontab.guru/> that could help you.

## Viewing newly downloaded images

I use a minimal image viewer called [sxiv](https://github.com/muennich/sxiv) which is available in the default repos of, at least, Fedora.

The following command let's you view all recently downloaded images:

{% highlight bash %}
$ sxiv $(find /home/ch/Pictures/instagram -mtime -10 -type f -name "*jpg")
{% endhighlight %}

where the first argument to `find` is the value of `--destination` in your configuration, and `-10` let's you view images from the last ten days.

## Caveats
### Notes about availability of images

All images uploaded to Instagram, even by those with their profiles set to private, are publicly accessible. Getting to know the URI is the only problem, that's why we use the scraping application, to find the correct URIs. Even though the URIs themselves do not require being authenticated to download, there may be rate limits or other counter measures against downloading images in bulk.

### Shortcuts in the code

- Feel free to `pip install --user` instead.
- Feel free to capture the stdout/stderr from the cronjob.

*[URI]: Uniform Resource Identifier
*[CDN]: Content delivery network
