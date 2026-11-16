// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplane

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
		// NOTE: the Dropped Connection will be transparently retried by Execute()
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
		pollers.PollingStatusInProgress, // working on it (dropped connection retried by Execute)
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

func TestPollerProvisioningState_FailedIncludesErrorDetails(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	body := `{"properties":{"provisioningState":"Failed"},"error":{"code":"DeploymentFailed","message":"At least one resource deployment operation failed.","details":[{"code":"Conflict","message":"The storage account name is already taken."}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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
		initialRetryDuration: 1,
		originalUri:          "/provisioning-state/poll",
		resourcePath:         "/provisioning-state/poll",
	}

	_, err := poller.Poll(ctx)
	if err == nil {
		t.Fatal("expected error for Failed provisioningState, got nil")
	}
	var failedErr pollers.PollingFailedError
	if !errors.As(err, &failedErr) {
		t.Fatalf("expected PollingFailedError, got %T: %v", err, err)
	}
	if !strings.Contains(failedErr.Error(), "The storage account name is already taken.") {
		t.Fatalf("expected PollingFailedError to include failure details, got %q", failedErr.Error())
	}
}
