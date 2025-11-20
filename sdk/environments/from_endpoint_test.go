// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environments

import (
	"context"
	"testing"
)

func TestFromEndpoint_Public(t *testing.T) {
	actual, err := FromEndpoint(context.Background(), "https://management.azure.com")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != "AzureCloud" {
		t.Fatalf("expected the Environment name to be `AzureCloud` but got %q", actual.Name)
	}

	endpoint, ok := actual.ResourceManager.Endpoint()
	if !ok {
		t.Fatalf("refreshing MetaData from Endpoint: no `ResourceManager` endpoint was defined")
	}
	if *endpoint != "https://management.azure.com" {
		t.Fatalf("expected the ResourceManager endpoint to be `https://management.azure.com` but got %q", *endpoint)
	}
}

func TestFromEndpoint_USGovernment(t *testing.T) {
	actual, err := FromEndpoint(context.Background(), "https://management.usgovcloudapi.net")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != "AzureUSGovernment" {
		t.Fatalf("expected the Environment name to be `AzureUSGovernment` but got %q", actual.Name)
	}

	endpoint, ok := actual.ResourceManager.Endpoint()
	if !ok {
		t.Fatalf("refreshing MetaData from Endpoint: no `ResourceManager` endpoint was defined")
	}
	if *endpoint != "https://management.usgovcloudapi.net" {
		t.Fatalf("expected the ResourceManager endpoint to be `https://management.usgovcloudapi.net` but got %q", *endpoint)
	}
}
