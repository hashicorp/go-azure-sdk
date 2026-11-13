// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplane

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

func TestPollerProvisioningState_InProvisioningState_Immediate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// Ensure we're parsing from the `provisioningState` field
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	dataplaneClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since we should be using
		// `apiVersion` (which otherwise gets parsed from the URI in `provisioningStatePollerFromResponse`)
		Client:     client.NewClient(server.URL, "Example", "2015-01-01"),
		ApiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:           helper.expectedApiVersion,
		client:               dataplaneClient,
		initialRetryDuration: 10,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}
	actual, err := poller.Poll(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual.Status != pollers.PollingStatusSucceeded {
		t.Fatalf("expected %q but got %q", string(pollers.PollingStatusSucceeded), string(actual.Status))
	}
}

func TestPollerProvisioningState_InStatus_Immediate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// Ensure we're parsing from the `status` field
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		ApiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:           helper.expectedApiVersion,
		client:               resourceManagerClient,
		initialRetryDuration: 10,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}
	actual, err := poller.Poll(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual.Status != pollers.PollingStatusSucceeded {
		t.Fatalf("expected %q but got %q", string(pollers.PollingStatusSucceeded), string(actual.Status))
	}
}

func TestPollerProvisioningState_InProvisioningState_DroppedThenInProgressThenSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 90*time.Second)
	defer cancel()

	// changing where we're setting this for the heck of it
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInProvisioningState(statusInProgress),
		responseThatDropsTheConnection(),
		responseWithStatusInProvisioningState(statusInProgress),
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since we should be using
		// `apiVersion` (which otherwise gets parsed from the URI in `provisioningStatePollerFromResponse`)
		Client:     client.NewClient(server.URL, "Example", "2015-01-01"),
		ApiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:            helper.expectedApiVersion,
		client:                resourceManagerClient,
		initialRetryDuration:  10,
		originalUri:           "/provisioning-state/poll",
		resourcePath:          "/provisioning-state/poll",
		maxDroppedConnections: 3,
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // working on it
		pollers.PollingStatusUnknown,
		// NOTE: the Dropped Connection will be ignored/silently retried
		pollers.PollingStatusInProgress, // working on it
		pollers.PollingStatusSucceeded,  // good
	}
	for i, expected := range expectedStatuses {
		t.Logf("Poll %d..", i)
		result, err := poller.Poll(ctx)
		if expected == pollers.PollingStatusUnknown {
			// the dropped connection is reported as an error, for the caller to poll again
			if !errors.As(err, &pollers.PollingDroppedConnectionError{}) {
				t.Fatalf("expected a dropped connection error but got: %+v", err)
			}

			continue
		}
		if err != nil {
			t.Fatal(err.Error())
		}
		if result.Status != expected {
			t.Fatalf("expected status to be %q but got %q", expected, result.Status)
		}
	}
	helper.assertCalled(t, 4)
}

func TestPollerProvisioningState_InStatus_DroppedThenInProgressThenSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// changing where we're setting this for the heck of it
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInStatusField(statusInProgress),
		responseThatDropsTheConnection(),
		responseWithStatusInStatusField(statusInProgress),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since we should be using
		// `apiVersion` (which otherwise gets parsed from the URI in `provisioningStatePollerFromResponse`)
		Client:     client.NewClient(server.URL, "Example", "2015-01-01"),
		ApiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:            helper.expectedApiVersion,
		client:                resourceManagerClient,
		initialRetryDuration:  10,
		originalUri:           "/provisioning-state/poll",
		resourcePath:          "/provisioning-state/poll",
		maxDroppedConnections: 3,
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // working on it
		pollers.PollingStatusUnknown,
		pollers.PollingStatusInProgress, // working on it
		pollers.PollingStatusSucceeded,  // good
	}
	for i, expected := range expectedStatuses {
		t.Logf("Poll %d..", i)
		result, err := poller.Poll(ctx)
		if expected == pollers.PollingStatusUnknown {
			// the dropped connection is reported as an error, for the caller to poll again
			if !errors.As(err, &pollers.PollingDroppedConnectionError{}) {
				t.Fatalf("expected a dropped connection error but got: %+v", err)
			}

			continue
		}
		if err != nil {
			t.Fatal(err.Error())
		}
		if result.Status != expected {
			t.Fatalf("expected status to be %q but got %q", expected, result.Status)
		}
	}
	// sanity-checking - expect 4 calls but 3 statuses (since the dropped connection is silently retried)
	helper.assertCalled(t, 4)
}

func TestPollerProvisioningState_InProvisioningState_Poll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// changing where we're setting this for the heck of it
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInProvisioningState(statusInProgress),
		responseWithStatusInProvisioningState(statusInProgress),
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since we should be using
		// `apiVersion` (which otherwise gets parsed from the URI in `provisioningStatePollerFromResponse`)
		Client:     client.NewClient(server.URL, "Example", "2015-01-01"),
		ApiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:           helper.expectedApiVersion,
		client:               resourceManagerClient,
		initialRetryDuration: 10,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}

	expectedStates := []pollers.PollingStatus{
		pollers.PollingStatusInProgress,
		pollers.PollingStatusInProgress,
		pollers.PollingStatusSucceeded,
	}
	for i, expected := range expectedStates {
		t.Logf("Poll %d..", i)
		actual, err := poller.Poll(ctx)
		if err != nil {
			t.Fatal(err.Error())
		}
		if actual.Status != expected {
			t.Fatalf("expected %q but got %q", string(expected), string(actual.Status))
		}
	}
}

func TestPollerProvisioningState_InStatus_Poll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// changing where we're setting this for the heck of it
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInStatusField(statusInProgress),
		responseWithStatusInStatusField(statusInProgress),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		ApiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:           helper.expectedApiVersion,
		client:               resourceManagerClient,
		initialRetryDuration: 10,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}

	expectedStates := []pollers.PollingStatus{
		pollers.PollingStatusInProgress,
		pollers.PollingStatusInProgress,
		pollers.PollingStatusSucceeded,
	}
	for i, expected := range expectedStates {
		t.Logf("Poll %d..", i)
		actual, err := poller.Poll(ctx)
		if err != nil {
			t.Fatal(err.Error())
		}
		if actual.Status != expected {
			t.Fatalf("expected %q but got %q", string(expected), string(actual.Status))
		}
	}
}

func TestPollerProvisioningState_FailedOrCancelled_IncludesTheResponse(t *testing.T) {
	testCases := []struct {
		state    string
		expected string
	}{
		{
			state:    "Failed",
			expected: "polling failed: the Azure API reported the state \"Failed\"",
		},
		{
			state:    "Canceled",
			expected: "polling was cancelled: the Azure API reported the state \"Canceled\"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.state, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
			defer cancel()

			body := fmt.Sprintf(`{"id":"/subscriptions/000/resourceGroups/rg/providers/Microsoft.Foo/things/t1","properties":{"provisioningState":%q,"provisioningError":{"code":"SubnetIsFull","message":"Subnet snet-app does not have enough free addresses for 3 instances."}}}`, tc.state)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(body))
			}))
			defer server.Close()

			dataplaneClient := &Client{
				Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
				ApiVersion: "2020-01-01",
			}
			poller := provisioningStatePoller{
				apiVersion:           "2020-02-01",
				client:               dataplaneClient,
				initialRetryDuration: 10,
				originalUri:          "/provisioning-state/poll",
				resourcePath:         "/provisioning-state/poll",
			}

			_, err := poller.Poll(ctx)
			if err == nil {
				t.Fatal("expected an error but didn't get one")
			}

			expected := fmt.Sprintf("%s\n\nAPI Response:\n\n----[start]----\n%s\n-----[end]-----", tc.expected, body)
			if err.Error() != expected {
				t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
			}

			// the body should be left for the caller to read
			var resp *client.Response
			var failed pollers.PollingFailedError
			var cancelled pollers.PollingCancelledError
			switch {
			case errors.As(err, &failed):
				resp = failed.HttpResponse

			case errors.As(err, &cancelled):
				resp = cancelled.HttpResponse

			default:
				t.Fatalf("expected a PollingFailedError or PollingCancelledError but got %T", err)
			}

			remaining, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("reading response body: %v", err)
			}
			if string(remaining) != body {
				t.Fatalf("expected the response body to still be %q, got %q", body, string(remaining))
			}
		})
	}
}

func TestPollerProvisioningState_DroppedConnection_ReportsWhyOnceOutOfAttempts(t *testing.T) {
	// a short deadline means a request is attempted once, rather than being retried by Execute()
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		dropConnection(t, w)
	}))
	defer server.Close()

	dataplaneClient := &Client{
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		ApiVersion: "2020-01-01",
	}
	pollerType := &provisioningStatePoller{
		apiVersion:            "2020-02-01",
		client:                dataplaneClient,
		initialRetryDuration:  10,
		originalUri:           "/provisioning-state/poll",
		resourcePath:          "/provisioning-state/poll",
		maxDroppedConnections: 3,
	}
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	err := poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	// the reason the request failed differs between platforms, so only check it is the request that is reported
	expected := fmt.Sprintf("executing request: Get \"%s/provisioning-state/poll?api-version=2020-02-01", server.URL)
	if !strings.HasPrefix(err.Error(), expected) {
		t.Fatalf("expected an error starting %q, got %q", expected, err.Error())
	}
	if actual := calls.Load(); actual != 3 {
		t.Fatalf("expected to be called 3 times but got %d", actual)
	}
}
