---
permalink: "jenkins-python-virtualenv"
layout: post
title: "Making Jenkins work with python's virtualenv"
date: 2013-03-31 19:08
tags: ci python
comments: true
---

If you use [virtualenv](http://www.virtualenv.org/en/latest/) to isolate your python project's environment, and want your code tested automatically -- read on!

## virtualenv isolates your project's python environment

virtualenv makes sure you lockdown your project's main directory and all subdirectories of it. This 'lockdown' means that you never touches your global python binary, or any globally installed libraries (like "sudo pip install flask" ).

Once locked down, you install all packages again, even those you have globally installed. This enables you to have one version of flask globally installed, but another version in your project. All dependencies can be listed in a separate file and validate a precise environment for you to work with. Tightly controlled dependencies is key to a deployment without surprises.

## Jenkins checks the health of your project for each change

Jenkins is a [CI server](https://en.wikipedia.org/wiki/Continuous_integration#Principles) which means it does a lot of repeating stuff so you can focus on doing more important stuff. More specifically, it listens for changes to your project's version control system (like git).

When changes are detected, the project is built and the test suite is executed. If any step fails, the CI server tells you that it did.

## Setup Jenkins, and make it use virtualenv

Jenkins needs some massaging before it handles the hijacked environment of virtualenv. This is how I did it for my local git repository:

 - Download [Jenkins](http://jenkins-ci.org/)
 - Start it, it should be up on http://localhost:8080
 - Add the [Git Plugin](https://wiki.jenkins-ci.org/display/JENKINS/Git+Plugin)
 - Setup a new project with these properties:
    - Source Code Management: add the URI to your local repository, /Users/you/Sites/asdf in my case. Make sure the jenkins user can read this directory, otherwise the Jenkins GUI will tell you something random about invalid git repo, without a hint about a permissions error.
    - Build Triggers: Poll SCM (with an interval like 0 * * * *). This is needed because
      - you're too lazy to build manually; and
      - you can not trigger builds with a git post-commit hook otherwise
    - Build > Execute shell. I've used two steps, one for setting up the environment and one for the actual tests:

```sh
# Setup a proper path, I call my virtualenv dir "venv" and
# I've got the virtualenv command installed in /usr/local/bin
PATH=$WORKSPACE/venv/bin:/usr/local/bin:$PATH
if [ ! -d "venv" ]; then
        virtualenv venv
fi
. venv/bin/activate
pip install -r requirements.txt --download-cache=/tmp/$JOB_NAME
```
        
and
        
```sh
. venv/bin/activate
python test_app.py
```

## Trigger Jenkins through git hook

You need to add a git hook which triggers a Jenkins build:

```sh
echo "curl http://localhost:8080/git/notifyCommit?url=/Users/you/Sites/asdf" >> .git/hooks/post-commit
chmod +x !$
```

Add an erroneous test, this will do:

```python
def test_bad(self):
    assert False
```

You should see a new build being queued up in Jenkins within a minute. If that doesn't work, execute the hook and watch the output for error messages:

```sh
./git/hooks/post-commit
```

Now, Jenkins should try to test your project but fail, and report the failure through the GUI. Tada.
