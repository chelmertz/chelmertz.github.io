---
date: "2010-11-29 01:23"
layout: post
title: "Nice ORDER BY-syntax: FIELD()"
permalink: "nice-order-by-syntax-field"
tags: [mysql]
---

SELECT
turtle,
pizza,
rating
FROM
heroes
ORDER BY
FIELD(turtle, ‘Donatello’, ‘Michelangelo’, ‘Raphael’, ‘Leonardo’)

This let’s you get the ratings of Donatello’s first, etc. Really practical.
