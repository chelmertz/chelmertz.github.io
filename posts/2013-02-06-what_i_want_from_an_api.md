---
layout: post
title: "What I want from an interface"
date: 2013-02-06 21:27
categories: thoughts
---

Consider this a checklist for your API, should work for REST, CLI or source code!

## What the interface is for
Have a scope and tell me what it is in a couple of sentences.

## How to
Let me learn the general approach and edge cases via a concise and result oriented documentation.

## Feedback
Scream at me when I did something wrong and, if I was close enough (and you have a lot of time), point me in the right direction.

## Consistency
For example, when treating errors you shouldn't alternate between simultaneous ways of telling your users something went wrong. Pick one and go with it.

- error callbacks (JS)
- separate errors from regular output (bash)
- exceptions (python et al.)
- magical return codes (C)
