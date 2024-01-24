#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

DIR="$(cd "$(dirname "$0")" && pwd)/.."

function publish {
  local version=$1

  echo "Tagging as '$version'.."
  git tag "$version"

  echo "Pushing Tags.."
  git push --tags
}

function main {
  local version=$1

  local sdkTag="sdk/$version"

  echo "Releasing the Base Layer.."
  publish "$sdkTag"
}

main "$1"
