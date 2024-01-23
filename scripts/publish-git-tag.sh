#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

function publish {
  local version=$1

  echo "Tagging as '$version'.."
  git tag "$version"

  echo "Pushing Tags.."
  git push --tags
}

function main {
  local version=$1

  publish "$version"
}

main "$1"
