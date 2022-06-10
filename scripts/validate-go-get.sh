#!/bin/bash

set -e

mkdir -p ./tmp/go-get
cd ./tmp/go-get
echo 'package main

import _ "github.com/hashicorp/go-azure-sdk/resource-manager/aadb2c/2021-04-01-preview/tenants"

func main() {
}
' > main.go
echo "module github.com/some/fake-repo

replace github.com/hashicorp/go-azure-sdk => ../../

go 1.18
  " > go.mod
go get github.com/hashicorp/go-azure-sdk@latest

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