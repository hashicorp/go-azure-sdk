#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

DIR="$(cd "$(dirname "$0")" && pwd)/.."

function updateSubmodule {
  local directory=$1
  local version=$2

  cd "${DIR}/${directory}"

  echo "Updating '${directory}/go.mod' to github.com/hashicorp/go-azure-sdk/sdk@${version}.."
  go mod edit -require="github.com/hashicorp/go-azure-sdk/sdk@${version}"

  echo "Running 'go mod tidy' in '${directory}'.."
  GOPROXY="direct,https://proxy.golang.org" go mod tidy

  cd "${DIR}"

  if [[ $(git status --porcelain "${directory}/go.mod" "${directory}/go.sum" | wc -l) -gt 0 ]]; then
    git add "${directory}/go.mod" "${directory}/go.sum"
    git commit -m "${directory}: updating to '${version}' of 'github.com/hashicorp/go-azure-sdk/sdk'"
  fi
}

function main {
  local version=${1:-}
  if [[ ! "$version" =~ ^v0\.[0-9]{8}\.1[0-9]{6}$ ]]; then
    echo "Invalid version format: $version" >&2
    exit 1
  fi

  local branch="auto-release/${version}"

  cd "${DIR}"
  git checkout -b "$branch"
  git config user.name "hc-github-team-tf-azure"
  git config user.email "<>"

  updateSubmodule "microsoft-graph" "$version"
  updateSubmodule "resource-manager" "$version"
  updateSubmodule "data-plane" "$version"

  if [[ -n "${BUNDLE_PATH:-}" ]]; then
    echo "Creating git bundle at ${BUNDLE_PATH} for refs/heads/${branch}.."
    git bundle create "${BUNDLE_PATH}" "origin/main..refs/heads/${branch}"
  fi

  if [[ -n "${GITHUB_OUTPUT:-}" ]]; then
    echo "branch_name=${branch}" >> "$GITHUB_OUTPUT"
  fi
}

main "$@"
