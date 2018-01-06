---
layout: post
title: "Handling forgotten passwords"
permalink: "/handling-forgotten-passwords/"
tags: [thoughts]
---

This is how I would let a user login again, after she lost her password:
<ol>
	<li>Provide a clear "Lost your password?" link at the login page and possibly in the "help section"</li>
	<li>Let your user provide one thing she surely remembers - her email</li>
	<li>Send a confirmation mail to the user with a text like
<blockquote>Hi, this is &lt;you&gt;. Someone requested a password reset on &lt;your site&gt;. If that was not you, you can safely ignore this email and nothing will happen.

If it is you, who requested your password to be reset, click this link &lt;yoursite.com/forgotten/3ttnjwnt32t&gt; to reset it.</blockquote>
The "3ttnjwnt32t" string is only mentioned in the letter, you save</li>
</ol>
store encrypted with a hash

password resets:
<ul>
	<li>request reset</li>
	<li>allow anybody to send a reset-request-email</li>
	<li>arrive at hidden page that lasts 1 hour</li>
	<li>ask for a new password and the email</li>
</ul>
<div>This also implies that an admin never should be able to change a password, only send out an email to the user who in turn changes it herself.</div>
