// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"net/http"
	"regexp"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestSetAuthHeader(t *testing.T) {
	req := &http.Request{}
	authorizer := &test.TestAuthorizer{}

	err := auth.SetAuthHeader(context.Background(), req, authorizer)
	if err != nil {
		t.Fatalf("received error: %+v", err)
	}

	expected := regexp.MustCompile("^Bearer [a-zA-Z0-9_-]+[.][a-zA-Z0-9_-]+[.][a-zA-Z0-9_-]+")
	if val := req.Header.Get("Authorization"); !expected.MatchString(val) {
		t.Fatalf("Authorization header mismatch, received: %q", val)
	}
}
