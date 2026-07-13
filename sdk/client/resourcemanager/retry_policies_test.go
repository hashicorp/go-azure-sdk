// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcemanager

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestRetryOn409ConflictFunc(t *testing.T) {
	cases := []struct {
		name          string
		statusCode    int
		headers       http.Header
		body          string
		expectedRetry bool
	}{
		{
			name:          "Not a 409",
			statusCode:    http.StatusOK,
			expectedRetry: false,
		},
		{
			name:          "409 with Retry-After header",
			statusCode:    http.StatusConflict,
			headers:       http.Header{"Retry-After": []string{"120"}},
			expectedRetry: true,
		},
		{
			name:          "409 with non-terminal status",
			statusCode:    http.StatusConflict,
			body:          `{"status": "Updating"}`,
			expectedRetry: true,
		},
		{
			name:          "409 with non-terminal provisioningState",
			statusCode:    http.StatusConflict,
			body:          `{"properties": {"provisioningState": "Creating"}}`,
			expectedRetry: true,
		},
		{
			name:          "409 with terminal status Failed",
			statusCode:    http.StatusConflict,
			body:          `{"status": "Failed"}`,
			expectedRetry: false,
		},
		{
			name:          "409 with terminal status Succeeded",
			statusCode:    http.StatusConflict,
			body:          `{"properties": {"provisioningState": "Succeeded"}}`,
			expectedRetry: false,
		},
		{
			name:          "409 with empty body",
			statusCode:    http.StatusConflict,
			body:          ``,
			expectedRetry: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tc.statusCode,
				Header:     tc.headers,
			}
			if tc.body != "" {
				resp.Body = io.NopCloser(bytes.NewBufferString(tc.body))
			}

			retry, err := RetryOn409ConflictFunc(resp, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if retry != tc.expectedRetry {
				t.Errorf("expected retry=%v, got %v", tc.expectedRetry, retry)
			}
		})
	}
}
