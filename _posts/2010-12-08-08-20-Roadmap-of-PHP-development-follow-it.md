---
layout: post
title: "Roadmap of PHP development - follow it"
permalink: "roadmap-of-php-development-follow-it"
tags: [php]
---

I just discovered an interesting wiki part of php.net, <a href="http://wiki.php.net/rfc"></a><a href="http://wiki.php.net/rfc">http://wiki.php.net/rfc</a>.

It contains drafts, proposals in an easy-on-the-eyes format. Some of the more interesting topics are
<h2><a href="http://wiki.php.net/rfc/functionarraydereferencing">Function Array Dereferencing</a></h2>
Example:
<div class="CodeRay">
<div class="code">
<pre>function a() { return array(1, 2, 3);}
echo a()[0];</pre>
</div>
</div>
Status is listed as “accepted” - finally! This syntax is so handy, the only real alternative seems to be <code>list()</code> which is a bit clunky. Besides, everybody knows this syntax from javascript.
<h2><a href="http://wiki.php.net/rfc/e-user-deprecated-warning">E_USER_DEPRECATED</a></h2>
Finally a way to properly mark parts of your API as <em>old</em>:
<div class="CodeRay">
<div class="code">
<pre>trigger_error("Use BlaBla::Bla() instead of ".__METHOD__, E_USER_DEPRECATED);</pre>
</div>
</div>
<h2><a href="http://wiki.php.net/rfc/docblockparser">docBlock Parser</a></h2>
A native parser for docblocks which will probably be a lot faster and also increase awareness and interest of actually documenting code. PHPUnit testcases and frameworks such as Doctrine 2 might gain some speed increase.
<h2><a href="http://wiki.php.net/rfc/typechecking">Strict type</a></h2>
There are some proposials (about time) for making a method signature like the following valid:
<div class="CodeRay">
<div class="code">
<pre>function(int $apples, string $collection_name) {}</pre>
</div>
</div>
<h2><a href="http://wiki.php.net/rfc/propertygetsetsyntax">C# style getters/setters</a></h2>
A really usable syntax, instead of <code>__get()</code>, <code>__set()</code> or <code>__call()</code>:
<div class="CodeRay">
<div class="code">
<pre>class TimePeriod
{
    private double seconds;

    public double Hours
    {
        get { return seconds / 3600; }
        set { seconds = value * 3600; }// The variable "value" holds the incoming value to be "set"
    }
}</pre>
</div>
</div>
The proposal then tries to include every possible aspect of this, from the beginning, good idea. Just by reading the discussion on the wiki (never read any mailing lists) I’m already getting tired of it. Too bad.
<h1>Conclusion</h1>
There exists some really nice discussions in an easy-to-read format. Since there are quite a few proposals already implemented, I probably got hoping for too many new features… In one way, it’s great to see the language improving, in another way: why is it taking so long?

Also, it feels nice to see inspiration coming from other languages (surprisingly javascript stands for quite a few). I wonder if it will become more and more like java, with all the oop/namespacing in the latest patches. (I hope so.) <a href="http://iamnearlythere.tumblr.com/post/1242918472/global-code-sucks-php-has-5468-global-functions">The more global functions we get rid off, the better.</a>
