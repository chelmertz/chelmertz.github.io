---
layout: post
title: "Requesting new magic method: `__toBoolean()`"
permalink: "/requesting-new-magic-method-__toboolean"
tags: [php]
---

What if you only want to display your active users in a very simple way? Consider the following:
<div class="CodeRay">
<div class="code">
<pre>class User {
  public $active = true;
  private $_name;
  public function __construct($name) {
    $this-&gt;_name = (string) $name;
  }
  public function __toBoolean() {
    return (bool) $this-&gt;active;
  }
  public function getName() {
    return $this-&gt;_name;
  }
}

$emily = new User('Emily');
$emily-&gt;active = false;
$sara = new User('Sara');

$users = array($emily, $sara);
foreach($users as $user) {
  if($user) {
    echo $user-&gt;getName();
  }
}</pre>
</div>
</div>
This could provide really convenient shortcuts, like only list unread mails, active users (see example), files in a directory that aren’t . or ..; just about everything!

The <code>__toBoolean()</code> would in some ways be equivalent to the method <a href="http://www.php.net/manual/en/class.filteriterator.php"><code>FilterIterator::accept()</code></a>.
