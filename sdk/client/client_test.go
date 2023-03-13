// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccClient(t *testing.T) {
	test.AccTest(t)

	ctx := context.TODO()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.MicrosoftGraph
	endpoint, ok := api.Endpoint()
	if !ok {
		t.Fatalf("missing endpoint for microsoft graph for this environment")
	}
	conn.Authorize(ctx, t, api)

	c := NewClient(*endpoint, "example", "2020-01-01")
	c.Authorizer = conn.Authorizer

	path := fmt.Sprintf("/v1.0/servicePrincipals/%s", conn.Claims.ObjectId)
	reqOpts := RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: nil,
		Path:          path,
	}
	req, err := c.NewRequest(ctx, reqOpts)
	if err != nil {
		t.Fatal(err)
	}

	_, err = req.ExecutePaged(ctx)
	if err != nil {
		t.Fatalf("ExecutePaged(): %v", err)
	}
}
