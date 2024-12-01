---
layout: post
title: "Trying out closure"
permalink: "trying-out-closure"
tags: [javascript, server]
---

Very easy way of minimizing a few bytes of every request using Google’s <em>Closure Compiler</em> which is a java program for compressing javascript.
<div class="CodeRay">
<div class="code">
<pre>$ curl -0 http://closure-compiler.googlecode.com/files/compiler-latest.zip
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
56 hello-compiled.js</pre>
</div>
</div>
<h2>Rinse and repeat</h2>
<div class="CodeRay">
<div class="code">
<pre>$ curl -0 http://closure-compiler.googlecode.com/files/compiler-latest.zip</pre>
</div>
</div>
Download the good stuff
<div class="CodeRay">
<div class="code">
<pre>$ unzip compiler-latest.zip -d ~/Applications/closure-compiler ; cd ~/Applications/closure-compiler</pre>
</div>
</div>
Locate the executable in a standard location and go there
<div class="CodeRay">
<div class="code">
<pre>$ echo "// A simple function.
function hello(longName) {
  alert('Hello, ' + longName);
}
hello('New User');" &gt; hello.js</pre>
</div>
</div>
Create a file with valid javascript in it
<div class="CodeRay">
<div class="code">
<pre>$ java -jar compiler.jar --js hello.js --js_output_file hello-compiled.js</pre>
</div>
</div>
Use the recently unzipped <em>compiler.jar</em> to take input from <em>hello.js</em> and generate its optimized version at <em>hello-compiled.js</em>. <em>Note that there’s only one file to use, compiler.jar. Usable.</em>
<div class="CodeRay">
<div class="code">
<pre>$ cat hello-compiled.js
function hello(a){alert("Hello, "+a)}hello("New User");</pre>
</div>
</div>
Dump the contents of the newly created file, just to verify what a minified javascript file could look like
<div class="CodeRay">
<div class="code">
<pre>$ wc -c hello.js
101 hello.js
$ wc -c hello-compiled.js
56 hello-compiled.js</pre>
</div>
</div>
Comparing bytes in each file shows almost 50% winnings, both in terms of lowering the storage needed but also for each request of that file.
