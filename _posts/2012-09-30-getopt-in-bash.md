---
permalink: "getopt-in-bash"
layout: post
title: "getopt in bash"
date: "2012-09-30 23:18"
comments: true
tags: bash
---

How to properly read command line arguments for a bash script and do it well even without getopt:

{% highlight bash %}
while :; do
        case "$1" in
                -h|--help)
                        echo "This is helping"
                        exit 0
                        ;;

                -s|--simple-flag)
                        echo "You passed a flag"
                        shift
                        ;;
                -p|--pass-option)
                        shift
                        echo "You set option to $1"
                        option="$1"
                        shift
                        ;;
                -*)
                        echo "That's a weird option"
                        exit 1
                *)
                        # No more args to read
                        break
                        ;;
        esac
done
{% endhighlight %}

I think the syntax is much cleaner than that of getopt's as well -- that is, if you've even got a bash version supporting getopt.

<ins>Edit: Thanks Willem D'Haese for pointing out the missing `esac`.</ins>
