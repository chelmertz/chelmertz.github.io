#!/bin/sh

sh j.sh bundle install
sh j.sh jekyll clean -s my_site
exec sh j.sh jekyll build -s my_site
