#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

DIR="$(cd "$(dirname "$0")" && pwd)/.."

function setupGitAuth {
  if [[ -n "${RELEASE_DEPLOY_KEY:-}" ]]; then
    key_file=$(mktemp "${RUNNER_TEMP:-/tmp}/release_ssh_key.XXXXXX")
    chmod 600 "$key_file"
    trap 'rm -f "$key_file"' EXIT
    printf '%s\n' "$RELEASE_DEPLOY_KEY" > "$key_file"
    export GIT_SSH_COMMAND="ssh -i ${key_file} -o IdentitiesOnly=yes -o StrictHostKeyChecking=accept-new"
    git remote set-url origin git@github.com:hashicorp/go-azure-sdk.git
  fi
}

function publishTagOnly {
  local tag=$1

  echo "Tagging as '$tag'.."
  git tag "$tag"

  echo "Pushing Tag '$tag'.."
  git push origin "refs/tags/$tag"
}

function publishBranchAndTag {
  local tag=$1

  echo "Validating modified files before pushing to main.."
  local disallowed
  disallowed=$(git diff --name-only origin/main..HEAD | grep -Ev '^(microsoft-graph|resource-manager|data-plane)/(go\.mod|go\.sum|vendor/)' || true)
  if [[ -n "$disallowed" ]]; then
    echo "::error::Refusing to push release! Unexpected files modified outside go.mod/go.sum/vendor:"
    echo "$disallowed"
    exit 1
  fi

  echo "Publishing main branch.."
  git push origin HEAD:refs/heads/main

  echo "Tagging as '$tag'.."
  git tag "$tag"

  echo "Pushing Tag '$tag'.."
  git push origin "refs/tags/$tag"
}

function updateSdkReferenceAndCommitChanges {
  local directory=$1
  local tag=$2

  cd "${directory}"

  echo "Setting the committer information.."
  git config user.name "hc-github-team-tf-azure"
  git config user.email "<>"

  echo "Updating the dependency on 'github.com/hashicorp/go-azure-sdk/sdk'.."
  if [[ $(uname) == "Darwin" ]]; then
    echo "Using BSD sed"
    sed -i "" "s/github.com\/hashicorp\/go-azure-sdk\/sdk.*/github.com\/hashicorp\/go-azure-sdk\/sdk $tag/g" go.mod
  else
    sed -i "s/github.com\/hashicorp\/go-azure-sdk\/sdk.*/github.com\/hashicorp\/go-azure-sdk\/sdk $tag/g" go.mod
  fi

  echo "Running 'go mod tidy'.."
  GOPROXY="direct,https://proxy.golang.org" go mod tidy

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
  if [[ ! "$version" =~ ^v0\.[0-9]{8}\.1[0-9]{6}$ ]]; then
    echo "Invalid version format: $version" >&2
    exit 1
  fi

  local microsoftGraphTag="microsoft-graph/$version"
  local resourceManagerTag="resource-manager/$version"
  local dataPlaneTag="data-plane/$version"
  local sdkTag="sdk/$version"

  setupGitAuth

  echo "Releasing the Base Layer.."
  publishTagOnly "$sdkTag"

  echo "Releasing the Microsoft Graph SDK.."
  updateSdkReferenceAndCommitChanges "microsoft-graph" "$version"
  publishBranchAndTag "$microsoftGraphTag"

  echo "Releasing the Resource Manager SDK.."
  updateSdkReferenceAndCommitChanges "resource-manager" "$version"
  publishBranchAndTag "$resourceManagerTag"

  echo "Releasing the Data plane SDK.."
  updateSdkReferenceAndCommitChanges "data-plane" "$version"
  publishBranchAndTag "$dataPlaneTag"
}

main "$1"
