// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

func TestErrorParser_FromResponse(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected string
	}{
		{
			name:     "code and message",
			body:     `<?xml version="1.0" encoding="utf-8"?><Error><Code>ContainerNotFound</Code><Message>The specified container does not exist.</Message></Error>`,
			expected: "ContainerNotFound: The specified container does not exist.",
		},
		{
			name:     "byte order mark before the document",
			body:     "\xef\xbb\xbf" + `<?xml version="1.0" encoding="utf-8"?><Error><Code>ContainerNotFound</Code><Message>The specified container does not exist.</Message></Error>`,
			expected: "ContainerNotFound: The specified container does not exist.",
		},
		{
			name:     "the reason is in an element specific to the error",
			body:     `<?xml version="1.0" encoding="utf-8"?><Error><Code>AuthenticationFailed</Code><Message>Server failed to authenticate the request.</Message><AuthenticationErrorDetail>Signed expiry time [Fri, 13 Nov 2026 09:00:00 GMT] must be after signed start time [Fri, 13 Nov 2026 10:00:00 GMT]</AuthenticationErrorDetail></Error>`,
			expected: "AuthenticationFailed: Server failed to authenticate the request.\n  AuthenticationErrorDetail: Signed expiry time [Fri, 13 Nov 2026 09:00:00 GMT] must be after signed start time [Fri, 13 Nov 2026 10:00:00 GMT]",
		},
		{
			name:     "several elements specific to the error, one of them nested",
			body:     `<?xml version="1.0" encoding="utf-8"?><Error><Code>InvalidQueryParameterValue</Code><Message>Value for one of the query parameters specified in the request URI is invalid.</Message><QueryParameterName>comp</QueryParameterName><QueryParameterValue>lst</QueryParameterValue><ExceptionDetails><ExceptionMessage>The value is not supported</ExceptionMessage></ExceptionDetails></Error>`,
			expected: "InvalidQueryParameterValue: Value for one of the query parameters specified in the request URI is invalid.\n  QueryParameterName: comp\n  QueryParameterValue: lst\n  ExceptionDetails: <ExceptionMessage>The value is not supported</ExceptionMessage>",
		},
		{
			name: "no body",
			body: "",
		},
		{
			name: "an error in JSON",
			body: `{"error":{"code":"FilesystemNotFound","message":"The specified filesystem does not exist."}}`,
		},
		{
			name: "something else entirely",
			body: `<html><head><title>502 Bad Gateway</title></head></html>`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &http.Response{
				Body: io.NopCloser(bytes.NewReader([]byte(tc.body))),
			}

			actual := ""
			if err := (&ErrorParser{}).FromResponse(resp); err != nil {
				actual = err.Error()
			}
			if actual != tc.expected {
				t.Fatalf("expected %q, got %q", tc.expected, actual)
			}

			// the body should be left for the caller to read
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("reading response body: %v", err)
			}
			if string(body) != tc.body {
				t.Fatalf("expected the response body to still be %q, got %q", tc.body, string(body))
			}
		})
	}
}

func TestClient_UnexpectedStatus(t *testing.T) {
	testCases := []struct {
		name     string
		method   string
		status   int
		headers  map[string]string
		body     string
		expected string
	}{
		{
			name:   "an XML error with a code and a message",
			method: http.MethodGet,
			status: http.StatusNotFound,
			headers: map[string]string{
				"Content-Type":    "application/xml",
				"x-ms-error-code": "ContainerNotFound",
				"x-ms-request-id": "11111111-aaaa-bbbb-cccc-000000000001",
			},
			body:     `<?xml version="1.0" encoding="utf-8"?><Error><Code>ContainerNotFound</Code><Message>The specified container does not exist.</Message></Error>`,
			expected: "unexpected status 404 (404 Not Found) with ContainerNotFound: The specified container does not exist.\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001",
		},
		{
			name:   "an XML error, with the request ID it already mentions left alone",
			method: http.MethodGet,
			status: http.StatusForbidden,
			headers: map[string]string{
				"Content-Type":    "application/xml",
				"x-ms-error-code": "AuthenticationFailed",
				"x-ms-request-id": "11111111-aaaa-bbbb-cccc-000000000001",
			},
			body:     "<?xml version=\"1.0\" encoding=\"utf-8\"?><Error><Code>AuthenticationFailed</Code><Message>Server failed to authenticate the request.\nRequestId:11111111-aaaa-bbbb-cccc-000000000001\nTime:2026-11-13T10:00:00.0000000Z</Message><AuthenticationErrorDetail>Signature did not match.</AuthenticationErrorDetail></Error>",
			expected: "unexpected status 403 (403 Forbidden) with AuthenticationFailed: Server failed to authenticate the request.\nRequestId:11111111-aaaa-bbbb-cccc-000000000001\nTime:2026-11-13T10:00:00.0000000Z\n  AuthenticationErrorDetail: Signature did not match.",
		},
		{
			name:   "a HEAD request, where the error is only in the headers",
			method: http.MethodHead,
			status: http.StatusNotFound,
			headers: map[string]string{
				"x-ms-error-code": "ContainerNotFound",
				"x-ms-request-id": "11111111-aaaa-bbbb-cccc-000000000001",
			},
			expected: "unexpected status 404 (404 Not Found) received with no body\nx-ms-error-code: ContainerNotFound\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001",
		},
		{
			name:   "an error in JSON, as Data Lake sends",
			method: http.MethodGet,
			status: http.StatusNotFound,
			headers: map[string]string{
				"Content-Type":    "application/json",
				"x-ms-error-code": "FilesystemNotFound",
				"x-ms-request-id": "11111111-aaaa-bbbb-cccc-000000000001",
			},
			body:     `{"error":{"code":"FilesystemNotFound","message":"The specified filesystem does not exist."}}`,
			expected: "unexpected status 404 (404 Not Found) with error: FilesystemNotFound: The specified filesystem does not exist.\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001",
		},
		{
			name:   "an error in JSON, as Tables sends",
			method: http.MethodGet,
			status: http.StatusNotFound,
			headers: map[string]string{
				"Content-Type":    "application/json;odata=minimalmetadata;streaming=true;charset=utf-8",
				"x-ms-request-id": "11111111-aaaa-bbbb-cccc-000000000001",
			},
			body:     `{"odata.error":{"code":"TableNotFound","message":{"lang":"en-US","value":"The table specified does not exist."}}}`,
			expected: "unexpected status 404 (404 Not Found) with error: TableNotFound: The table specified does not exist.\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, v := range tc.headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer server.Close()

			storageClient, err := NewStorageClient(server.URL, "blob", "2023-11-03")
			if err != nil {
				t.Fatalf("building client: %v", err)
			}

			req, err := storageClient.NewRequest(ctx, client.RequestOptions{
				ContentType:         "application/xml; charset=utf-8",
				ExpectedStatusCodes: []int{http.StatusOK},
				HttpMethod:          tc.method,
				Path:                "/example",
			})
			if err != nil {
				t.Fatalf("building request: %v", err)
			}

			resp, err := req.Execute(ctx)
			if err == nil {
				t.Fatal("expected an error for an unexpected status")
			}
			if err.Error() != tc.expected {
				t.Fatalf("expected error:\n%s\n\ngot:\n%s", tc.expected, err.Error())
			}

			// whatever the error made of it, the body is left for the caller to read
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("reading response body: %v", err)
			}
			if string(body) != tc.body {
				t.Fatalf("expected the response body to still be %q, got %q", tc.body, string(body))
			}
		})
	}
}
