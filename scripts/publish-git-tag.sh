#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -euo pipefail

function verifyReleaseSha {
  local repo="${REPO:-hashicorp/go-azure-sdk}"
  local head_sha
  head_sha=$(git rev-parse HEAD)

  if [[ -n "${RELEASE_SHA:-}" ]]; then
    if [[ "$head_sha" != "$RELEASE_SHA" ]]; then
      echo "::error::Checked-out HEAD ($head_sha) does not match expected RELEASE_SHA ($RELEASE_SHA)." >&2
      exit 1
    fi

    local remote_main
    remote_main=$(git ls-remote "https://github.com/${repo}.git" refs/heads/main | cut -f1)
    if [[ "$remote_main" != "$RELEASE_SHA" ]]; then
      echo "::error::Remote main ($remote_main) has advanced past RELEASE_SHA ($RELEASE_SHA). Skipping tag creation so the newer run releases cleanly." >&2
      exit 1
    fi
  fi
}

function verifySubmoduleGoModVersion {
  local directory=$1
  local version=$2

  if ! grep -Eq "github\.com/hashicorp/go-azure-sdk/sdk[[:space:]]+${version}([[:space:]]|$)" "${directory}/go.mod"; then
    echo "::error::${directory}/go.mod does not require github.com/hashicorp/go-azure-sdk/sdk ${version}" >&2
    exit 1
  fi
}

function publishTag {
  local tag=$1
  local repo="${REPO:-hashicorp/go-azure-sdk}"

  echo "Tagging HEAD as '$tag'.."
  git tag "$tag"

  echo "Pushing Tag '$tag'.."
  if [[ -n "${GH_TOKEN:-}" ]]; then
    local auth_header
    auth_header=$(printf 'x-access-token:%s' "$GH_TOKEN" | base64 | tr -d '\n')
    git -c "http.https://github.com/.extraheader=AUTHORIZATION: basic ${auth_header}" \
      push "https://github.com/${repo}.git" "refs/tags/$tag"
  else
    git push origin "refs/tags/$tag"
  fi
}

function main {
  local mode=${1:-}
  local version=${2:-}

  if [[ "$mode" != "sdk" && "$mode" != "submodules" ]]; then
    echo "Usage: $0 <sdk|submodules> <v0.YYYYMMDD.1HHMMSS>" >&2
    exit 1
  fi

  if [[ ! "$version" =~ ^v0\.[0-9]{8}\.1[0-9]{6}$ ]]; then
    echo "Invalid version format: $version" >&2
    exit 1
  fi

  verifyReleaseSha

  if [[ "$mode" == "sdk" ]]; then
    echo "Releasing the Base Layer.."
    publishTag "sdk/$version"
  else
    echo "Verifying submodule go.mod files reference sdk $version.."
    verifySubmoduleGoModVersion "microsoft-graph" "$version"
    verifySubmoduleGoModVersion "resource-manager" "$version"
    verifySubmoduleGoModVersion "data-plane" "$version"

    echo "Releasing the Microsoft Graph SDK.."
    publishTag "microsoft-graph/$version"

    echo "Releasing the Resource Manager SDK.."
    publishTag "resource-manager/$version"

    echo "Releasing the Data Plane SDK.."
    publishTag "data-plane/$version"
  fi
}

main "$@"
