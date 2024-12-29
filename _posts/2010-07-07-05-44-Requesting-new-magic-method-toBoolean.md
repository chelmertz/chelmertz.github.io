---
date: "2010-07-07 05:44"
title: "Requesting new magic method: `__toBoolean()`"
permalink: "requesting-new-magic-method-__toboolean"
tags: [php]
---

What if you only want to display your active users in a very simple way? Consider the following:

```
class User {
  public $active = true;
  private $_name;
  public function __construct($name) {
    $this->_name = (string) $name;
  }
  public function __toBoolean() {
    return (bool) $this->active;
  }
  public function getName() {
    return $this->_name;
  }
}

$emily = new User('Emily');
$emily->active = false;
$sara = new User('Sara');

$users = array($emily, $sara);
foreach($users as $user) {
  if($user) {
    echo $user->getName();
  }
}
```

This could provide really convenient shortcuts, like only list unread mails, active users (see example), files in a directory that aren’t . or ..; just about everything!

The `__toBoolean()` would in some ways be equivalent to the method [`FilterIterator::accept()`](http://www.php.net/manual/en/class.filteriterator.php).
