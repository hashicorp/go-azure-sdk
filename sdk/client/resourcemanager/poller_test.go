// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

func TestNewPoller_LongRunningOperation(t *testing.T) {
	testData := []struct {
		response *client.Response
		valid    bool
	}{
		{
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Location"): []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789"},
					},
					Request: &http.Request{
						URL: &url.URL{
							Path: "/example",
						},
					},
				},
			},
			valid: true,
		},
		{
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusAccepted,
					Header: http.Header{
						http.CanonicalHeaderKey("Location"): []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789"},
					},
					Request: &http.Request{
						URL: &url.URL{
							Path: "/example",
						},
					},
				},
			},
			valid: true,
		},
		{
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Location"): []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789"},
					},
					Request: &http.Request{
						URL: &url.URL{
							Path: "/example",
						},
					},
				},
			},
			valid: true,
		},
		{
			// no LRO header, ergo this isn't acceptable
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusAccepted,
					Header:     http.Header{},
					Request: &http.Request{
						URL: &url.URL{
							Path: "/example",
						},
					},
				},
			},
			valid: false,
		},
	}

	for _, v := range testData {
		localApi := environments.NewApiEndpoint("Example", "https://async-url-test.local", nil)
		client, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("building client: %+v", err)
		}
		_, err = resourcemanager.PollerFromResponse(v.response, client)
		if v.valid {
			if err != nil {
				t.Fatal(fmt.Errorf("building poller from response: %+v", err))
			}
		} else {
			if err == nil {
				t.Fatalf("expected an error but didn't get one")
			}
		}
	}
}

func TestNewPoller_LongRunningOperationWithSelfReference(t *testing.T) {
	testData := []struct {
		response *client.Response
		valid    bool
	}{
		{
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
						http.CanonicalHeaderKey("Location"):     []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789?api-version=2020-01-01"},
					},
					Request: &http.Request{
						Method: http.MethodPost,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
			valid: true,
		},
		{
			// Seen in API Management - API and API Schema
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusAccepted,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
						http.CanonicalHeaderKey("Location"):     []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789?api-version=2020-01-01&asyncId=abc123&asyncCode=201"},
					},
					Request: &http.Request{
						Method: http.MethodPut,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
			valid: true,
		},
		{
			// Seen in API Management - API and API Schema
			response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
						http.CanonicalHeaderKey("Location"):     []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789?api-version=2020-01-01&asyncId=abc123&asyncCode=201"},
					},
					Request: &http.Request{
						Method: http.MethodPut,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
			valid: true,
		},
	}

	for _, v := range testData {
		localApi := environments.NewApiEndpoint("Example", "https://async-url-test.local", nil)
		client, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("building client: %+v", err)
		}
		_, err = resourcemanager.PollerFromResponse(v.response, client)
		if v.valid {
			if err != nil {
				t.Fatal(fmt.Errorf("building poller from response: %+v", err))
			}
		} else {
			if err == nil {
				t.Fatalf("expected an error but didn't get one")
			}
		}
	}
}

func TestNewPoller_ProvisioningState_ValidResourceManagerResource(t *testing.T) {
	// validates that a Resource Manager Resource is valid
	// e.g. `/some/resource`
	testCases := []struct {
		Response client.Response
	}{
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPost,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPut,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPost,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPut,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
	}

	for _, v := range testCases {
		localApi := environments.NewApiEndpoint("Example", "https://async-url-test.local", nil)
		client, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("building client: %+v", err)
		}

		if _, err = resourcemanager.PollerFromResponse(&v.Response, client); err != nil {
			t.Fatal(fmt.Errorf("building poller from response: %+v", err))
		}
	}
}

func TestNewPoller_ProvisioningState_ValidResourceManagerResourceOperation(t *testing.T) {
	// validates that an Operation on a Resource Manager Resource is valid
	// e.g. `/some/resource/start`
	testCases := []struct {
		Response client.Response
	}{
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPost,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPut,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPost,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: client.Response{
				Response: &http.Response{
					StatusCode: http.StatusCreated,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPut,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
	}

	for _, v := range testCases {
		localApi := environments.NewApiEndpoint("Example", "https://async-url-test.local", nil)
		client, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("building client: %+v", err)
		}

		if _, err = resourcemanager.PollerFromResponse(&v.Response, client); err != nil {
			t.Fatal(fmt.Errorf("building poller from response: %+v", err))
		}
	}
}

func TestNewPoller_ProvisioningState_Invalid(t *testing.T) {
	testCases := []struct {
		Response *client.Response
	}{
		{
			// Dropped response should raise an error
			Response: nil,
		},
		{
			Response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{
						// no content type header
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						// different content types
						http.CanonicalHeaderKey("Content-Type"): []string{"application/xml"},
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							u, _ := url.Parse("https://async-url-test.local/example/uri/start?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
		{
			Response: &client.Response{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
					},
					Request: &http.Request{
						Method: http.MethodPatch,
						URL: func() *url.URL {
							// non-resource manager operations aren't valid
							u, _ := url.Parse("https://async-url-test.local/example?api-version=2020-01-01")
							return u
						}(),
					},
				},
			},
		},
	}

	for i, v := range testCases {
		t.Logf("iteration %d", i)
		localApi := environments.NewApiEndpoint("Example", "https://async-url-test.local", nil)
		client, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("building client: %+v", err)
		}

		if _, err = resourcemanager.PollerFromResponse(v.Response, client); err == nil {
			t.Fatal("expected an error but didn't get one")
		}
	}
}
