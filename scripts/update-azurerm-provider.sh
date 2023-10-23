#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

function prepare {
  local workingDirectory=$1

  echo "Recreating the working directory at '${workingDirectory}'.."
  rm -rf "${workingDirectory}"
  mkdir -p "${workingDirectory}"

  local repositoryDirectory="terraform-provider-azurerm"

  echo "Cloning AzureRM.."
  pushd "${workingDirectory}"
  git clone git@github.com:hashicorp/terraform-provider-azurerm.git "${repositoryDirectory}"

  echo "Configuring Git in that repository.."
  cd "${repositoryDirectory}"
  git config user.name "hc-github-team-tf-azure"
  git config user.email "<>"

  echo "Returning to the original directory.."
  popd
}

function runUpdaterTool {
  local workingDirectory=$1
  local newSdkVersion=$2
  local branch="auto-deps-pr/updating-go-azure-sdk-to-${newSdkVersion}"

  echo "Moving into the AzureRM Provider directory.."
  pwd
  pushd "${workingDirectory}/terraform-provider-azurerm"

  echo "Checking out a new branch.."
  git checkout -b "${branch}"

  echo "Building the updater tool.."
  cd ./internal/tools/update-go-azure-sdk
  go build .

  echo "Running the updater tool.."
  ./update-go-azure-sdk --new-sdk-version="${newSdkVersion}" --azurerm-repo-path=../../../ --go-sdk-repo-path=../../../../../ --output-file="../../../../pr-description.txt"

  echo "Pushing the branch"
  git push origin "$branch" -f

  echo "Returning to the original directory"
  popd
}

function main {
  local workingDirectory="./tmp"
  local goAzureSdkCheckout="$(pwd)"
  local newSdkVersion=$1

  prepare "$workingDirectory"
  runUpdaterTool "$workingDirectory" "$newSdkVersion"

  exit 0
}

# 1 = SDK Version
main "$1"
