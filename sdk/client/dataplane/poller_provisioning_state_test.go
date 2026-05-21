// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplane

import (
	"context"
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
		pollers.PollingStatusUnknown,
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

func TestPollerProvisioningState_VCRErrorHandling(t *testing.T) {
	testCases := []struct {
		name                     string
		errMsg                   string
		expectReplayMiss         bool
		expectResultStatus       *pollers.PollingStatus
		expectError              bool
		expectDroppedConnections int
	}{
		{
			name:                     "replay miss url error",
			errMsg:                   client.VCRInteractionNotFoundErrMsg,
			expectReplayMiss:         true,
			expectError:              true,
			expectDroppedConnections: 0,
		},
		{
			name:                     "non retryable url error",
			errMsg:                   "unsupported protocol scheme",
			expectResultStatus:       func() *pollers.PollingStatus { s := pollers.PollingStatusUnknown; return &s }(),
			expectDroppedConnections: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			dataplaneClient := &Client{
				Client:     client.NewClient("https://management.azure.com", "Example", "2020-01-01"),
				ApiVersion: "2020-01-01",
			}
			dataplaneClient.SetTransport(errRoundTripper{errMsg: testCase.errMsg}, client.TransportModeVCRReplay)

			poller := provisioningStatePoller{
				apiVersion:            "2020-01-01",
				client:                dataplaneClient,
				initialRetryDuration:  10 * time.Millisecond,
				originalUri:           "/provisioning-state/poll",
				resourcePath:          "/provisioning-state/poll",
				maxDroppedConnections: 3,
			}

			result, err := poller.Poll(ctx)
			if testCase.expectError {
				if err == nil {
					t.Fatal("expected polling to return an error, but got no error")
				}
				if result != nil {
					t.Fatalf("expected no poll result, got: %+v ,error %T (err=%v)", result, err, err)
				}
				if testCase.expectReplayMiss != client.IsVCRReplayMissError(err) {
					t.Fatalf("expected IsVCRReplayMissError=%t, got %t (err=%v)", testCase.expectReplayMiss, client.IsVCRReplayMissError(err), err)
				}
				if !strings.Contains(err.Error(), testCase.errMsg) {
					t.Fatalf("expected error to contain %q, got %v", testCase.errMsg, err)
				}
			}
			if testCase.expectResultStatus != nil && (result == nil || result.Status != *testCase.expectResultStatus) {
				t.Fatalf("expected poll result with status %q, got: %+v", *testCase.expectResultStatus, result)
			}
			if poller.droppedConnectionCount != testCase.expectDroppedConnections {
				t.Fatalf("expected droppedConnectionCount to be %d, got %d", testCase.expectDroppedConnections, poller.droppedConnectionCount)
			}
		})
	}
}
