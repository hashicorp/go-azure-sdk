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

	if actual.Name != AzurePublicCloud {
		t.Fatalf("expected the Environment name to be `%s` but got %q", AzurePublicCloud, actual.Name)
	}
}

func TestFromEndpoint_USGovernment(t *testing.T) {
	actual, err := FromEndpoint(context.Background(), "https://management.usgovcloudapi.net", "AzureUSGovernment")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != AzureUSGovernmentCloud {
		t.Fatalf("expected the Environment name to be `%s` but got %q", AzureUSGovernmentCloud, actual.Name)
	}
}

func TestFromEndpoint_China(t *testing.T) {
	actual, err := FromEndpoint(context.Background(), "https://management.chinacloudapi.cn", "AzureUSGovernment")
	if err != nil {
		t.Fatalf("loading from endpoint: %+v", err)
	}

	if actual.Name != AzureChinaCloud {
		t.Fatalf("expected the Environment name to be `%s` but got %q", AzureChinaCloud, actual.Name)
	}
}
