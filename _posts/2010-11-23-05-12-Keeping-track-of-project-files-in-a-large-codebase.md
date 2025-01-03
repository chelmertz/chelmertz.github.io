---
date: "2010-11-23 05:12"
title: "Keeping track of project files in a large codebase"
permalink: "keeping-track-of-project-files-in-a-large-codebase"
tags: [thoughts]
---

Insight: symlinks should really be abused!

Let’s say you’re in the same position as me, working within a large codebase that requires _focus over multiple folders_ even for working with just one project. Perhaps you need to add logic, content, conditions, markup and css. These areas should all be separated, so naturally you’ll be clicking/`cd`-ing a lot… you see where this is going, right?

Make a temporary, meaningless folder named YourProject2010 and start pulling symlinks into it so you’ll never have to leave home again :) That would mean something like this:

```shell
cd /var/www/your-dev/YourProject2010
ln -s /var/www/your-dev/css/campaign_file.css .
ln -s /var/www/your-dev/js/campaign_file.js .
ln -s /var/www/your-dev/lib/campaign_file.php .
```

When you’re all done with the project, merging your branch with master and

```shell
rm -r /var/www/your-dev/YourProject2010
```

are the only two steps you need wipe out your traces and tell your boss that “it was a hard project to pull off because of its scope”.
