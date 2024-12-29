---
date: "2010-12-13 01:43"
title: "Trying out closure"
permalink: "trying-out-closure"
tags: [javascript, server]
---

Very easy way of minimizing a few bytes of every request using Google’s <em>Closure Compiler</em> which is a java program for compressing javascript.

```
$ curl -0 http://closure-compiler.googlecode.com/files/compiler-latest.zip
$ unzip compiler-latest.zip -d ~/Applications/closure-compiler ; cd ~/Applications/closure-compiler
$ echo "// A simple function.
function hello(longName) {
  alert('Hello, ' + longName);
}
hello('New User');" &gt; hello.js
$ java -jar compiler.jar --js hello.js --js_output_file hello-compiled.js
$ cat hello-compiled.js
function hello(a){alert("Hello, "+a)}hello("New User");
$ wc -c hello.js
101 hello.js
$ wc -c hello-compiled.js
56 hello-compiled.js
```

<h2>Rinse and repeat</h2>

```
$ curl -0 http://closure-compiler.googlecode.com/files/compiler-latest.zip
```

Download the good stuff

```
$ unzip compiler-latest.zip -d ~/Applications/closure-compiler ; cd ~/Applications/closure-compiler
```

Locate the executable in a standard location and go there

```
$ echo "// A simple function.
function hello(longName) {
  alert('Hello, ' + longName);
}
hello('New User');" &gt; hello.js
```

Create a file with valid javascript in it

```
$ java -jar compiler.jar --js hello.js --js_output_file hello-compiled.js
```

Use the recently unzipped <em>compiler.jar</em> to take input from <em>hello.js</em> and generate its optimized version at <em>hello-compiled.js</em>. <em>Note that there’s only one file to use, compiler.jar. Usable.</em>

```
$ cat hello-compiled.js
function hello(a){alert("Hello, "+a)}hello("New User");
```

Dump the contents of the newly created file, just to verify what a minified javascript file could look like

```
$ wc -c hello.js
101 hello.js
$ wc -c hello-compiled.js
56 hello-compiled.js
```

Comparing bytes in each file shows almost 50% winnings, both in terms of lowering the storage needed but also for each request of that file.
