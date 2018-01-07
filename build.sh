#!/bin/sh

sh j.sh bundle install
exec sh j.sh jekyll build -s my_site
