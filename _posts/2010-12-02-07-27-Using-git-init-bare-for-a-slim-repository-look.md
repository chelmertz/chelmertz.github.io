---
date: "2010-12-02 07:27"
title: "Using git init --bare for a slim repository look"
permalink: "using-git-init-bare-for-a-slim-repository-look"
tags: [git, server]
---

Until now, I’ve mainly worked alone with my git repositories. Sometime pushing to <a href="https://github.com/chelmertz">my github-account</a>, but always using the service as a kind of backup solution. Until now. At work, we’re setting up our new environment with git and I’m lucky enough to learn it while doing it.

This setting will differ quite a lot from what I’m used to:
<ul>
	<li>for private projects, the git repository is my project folder;</li>
	<li>for work, the repository has to be <strong>synced across multiple users</strong>, a base for <strong>deployment</strong> and <strong>act as a more important repository</strong> than our local ones, kind of what SVN was doing and what, I’m told, git isn’t really made for.</li>
</ul>
Also, exercising a plain <code>git init</code> routine creates a complete repository with all source files, branches &amp; stuff; which really are ment to keep track off and develop in. Let’s not, let’s just keep the repository there and only <code>git push</code> there so <strong>all development will be done in each user’s local repository</strong>.

Let’s get to the good stuff:
<div class="CodeRay">
<div class="code">
<pre># create a user named 'git' and make that user the owner of your repositories, and nothing else
$ mkdir -p /home/git/repositories
$ cd /home/git/repositories
$ mkdir your-project.git &amp;&amp; cd your-project.git
$ git init --bare
$ git config receive.denyCurrentBranch "ignore"
$ mkdir -p ~/www/your-project &amp;&amp; cd ~/www/your-project
$ echo "lalala, random file" &gt; random.txt &amp;&amp; git add . &amp;&amp; git commit -m "First commit"
$ git remote add origin /home/git/repositories/your-project.git
$ git push origin master</pre>
</div>
</div>
There you go, a fully capable external repository, local on your dev machine, ready to export to a staging/testing/Q&amp;A-setting. For an easy setup of this, go read <strong><a href="http://toroid.org/ams/git-website-howto">Using Git to manage a web site</a></strong> which <em>worked out of the box</em> for us.

If you get any errors whilst doing this, google it or just try to parse it (for us, it mostly came down to <code>chown</code>- and <code>chmod</code>-ing the <em>-R /home/git/repositories/</em> folders, so using <code>sudo</code> helps a lot).
