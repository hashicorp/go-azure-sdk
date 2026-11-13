// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
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

func TestPollerLRO_AcceptedThenNoContent(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint([]expectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithHttpStatusCode(http.StatusNoContent),
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
	helpers.assertCalled(t, 2)
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

func TestPollerLRO_FailedOrCancelled_IncludesTheErrorAndTheResponse(t *testing.T) {
	testCases := []struct {
		status   string
		expected string
	}{
		{
			status:   "Failed",
			expected: "polling failed: the Azure API returned the following error:",
		},
		{
			status:   "Canceled",
			expected: "polling was cancelled: the Azure API returned the following error:",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.status, func(t *testing.T) {
			ctx := context.TODO()
			body := fmt.Sprintf(`{"id":"/subscriptions/000/providers/Microsoft.Foo/locations/westeurope/operations/1","status":%q,"error":{"code":"SubnetIsFull","message":"Subnet snet-app does not have enough free addresses for 3 instances."}}`, tc.status)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(body))
			}))
			defer server.Close()

			helpers := newLongRunningOperationsEndpoint([]expectedResponse{})
			response := &client.Response{
				Response: helpers.response(),
			}
			client := client.NewClient(server.URL, "MyService", "2020-02-01")
			poller, err := longRunningOperationPollerFromResponse(response, client)
			if err != nil {
				t.Fatal(err.Error())
			}

			_, err = poller.Poll(ctx)
			if err == nil {
				t.Fatal("expected an error but didn't get one")
			}

			expected := fmt.Sprintf("%s\n\nStatus: %q\nCode: \"SubnetIsFull\"\nMessage: \"Subnet snet-app does not have enough free addresses for 3 instances.\"\nActivity Id: \"\"\n\n---\n\nAPI Response:\n\n----[start]----\n%s\n-----[end]-----\n", tc.expected, tc.status, body)
			if err.Error() != expected {
				t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
			}
		})
	}
}

func TestPollerLRO_ResponseThatCannotBeRead_IncludesTheResponse(t *testing.T) {
	testCases := []struct {
		name        string
		contentType string
		expected    string
	}{
		{
			name:        "not JSON",
			contentType: "text/html",
			expected:    "internal-error: polling support for the Content-Type \"text/html\" was not implemented, the API returned 200 with: <html><head><title>Sign in</title></head></html>",
		},
		{
			name:        "said to be JSON but is not",
			contentType: "application/json",
			expected:    "unmarshalling response body: invalid character '<' looking for beginning of value, the API returned 200 with: <html><head><title>Sign in</title></head></html>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				_, _ = w.Write([]byte("<html><head><title>Sign in</title></head></html>"))
			}))
			defer server.Close()

			helpers := newLongRunningOperationsEndpoint([]expectedResponse{})
			response := &client.Response{
				Response: helpers.response(),
			}
			client := client.NewClient(server.URL, "MyService", "2020-02-01")
			poller, err := longRunningOperationPollerFromResponse(response, client)
			if err != nil {
				t.Fatal(err.Error())
			}

			_, err = poller.Poll(ctx)
			if err == nil {
				t.Fatal("expected an error but didn't get one")
			}
			if err.Error() != tc.expected {
				t.Fatalf("expected error %q, got %q", tc.expected, err.Error())
			}
		})
	}
}

func TestPollerLRO_DroppedConnection_ReportsWhyOnceOutOfAttempts(t *testing.T) {
	// a short deadline means a request is attempted once, rather than being retried by Execute()
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		dropConnection(t, w)
	}))
	defer server.Close()

	helpers := newLongRunningOperationsEndpoint([]expectedResponse{})
	response := &client.Response{
		Response: helpers.response(),
	}
	client := client.NewClient(server.URL, "MyService", "2020-02-01")
	pollerType, err := longRunningOperationPollerFromResponse(response, client)
	if err != nil {
		t.Fatal(err.Error())
	}
	pollerType.initialRetryDuration = 10 * time.Millisecond
	poller := pollers.NewPoller(pollerType, 10*time.Millisecond, pollers.DefaultNumberOfDroppedConnectionsToAllow)

	err = poller.PollUntilDone(ctx)
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	// the reason the request failed differs between platforms, so only check it is the request that is reported
	expected := fmt.Sprintf("Get %q: ", server.URL+"/lro/poll")
	if !strings.HasPrefix(err.Error(), expected) {
		t.Fatalf("expected an error starting %q, got %q", expected, err.Error())
	}
	if actual := calls.Load(); actual != 3 {
		t.Fatalf("expected to be called 3 times but got %d", actual)
	}
}
