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

func TestPollerLRO_InProvisioningState_ImmediateSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("Polling..")
	result, err := poller.Poll(ctx)
	if err != nil {
		if result.Status == pollers.PollingStatusCancelled || result.Status == pollers.PollingStatusFailed {
			return
		}

		t.Fatal(err.Error())
	}
	if result.Status != pollers.PollingStatusSucceeded {
		t.Fatalf("expected status to be %q but got %q", pollers.PollingStatusSucceeded, result.Status)
	}
	// sanity-checking
	helpers.assertCalled(t, 1)
}

func TestPollerLRO_InStatus_ImmediateSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("Polling..")
	result, err := poller.Poll(ctx)
	if err != nil {
		if result.Status == pollers.PollingStatusCancelled || result.Status == pollers.PollingStatusFailed {
			return
		}

		t.Fatal(err.Error())
	}
	if result.Status != pollers.PollingStatusSucceeded {
		t.Fatalf("expected status to be %q but got %q", pollers.PollingStatusSucceeded, result.Status)
	}
	// sanity-checking
	helpers.assertCalled(t, 1)
}

func TestPollerLRO_InProvisioningState_AcceptedThenImmediateSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 202 Accepted
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
	// sanity-checking
	helpers.assertCalled(t, 2)
}

func TestPollerLRO_InStatus_AcceptedThenImmediateSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 202 Accepted
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
	// sanity-checking
	helpers.assertCalled(t, 2)
}

func TestPollerLRO_InProvisioningState_AcceptedThenInProgressThenSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatusInProvisioningState(statusInProgress),
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 202 Accepted
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
	// sanity-checking
	helpers.assertCalled(t, 3)
}

func TestPollerLRO_InStatus_AcceptedThenInProgressThenSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatusInStatusField(statusInProgress),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 202 Accepted
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
	// sanity-checking
	helpers.assertCalled(t, 3)
}

func TestPollerLRO_NoTerminalStatusThenSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithNoTerminalState(),
		responseWithNoTerminalState(),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress,
		pollers.PollingStatusInProgress,
		pollers.PollingStatusInProgress,
		pollers.PollingStatusSucceeded,
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
	// sanity-checking
	helpers.assertCalled(t, 4)
}

func TestPollerLRO_ArbitraryStatusThenSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatusInStatusField("Reticulating Splines"),
		responseWithStatusInStatusField("Retracting Phong Shader"),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress,
		pollers.PollingStatusInProgress,
		pollers.PollingStatusInProgress,
		pollers.PollingStatusSucceeded,
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
	// sanity-checking
	helpers.assertCalled(t, 4)
}

func TestPollerLRO_InProvisioningState_AcceptedThenDroppedThenInProgressThenSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseThatDropsTheConnection(),
		responseWithStatusInProvisioningState(statusInProgress),
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 202 Accepted
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
	// sanity-checking - expect 4 calls but 3 statuses (since the dropped connection is silently retried)
	helpers.assertCalled(t, 4)
}

func TestPollerLRO_InStatus_AcceptedThenDroppedThenInProgressThenSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseThatDropsTheConnection(),
		responseWithStatusInStatusField(statusInProgress),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 202 Accepted
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
	helpers.assertCalled(t, 4)
}

func TestPollerLRO_InProvisioningState_404ThenImmediateSuccess(t *testing.T) {
	// This scenario handles the API returning a 404 initially, then succeeded
	// This happens when there's an API bug, since we shouldn't get the initial 404 - and can be
	// seen in this issue: https://github.com/hashicorp/go-azure-sdk/issues/886 (although this test
	// assumes we've immediately passed)
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 404
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
	// sanity-checking
	helpers.assertCalled(t, 2)
}

func TestPollerLRO_InStatus_404ThenImmediateSuccess(t *testing.T) {
	// This scenario handles the API returning a 404 initially, then succeeded
	// This happens when there's an API bug, since we shouldn't get the initial 404 - and can be
	// seen in this issue: https://github.com/hashicorp/go-azure-sdk/issues/886 (although this test
	// assumes we've immediately passed)
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 404
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
	// sanity-checking
	helpers.assertCalled(t, 2)
}

func TestPollerLRO_InProvisioningState_404ThenInProgressThenSucceeded(t *testing.T) {
	// This scenario handles the API returning a 404 initially, then in progress, then succeeded
	// This happens when there's an API bug, since we shouldn't get the initial 404 - and can be
	// seen in this issue: https://github.com/hashicorp/go-azure-sdk/issues/886
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatusInProvisioningState(statusInProgress),
		responseWithHttpStatusCode(http.StatusNotFound), // let's fake some eventual consistency
		responseWithStatusInProvisioningState(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 404
		pollers.PollingStatusInProgress, // the in progress
		pollers.PollingStatusInProgress, // another 404 for good measure, call it eventual consistency
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
	// sanity-checking
	helpers.assertCalled(t, 4)
}

func TestPollerLRO_InStatus_404ThenInProgressThenSucceeded(t *testing.T) {
	// This scenario handles the API returning a 404 initially, then in progress, then succeeded
	// This happens when there's an API bug, since we shouldn't get the initial 404 - and can be
	// seen in this issue: https://github.com/hashicorp/go-azure-sdk/issues/886
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatusInStatusField(statusInProgress),
		responseWithHttpStatusCode(http.StatusNotFound), // let's fake some eventual consistency
		responseWithStatusInStatusField(statusSucceeded),
	})
	server := httptest.NewServer(http.HandlerFunc(helpers.endpoint(t)))
	defer server.Close()

	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	poller, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedStatuses := []pollers.PollingStatus{
		pollers.PollingStatusInProgress, // the 404
		pollers.PollingStatusInProgress, // the in progress
		pollers.PollingStatusInProgress, // another 404 for good measure, call it eventual consistency
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
	// sanity-checking
	helpers.assertCalled(t, 4)
}

func TestPollerLRO_VCRErrorHandling(t *testing.T) {
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

			helpers := newLongRunningOperationsEndpoint([]expectedResponse{})
			response := &client.Response{
				Response: helpers.response(),
			}

			baseClient := client.NewClient("https://management.azure.com", "Example", "2020-01-01")
			baseClient.SetTransport(errRoundTripper{errMsg: testCase.errMsg}, client.TransportModeVCRReplay)

			poller, err := longRunningOperationPollerFromResponse(response, baseClient)
			if err != nil {
				t.Fatal(err.Error())
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
