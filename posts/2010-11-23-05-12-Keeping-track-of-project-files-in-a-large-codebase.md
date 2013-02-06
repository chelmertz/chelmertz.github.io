---
layout: post
title: "Keeping track of project files in a large codebase"
permalink: "/keeping-track-of-project-files-in-a-large-codebase"
categories: [thoughts]
---

Insight: symlinks should really be abused!

Let’s say you’re in the same position as me, working within a large codebase that requires <strong>focus over multiple folders</strong>* even for working with just one project. Perhaps you need to add logic, content, conditions, markup and css. These areas should all be separated, so naturally you’ll be clicking/<code>cd</code>-ing a lot… you see where this is going, right?

Make a temporary, meaningless folder named YourProject2010 and start pulling symlinks into it so you’ll never have to leave home again :) That would mean something like this:
<div class="CodeRay">
<div class="code">
<pre>cd /var/www/your-dev/YourProject2010
ln -s /var/www/your-dev/css/campaign_file.css .
ln -s /var/www/your-dev/js/campaign_file.js .
ln -s /var/www/your-dev/lib/campaign_file.php .</pre>
</div>
</div>
When you’re all done with the project, merging your branch with master and
<div class="CodeRay">
<div class="code">
<pre>rm -r /var/www/your-dev/YourProject2010</pre>
</div>
</div>
are the only two steps you need wipe out your traces and tell your boss that “it was a hard project to pull off because of its scope”.
