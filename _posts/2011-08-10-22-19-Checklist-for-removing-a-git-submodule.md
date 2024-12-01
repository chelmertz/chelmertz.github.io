---
layout: post
title: "Checklist for removing a git submodule"
permalink: "checklist-removing-git-submodule"
tags: [git]
---

<blockquote cite="https://git.wiki.kernel.org/index.php/GitSubmoduleTutorial">
<ol>
	<li>Delete the relevant line from the <em>.gitmodules</em> file.</li>
	<li>Delete the relevant section from <em>.git/config</em>.</li>
	<li>Run <em>git rm --cached path_to_submodule</em> (no trailing slash).</li>
	<li>Commit the superproject.</li>
	<li>Delete the now untracked submodule files.</li>
</ol>
</blockquote>
From <a href="https://git.wiki.kernel.org/index.php/GitSubmoduleTutorial">https://git.wiki.kernel.org/index.php/GitSubmoduleTutorial</a>
