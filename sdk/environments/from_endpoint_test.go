// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environments

import (
	"context"
	"testing"
)

func TestFromEndpoint_Public(t *testing.T) {
	actual, err := FromEndpoint(context.Background(), "https://management.azure.com", "")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != "AzureCloud" {
		t.Fatalf("expected the Environment name to be `AzureCloud` but got %q", actual.Name)
	}
}

func TestFromEndpoint_USGovernment(t *testing.T) {
	t.Skip("Skipping because USGovernment ARM metadata service has reverted to older schema 2019-05-01 (manicminer, 2023-02-02)")

	actual, err := FromEndpoint(context.Background(), "https://management.usgovcloudapi.net", "AzureUSGovernment")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != "AzureCloud" {
		t.Fatalf("expected the Environment name to be `AzureCloud` but got %q", actual.Name)
	}
}
