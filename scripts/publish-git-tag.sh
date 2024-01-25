#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

DIR="$(cd "$(dirname "$0")" && pwd)/.."

function publish {
  local version=$1

  echo "Publishing the branch (if required).."
  git push

  echo "Tagging as '$version'.."
  git tag "$version"

  echo "Pushing Tags.."
  git push --tags
}

function updateSdkReferenceAndCommitChanges {
  local directory=$1
  local tag=$2

  cd "${directory}"

  echo "Updating the dependency on 'github.com/hashicorp/go-azure-sdk/sdk'.."
  if [[ $(uname) == "Darwin" ]]; then
    echo "Using BSD sed"
    sed -i "" "s/github.com\/hashicorp\/go-azure-sdk\/sdk.*/github.com\/hashicorp\/go-azure-sdk\/sdk $tag/g" go.mod
  else
    sed -i "s/github.com\/hashicorp\/go-azure-sdk\/sdk.*/github.com\/hashicorp\/go-azure-sdk\/sdk $tag/g" go.mod
  fi

  echo "Running 'go mod tidy'.."
  go mod tidy

  echo "Running 'go mod vendor'.."
  go mod vendor

  echo "Adding the updated go.mod/go.sum files.."
  git add --all

  echo "Committing the changes.."
  git commit -m "$directory: updating to '$tag' of 'github.com/hashicorp/go-azure-sdk/sdk'"

  cd "${DIR}"
}

function main {
  local version=$1

  local microsoftGraphTag="microsoft-graph/$version"
  local resourceManagerTag="resource-manager/$version"
  local sdkTag="sdk/$version"

  echo "Releasing the Base Layer.."
  publish "$sdkTag"

  echo "Releasing the Microsoft Graph SDK.."
  updateSdkReferenceAndCommitChanges "microsoft-graph" "$version"
  publish "$microsoftGraphTag"

  echo "Releasing the Resource Manager SDK.."
  updateSdkReferenceAndCommitChanges "resource-manager" "$version"
  publish "$resourceManagerTag"
}

main "$1"
