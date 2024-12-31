---
date: "2010-12-08 08:20"
title: "Roadmap of PHP development - follow it"
permalink: "roadmap-of-php-development-follow-it"
tags: [php]
---

I just discovered an interesting wiki part of php.net, http://wiki.php.net/rfc

It contains drafts, proposals in an easy-on-the-eyes format. Some of the more interesting topics are

## [Function Array Dereferencing](http://wiki.php.net/rfc/functionarraydereferencing)
Example:

```php
function a() { return array(1, 2, 3);}
echo a()[0];
```

Status is listed as “accepted” - finally! This syntax is so handy, the only real alternative seems to be `list()` which is a bit clunky. Besides, everybody knows this syntax from javascript.

## [E_USER_DEPRECATED](http://wiki.php.net/rfc/e-user-deprecated-warning)
Finally a way to properly mark parts of your API as _old_:

```php
trigger_error("Use BlaBla::Bla() instead of ".__METHOD__, E_USER_DEPRECATED);
```

## [docBlock Parser](http://wiki.php.net/rfc/docblockparser)
A native parser for docblocks which will probably be a lot faster and also increase awareness and interest of actually documenting code. PHPUnit testcases and frameworks such as Doctrine 2 might gain some speed increase.

## [Strict type](http://wiki.php.net/rfc/typechecking)
There are some proposials (about time) for making a method signature like the following valid:

```php
function(int $apples, string $collection_name) {}
```

## [C# style getters/setters](http://wiki.php.net/rfc/propertygetsetsyntax)
A really usable syntax, instead of `__get()`, `__set()` or `__call()`:

```php
class TimePeriod
{
    private double seconds;

    public double Hours
    {
        get { return seconds / 3600; }
        set { seconds = value * 3600; }// The variable "value" holds the incoming value to be "set"
    }
}
```

The proposal then tries to include every possible aspect of this, from the beginning, good idea. Just by reading the discussion on the wiki (never read any mailing lists) I’m already getting tired of it. Too bad.

## Conclusion

There exists some really nice discussions in an easy-to-read format. Since there are quite a few proposals already implemented, I probably got hoping for too many new features… In one way, it’s great to see the language improving, in another way: why is it taking so long?

Also, it feels nice to see inspiration coming from other languages (surprisingly javascript stands for quite a few). I wonder if it will become more and more like java, with all the oop/namespacing in the latest patches. (I hope so.) [The more global functions we get rid off, the better.](/global-code-sucks-php-has-5468-global-functions/)