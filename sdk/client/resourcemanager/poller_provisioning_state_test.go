// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"context"
	"net/http"
	"net/http/httptest"
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

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
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
		t.Fatalf(err.Error())
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
		apiVersion: "2020-01-01",
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
		t.Fatalf(err.Error())
	}
	if actual.Status != pollers.PollingStatusSucceeded {
		t.Fatalf("expected %q but got %q", string(pollers.PollingStatusSucceeded), string(actual.Status))
	}
}

func TestPollerProvisioningState_OkWithNoBody_Immediate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// In this instance we're not setting any `status` value
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusOK),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
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
		t.Fatalf(err.Error())
	}
	if actual.Status != pollers.PollingStatusSucceeded {
		t.Fatalf("expected %q but got %q", string(pollers.PollingStatusSucceeded), string(actual.Status))
	}
}

func TestPollerProvisioningState_OkWithNoBody_AfterPolling(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	// changing where we're setting this for the heck of it
	helper := newProvisioningStateEndpoint([]expectedResponse{
		responseWithStatusInProvisioningState(statusInProgress), // intentionally in provisioningState
		responseWithStatusInStatusField(statusInProgress),       // intentionally in status
		responseWithHttpStatusCode(http.StatusOK),
	})
	server := httptest.NewServer(http.HandlerFunc(helper.endpoint(t)))
	defer server.Close()

	resourceManagerClient := &Client{
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
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
			t.Fatalf(err.Error())
		}
		if actual.Status != expected {
			t.Fatalf("expected %q but got %q", string(expected), string(actual.Status))
		}
	}
}

func TestPollerProvisioningState_InProvisioningState_DroppedThenInProgressThenSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
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
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:           helper.expectedApiVersion,
		client:               resourceManagerClient,
		initialRetryDuration: 10,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // working on it
		// NOTE: the Dropped Connection will be ignored/silently retried
		pollers.PollingStatusInProgress, // working on it
		pollers.PollingStatusSucceeded,  // good
	}
	for i, expected := range expectedStatuses {
		t.Logf("Poll %d..", i)
		result, err := poller.Poll(ctx)
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
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
	}
	poller := provisioningStatePoller{
		apiVersion:           helper.expectedApiVersion,
		client:               resourceManagerClient,
		initialRetryDuration: 10,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // working on it
		// NOTE: the Dropped Connection will be ignored/silently retried
		pollers.PollingStatusInProgress, // working on it
		pollers.PollingStatusSucceeded,  // good
	}
	for i, expected := range expectedStatuses {
		t.Logf("Poll %d..", i)
		result, err := poller.Poll(ctx)
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
		// NOTE: the use of a different API Version here is _intentional_ to ensure it's unused since the
		Client:     client.NewClient(server.URL, "Example", "2020-01-01"),
		apiVersion: "2020-01-01",
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
			t.Fatalf(err.Error())
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
		apiVersion: "2020-01-01",
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
			t.Fatalf(err.Error())
		}
		if actual.Status != expected {
			t.Fatalf("expected %q but got %q", string(expected), string(actual.Status))
		}
	}
}
