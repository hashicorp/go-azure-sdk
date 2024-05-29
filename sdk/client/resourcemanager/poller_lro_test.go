// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

func TestPollerLRO_InProvisioningState_TerminalStates(t *testing.T) {
	// This test ensures that each of the Custom Statuses works in the provisioningState field
	ctx := context.TODO()
	for state, expected := range longRunningOperationCustomStatuses {
		t.Run(fmt.Sprintf("%s/%s", string(state), string(expected)), func(t *testing.T) {
			helpers := newLongRunningOperationsEndpoint(true, []longRunningOperationsExpectedResponse{
				responseWithStatus(state),
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
				if expected == pollers.PollingStatusCancelled || expected == pollers.PollingStatusFailed {
					return
				}

				t.Fatal(err.Error())
			}
			if result.Status != expected {
				t.Fatalf("expected status to be %q but got %q", expected, result.Status)
			}
			// sanity-checking
			helpers.assertCalled(t, 1)
		})
	}
}

func TestPollerLRO_InStatus_TerminalStates(t *testing.T) {
	// This test ensures that each of the Custom Statuses works in the status field
	ctx := context.TODO()
	for state, expected := range longRunningOperationCustomStatuses {
		t.Run(fmt.Sprintf("%s/%s", string(state), string(expected)), func(t *testing.T) {
			helpers := newLongRunningOperationsEndpoint(false, []longRunningOperationsExpectedResponse{
				responseWithStatus(state),
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
				if expected == pollers.PollingStatusCancelled || expected == pollers.PollingStatusFailed {
					return
				}

				t.Fatal(err.Error())
			}
			if result.Status != expected {
				t.Fatalf("expected status to be %q but got %q", expected, result.Status)
			}
			// sanity-checking
			helpers.assertCalled(t, 1)
		})
	}
}

func TestPollerLRO_InProvisioningState_ImmediateSuccess(t *testing.T) {
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint(true, []longRunningOperationsExpectedResponse{
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(false, []longRunningOperationsExpectedResponse{
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(true, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(false, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(true, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatus(statusInProgress),
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(false, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusAccepted),
		responseWithStatus(statusInProgress),
		responseWithStatus(statusSucceeded),
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

func TestPollerLRO_InProvisioningState_404ThenImmediateSuccess(t *testing.T) {
	// This scenario handles the API returning a 404 initially, then succeeded
	// This happens when there's an API bug, since we shouldn't get the initial 404 - and can be
	// seen in this issue: https://github.com/hashicorp/go-azure-sdk/issues/886 (although this test
	// assumes we've immediately passed)
	ctx := context.TODO()
	helpers := newLongRunningOperationsEndpoint(true, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(false, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(true, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatus(statusInProgress),
		responseWithHttpStatusCode(http.StatusNotFound), // let's fake some eventual consistency
		responseWithStatus(statusSucceeded),
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
	helpers := newLongRunningOperationsEndpoint(false, []longRunningOperationsExpectedResponse{
		responseWithHttpStatusCode(http.StatusNotFound),
		responseWithStatus(statusInProgress),
		responseWithHttpStatusCode(http.StatusNotFound), // let's fake some eventual consistency
		responseWithStatus(statusSucceeded),
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
