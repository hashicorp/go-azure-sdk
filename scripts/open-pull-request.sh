# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

function main {
  local prTitle="$1"
  local prBody="$2"
  local branch="$3"
  local label="$4"

  # this runs everytime the PR gets pushed too, whilst you can only create a PR a single time
  # so we should be smarter, but just trying to trigger this should be sufficient for now
  result=$(gh pr create --title "$prTitle" --body "$prBody" -B "$branch" -l "$label")
  echo "$result"
  exit 0
}

main "$1" "$2" "$3" "$4"