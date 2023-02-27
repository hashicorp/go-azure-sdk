// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type options struct{}

func (o options) ToHeaders() *client.Headers {
	return nil
}

func (o options) ToOData() *odata.Query {
	return nil
}

func (o options) ToQuery() *client.QueryParams {
	return nil
}

func TestNewGetRequest(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
	defer cancel()

	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options{},
		Path:          "/",
	}
	localApi := environments.NewApiEndpoint("Example", "http://localhost", nil)
	resourceManagerClient, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
	if err != nil {
		t.Fatalf("building client: %+v", err)
	}
	if _, err = resourceManagerClient.NewRequest(ctx, opts); err != nil {
		t.Fatalf("building new request: %+v", err)
	}
}
