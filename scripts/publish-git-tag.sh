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
  local version=$1

  echo "Configuring Git author information.."
  git config --global user.name "hc-github-team-tf-azure"
  git config --global user.email '<>'

  echo "Tagging the 'sdk' module.."
  local sdkTag="sdk/$version"
  publish "$version"

  echo "Tagging the 'microsoft-graph' module.."
  updateSdkReferenceThenPublish "microsoft-graph" "$version"

  echo "Tagging the 'resource-manager' module.."
  updateSdkReferenceThenPublish "resource-manager" "$version"
}

main "$1"