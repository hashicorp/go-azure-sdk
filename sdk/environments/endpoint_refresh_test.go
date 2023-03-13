// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environments

import (
	"context"
	"testing"
	"time"
)

func TestEndpointRefresh_China(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := AzureChina().RefreshMetaDataFromEndpoint(ctx); err != nil {
		t.Fatalf("refreshing metadata from endpoint for Azure China: %+v", err)
	}
}

func TestEndpointRefresh_Public(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := AzurePublic().RefreshMetaDataFromEndpoint(ctx); err != nil {
		t.Fatalf("refreshing metadata from endpoint for Azure Public: %+v", err)
	}
}

func TestEndpointRefresh_USGovernment(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := AzureUSGovernment().RefreshMetaDataFromEndpoint(ctx); err != nil {
		t.Fatalf("refreshing metadata from endpoint for Azure US Government: %+v", err)
	}
}
