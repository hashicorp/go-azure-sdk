#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


set -e

mkdir -p ./tmp/go-get
cd ./tmp/go-get
echo 'package main

import _ "github.com/hashicorp/go-azure-sdk/resource-manager/aadb2c/2021-04-01-preview/tenants"
import _ "github.com/hashicorp/go-azure-sdk/sdk/environments"

func main() {
}
' > main.go
echo "module github.com/some/fake-repo

replace github.com/hashicorp/go-azure-sdk/resource-manager => ../../resource-manager
replace github.com/hashicorp/go-azure-sdk/sdk => ../../sdk

go 1.21
" > go.mod

# Update the Base Layer
go get github.com/hashicorp/go-azure-sdk/sdk@latest

# Update the Resource Manager SDK
go get github.com/hashicorp/go-azure-sdk/resource-manager@latest

if [[ ! $(go mod tidy) -eq 0 ]]; then
  echo "Go Mod Tidy failed"
  exit 1
fi

if [[ ! $(go mod vendor) -eq 0 ]]; then
  echo "Go Mod vendor failed"
  exit 1
fi

cd ../../
rm -rf ./tmp