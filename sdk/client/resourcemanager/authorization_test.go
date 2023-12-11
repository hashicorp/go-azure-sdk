// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager_test

import (
	"context"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAuthorizeResourceManagerRequest(t *testing.T) {
	req := &http.Request{}
	authorizer := &test.TestAuthorizer{}

	err := resourcemanager.AuthorizeResourceManagerRequest(context.Background(), req, authorizer)
	if err != nil {
		t.Fatalf("received error: %+v", err)
	}

	expected := regexp.MustCompile("^Bearer [a-zA-Z0-9_-]+[.][a-zA-Z0-9_-]+[.][a-zA-Z0-9_-]+")
	if val := req.Header.Get("Authorization"); !expected.MatchString(val) {
		t.Fatalf("Authorization header mismatch, received: %q", val)
	}
	auxVals := strings.Split(req.Header.Get("X-Ms-Authorization-Auxiliary"), ",")
	if len(auxVals) != 3 {
		t.Fatalf("X-Ms-Authorization-Auxiliary mismatch, should contain 3 tokens, received: %q", strings.Join(auxVals, ","))
	}
	for i, auxVal := range auxVals {
		if !expected.MatchString(strings.TrimSpace(auxVal)) {
			t.Fatalf("Authorization header mismatch for token %d, received: %q", i, auxVal)
		}
	}
}
