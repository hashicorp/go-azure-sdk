#!/bin/bash

# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

function main {
  local prDescriptionFileName="$1"

  base64dValue=$(base64 -i "$prDescriptionFileName")
  output='description=<<EOF'"$base64dValue"'EOF'
  echo "$output"
  exit 0
}

main "$1"