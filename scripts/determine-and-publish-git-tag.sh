#!/bin/bash

set -e

function determineGitTag {
  # Whilst we could do something fancy here to determine the version, tbqh
  # it doesn't matter providing it's:
  # a) 0-based (so we don't bump the Go module version/trigger import path changes)
  # b) ordered
  # c) clear when it was released
  #
  # whilst an odd choice for an SDK, makes sense insofar as the API Versions
  # themselves are versioned too - so we can use a date/time stamp for now:
  # e.g.v0.YYYYMMDD.HHMMSS (v0.20220504.115712)
  date '+v0.%Y%m%d.%H%M%S'
}

function publish {
  local version=$1

  echo "Tagging as '$version'.."
  git tag "$version"

  echo "Pushing Tags.."
  git push --tags
}

function main {
  local gitTag
  gitTag=$(determineGitTag)
  publish "$gitTag"
}

main
