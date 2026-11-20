// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
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
	resourceManagerClient, err := resourcemanager.NewClient(localApi, "example", "2020-02-01")
	if err != nil {
		t.Fatalf("building client: %+v", err)
	}
	if _, err = resourceManagerClient.NewRequest(ctx, opts); err != nil {
		t.Fatalf("building new request: %+v", err)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestPutIdempotencyDefaultsAndOptOut(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(30*time.Second))
	defer cancel()

	localApi := environments.NewApiEndpoint("Example", "http://localhost", nil)

	runPut := func(disableClientDefault bool, reqIdempotent *bool) (int, error) {
		c, err := resourcemanager.NewClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("NewClient error: %v", err)
		}
		c.AuthorizeRequest = nil
		c.DisableIdempotentPutRetries = disableClientDefault

		attempts := 0
		c.SetTransport(roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			attempts++
			if attempts == 1 {
				return nil, io.EOF
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       io.NopCloser(bytes.NewBufferString(`{"ok":true}`)),
				Request:    req,
			}, nil
		}))

		req, err := c.NewRequest(ctx, client.RequestOptions{
			ContentType:         "application/json",
			ExpectedStatusCodes: []int{http.StatusOK},
			HttpMethod:          http.MethodPut,
			Idempotent:          reqIdempotent,
			Path:                "/resource",
		})
		if err != nil {
			t.Fatalf("NewRequest error: %v", err)
		}
		_, err = req.Execute(ctx)
		return attempts, err
	}

	// Default Resource Manager PUT retries on nil-response EOF
	if attempts, err := runPut(false, nil); err != nil || attempts != 2 {
		t.Fatalf("expected default PUT to retry and succeed in 2 attempts, got attempts=%d err=%v", attempts, err)
	}

	// Client-level opt-out disables PUT retry
	if attempts, err := runPut(true, nil); err == nil || attempts != 1 {
		t.Fatalf("expected DisableIdempotentPutRetries=true to fail after 1 attempt, got attempts=%d err=%v", attempts, err)
	}

	// Request-level opt-out disables PUT retry
	if attempts, err := runPut(false, pointer.To(false)); err == nil || attempts != 1 {
		t.Fatalf("expected RequestOptions.Idempotent=false to fail after 1 attempt, got attempts=%d err=%v", attempts, err)
	}
}
