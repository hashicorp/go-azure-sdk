// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplane

import (
	"encoding/json"
	"net/http"
	"testing"
)

// longRunningOperationsEndpoint is a test helper struct which makes testing the expected behaviour of the
// long-running operations endpoint easier. Originally these started out as a function per test, but with
// so much boilerplate, this helped make the tests more readable, so whilst this has some boilerplate I
// suspect it's worth it for the tradeoff of having tests that are more understandable.
type longRunningOperationsEndpoint struct {
	numberOfTimesCalled int
	responses           []expectedResponse
}

func newLongRunningOperationsEndpoint(responses []expectedResponse) longRunningOperationsEndpoint {
	return longRunningOperationsEndpoint{
		numberOfTimesCalled: 0,
		responses:           responses,
	}
}

func (e *longRunningOperationsEndpoint) endpoint(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Server request %q %q", r.Method, r.RequestURI)

		if r.Method != http.MethodGet && r.RequestURI != "/lro/poll" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}

		e.numberOfTimesCalled++
		if e.numberOfTimesCalled > len(e.responses) {
			t.Fatal("called too many times")
		}
		response := e.responses[e.numberOfTimesCalled-1]
		if response.dropsConnection {
			dropConnection(t, w)
			return
		}
		if response.httpStatusCode != nil {
			w.WriteHeader(*response.httpStatusCode)
			return
		}

		if response.status == nil {
			t.Fatalf("missing either a httpStatusCode or a status to return")
		}
		payload := operationResult{
			Properties: operationResultProperties{
				ProvisioningState: "",
			},
			Status: "",
		}
		if response.useProvisioningState {
			payload.Properties.ProvisioningState = *response.status
		}
		if response.useStatus {
			payload.Status = *response.status
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Fatal(err.Error())
			return
		}
	}
}

func (e longRunningOperationsEndpoint) assertCalled(t *testing.T, expected int) {
	if e.numberOfTimesCalled != expected {
		t.Fatalf("expected to be called %d times but got %d", expected, e.numberOfTimesCalled)
	}
}

func (e longRunningOperationsEndpoint) response() *http.Response {
	return &http.Response{
		Header: map[string][]string{
			http.CanonicalHeaderKey("Azure-AsyncOperation"): {"http://localhost/lro/poll"},
			http.CanonicalHeaderKey("Retry-After"):          {"1"},
		},
	}
}
