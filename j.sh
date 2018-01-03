#!/bin/sh
JEKYLL_VERSION=3.5

# use with args like
# jekyll new my_site
# jekyll build
# jekyll serve

docker run --rm \
  --volume="$PWD:/srv/jekyll" \
  -p 4000:4000 \
  -it jekyll/jekyll:$JEKYLL_VERSION \
  $@
