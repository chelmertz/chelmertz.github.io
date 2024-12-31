---
date: "2011-08-19 21:07"
title: "Handling forgotten passwords"
permalink: "handling-forgotten-passwords"
tags: [thoughts]
---

This is how I would let a user login again, after she lost her password:

- Provide a clear "Lost your password?" link at the login page and possibly in the "help section"
- Let your user provide one thing she surely remembers - her email
- Send a confirmation mail to the user with a text like
    > Hi, this is _you_;. Someone requested a password reset on _your site_. If that was not you, you can safely ignore this email and nothing will happen.
	> 
	> If it is you, who requested your password to be reset, click this link _yoursite.com/forgotten/3ttnjwnt32t_ to reset it.

The "3ttnjwnt32t" string is only mentioned in the letter, for a single user. You store it temporarily and use it to get the user.

Password resets:

- request reset
- allow anybody to send a reset-request-email
- arrive at hidden page that lasts 1 hour
- ask for a new password and the email

This also implies that an admin never should be able to change a password, only send out an email to the user who in turn changes it herself.
