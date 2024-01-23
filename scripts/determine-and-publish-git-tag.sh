#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

DIR="$(cd "$(dirname "$0")" && pwd)/.."

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
  #
  # turns out that Go doesn't like the minor/patch to start with a 0, else
  # the Go Module Proxy returns:
  # > bad request: version "v0.20220622.050833" is not canonical (wanted "")
  # as such we'll prefix the Hours-Minutes-Seconds segment with a `1` for now
  date '+v0.%Y%m%d.1%H%M%S'
}

function publish {
  local version=$1

  echo "Tagging as '$version'.."
  git tag "$version"

  echo "Pushing Tags.."
  git push --tags
}

function updateSdkReferenceThenPublish {
  local directory=$1
  local tag=$2

  cd "${directory}"

  echo "Checking out a working branch from 'sdk/$tag'.."
  git checkout -b "$directory/$tag" "sdk/$tag"

  echo "Updating the dependency on 'github.com/hashicorp/go-azure-sdk/sdk'.."
  go get "github.com/hashicorp/go-azure-sdk/sdk@$sdkTag"

  echo "Running 'go mod tidy'.."
  go mod tidy

  echo "Running 'go mod vendor'.."
  go mod vendor

  echo "Adding the updated go.mod/go.sum files.."
  git add --all

  echo "Committing the changes.."
  git commit -m "$directory: updating to '$sdkTag' of 'github.com/hashicorp/go-azure-sdk/sdk'"

  cd "${DIR}"
}

function main {
  local gitTag
  gitTag=$(determineGitTag)

  echo "Configuring Git author information.."
  git config --global user.name "hc-github-team-tf-azure"
  git config --global user.email '<>'

  echo "Tagging the 'sdk' module.."
  local sdkTag="sdk/$gitTag"
  publish "$sdkTag"

  echo "Tagging the 'microsoft-graph' module.."
  local resourceManagerTag="microsoft-graph/$gitTag"
  updateSdkReferenceThenPublish "microsoft-graph" "$gitTag"

  echo "Tagging the 'resource-manager' module.."
  local resourceManagerTag="resource-manager/$gitTag"
  updateSdkReferenceThenPublish "resource-manager" "$gitTag"
}

main
