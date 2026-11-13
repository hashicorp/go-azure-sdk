// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

func TestExecute_UnexpectedStatus(t *testing.T) {
	testCases := []struct {
		name     string
		status   int
		headers  map[string]string
		body     string
		expected string
	}{
		// the shapes of error that have always been reported in full
		{
			name:     "a code and a message",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"InvalidParameter","message":"The value of parameter sku is invalid."}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: InvalidParameter: The value of parameter sku is invalid.",
		},
		{
			name:     "a message wrapped in an object",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"odata.error":{"code":"Request_BadRequest","message":{"lang":"en","value":"Invalid value specified for property 'displayName'."}}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: Request_BadRequest: Invalid value specified for property 'displayName'.",
		},
		{
			name:     "a message inside the inner error",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"BadRequest","message":"The server could not process the request.","innerError":{"message":"1034: Policy contains invalid applications"}}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: BadRequest: The server could not process the request.: 1034: Policy contains invalid applications",
		},
		{
			name:     "a code and a message that are not inside an error",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"Code":"BadRequest","Message":"The parameter LinuxFxVersion has an invalid value.","Details":[{"Message":"The parameter LinuxFxVersion has an invalid value."},{"Code":"BadRequest"}]}`,
			expected: "unexpected status 400 (400 Bad Request) with response: {\"Code\":\"BadRequest\",\"Message\":\"The parameter LinuxFxVersion has an invalid value.\",\"Details\":[{\"Message\":\"The parameter LinuxFxVersion has an invalid value.\"},{\"Code\":\"BadRequest\"}]}",
		},
		{
			name:     "an error that cannot be parsed, as its details are text",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"BadRequest","message":"Validation failed.","details":"The property 'sku.tier' is required when 'sku.name' is Premium."}}`,
			expected: "unexpected status 400 (400 Bad Request) with response: {\"error\":{\"code\":\"BadRequest\",\"message\":\"Validation failed.\",\"details\":\"The property 'sku.tier' is required when 'sku.name' is Premium.\"}}",
		},
		{
			name:     "JSON that is not an object",
			status:   http.StatusNotFound,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `"The requested resource was not found."`,
			expected: "unexpected status 404 (404 Not Found) with response: \"The requested resource was not found.\"",
		},
		{
			name:     "a page from something in front of the API",
			status:   http.StatusForbidden,
			headers:  map[string]string{"Content-Type": "text/html"},
			body:     "<html><head><title>403 Forbidden</title></head><body><h1>Forbidden</h1></body></html>",
			expected: "unexpected status 403 (403 Forbidden) with response: <html><head><title>403 Forbidden</title></head><body><h1>Forbidden</h1></body></html>",
		},
		{
			name:     "plain text",
			status:   http.StatusForbidden,
			headers:  map[string]string{"Content-Type": "text/plain"},
			body:     "IP address 203.0.113.7 is not allowed",
			expected: "unexpected status 403 (403 Forbidden) with response: IP address 203.0.113.7 is not allowed",
		},
		{
			name:     "no body",
			status:   http.StatusNotFound,
			expected: "unexpected status 404 (404 Not Found) received with no body",
		},

		// the parts of an error that explain it, but are not its code or message
		{
			name:     "the reason is in the details",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"InvalidRequestFormat","message":"Cannot parse the request.","details":[{"code":"InvalidJsonReferenceFormat","message":"Reference Id /subscriptions/000 is not formatted correctly."}]}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: InvalidRequestFormat: Cannot parse the request.\n  InvalidJsonReferenceFormat: Reference Id /subscriptions/000 is not formatted correctly.",
		},
		{
			name:     "the message is empty and the reason is in the details",
			status:   http.StatusConflict,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"Conflict","message":"","details":[{"code":"AnotherOperationInProgress","message":"Another operation on this resource is in progress."}]}}`,
			expected: "unexpected status 409 (409 Conflict) with error: Conflict: \n  AnotherOperationInProgress: Another operation on this resource is in progress.",
		},
		{
			name:     "details inside details, each naming what it is about",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"InvalidTemplateDeployment","message":"The template deployment failed.","details":[{"code":"RequestDisallowedByPolicy","target":"examplestorage","message":"Resource 'examplestorage' was disallowed by policy.","details":[{"code":"LocationNotAllowed","message":"Location 'westus9' is not in the allowed list."}]}]}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: InvalidTemplateDeployment: The template deployment failed.\n  RequestDisallowedByPolicy: Resource 'examplestorage' was disallowed by policy. (target: examplestorage)\n    LocationNotAllowed: Location 'westus9' is not in the allowed list.",
		},
		{
			name:     "what the error is about",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"InvalidParameter","message":"The value is invalid.","target":"properties.sku.name"}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: InvalidParameter: The value is invalid.\n  target: properties.sku.name",
		},
		{
			name:     "the policy that denied the request",
			status:   http.StatusForbidden,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"RequestDisallowedByPolicy","message":"Resource 'examplestorage' was disallowed by policy.","additionalInfo":[{"type":"PolicyViolation","info":{"policyDefinitionDisplayName":"Secure transfer to storage accounts should be enabled"}}]}}`,
			expected: "unexpected status 403 (403 Forbidden) with error: RequestDisallowedByPolicy: Resource 'examplestorage' was disallowed by policy.\n  PolicyViolation: {\"policyDefinitionDisplayName\":\"Secure transfer to storage accounts should be enabled\"}",
		},
		{
			name:     "the reason is in fields of the inner error",
			status:   http.StatusConflict,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"AllocationFailed","message":"Allocation failed.","innererror":{"exceptiontype":"Microsoft.Compute.AllocationFailedException","errordetail":"Zone 2 has no capacity for Standard_D64s_v5."}}}`,
			expected: "unexpected status 409 (409 Conflict) with error: AllocationFailed: Allocation failed.\n  innererror:\n    errordetail: Zone 2 has no capacity for Standard_D64s_v5.\n    exceptiontype: Microsoft.Compute.AllocationFailedException",
		},
		{
			name:     "fields beside the code and message",
			status:   http.StatusBadRequest,
			headers:  map[string]string{"Content-Type": "application/json"},
			body:     `{"error":{"code":"BadRequest","message":"The request was invalid.","possibleCauses":"The name is already in use.","recommendedAction":"Choose another name.","retryable":false}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: BadRequest: The request was invalid.\n  possibleCauses: The name is already in use.\n  recommendedAction: Choose another name.\n  retryable: false",
		},

		// the response headers that identify the request, or explain the failure when there is no body to
		{
			name:   "the IDs that identify the request to Resource Manager",
			status: http.StatusBadRequest,
			headers: map[string]string{
				"Content-Type":                "application/json",
				"x-ms-request-id":             "11111111-aaaa-bbbb-cccc-000000000001",
				"x-ms-correlation-request-id": "22222222-aaaa-bbbb-cccc-000000000002",
			},
			body:     `{"error":{"code":"InvalidParameter","message":"The value of parameter sku is invalid."}}`,
			expected: "unexpected status 400 (400 Bad Request) with error: InvalidParameter: The value of parameter sku is invalid.\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001\nx-ms-correlation-request-id: 22222222-aaaa-bbbb-cccc-000000000002",
		},
		{
			name:   "the IDs that identify the request to Microsoft Graph",
			status: http.StatusForbidden,
			headers: map[string]string{
				"Content-Type":      "application/json",
				"request-id":        "33333333-aaaa-bbbb-cccc-000000000003",
				"client-request-id": "44444444-aaaa-bbbb-cccc-000000000004",
			},
			body:     `{"error":{"code":"Authorization_RequestDenied","message":"Insufficient privileges to complete the operation."}}`,
			expected: "unexpected status 403 (403 Forbidden) with error: Authorization_RequestDenied: Insufficient privileges to complete the operation.\nrequest-id: 33333333-aaaa-bbbb-cccc-000000000003\nclient-request-id: 44444444-aaaa-bbbb-cccc-000000000004",
		},
		{
			name:   "an ID is only included once",
			status: http.StatusNotFound,
			headers: map[string]string{
				"Content-Type":                "application/json",
				"request-id":                  "33333333-aaaa-bbbb-cccc-000000000003",
				"client-request-id":           "33333333-aaaa-bbbb-cccc-000000000003",
				"x-ms-request-id":             "44444444-aaaa-bbbb-cccc-000000000004",
				"x-ms-correlation-request-id": "44444444-aaaa-bbbb-cccc-000000000004",
			},
			body:     `{"error":{"code":"Request_ResourceNotFound","message":"Resource does not exist.","innerError":{"request-id":"33333333-aaaa-bbbb-cccc-000000000003"}}}`,
			expected: "unexpected status 404 (404 Not Found) with error: Request_ResourceNotFound: Resource does not exist.\n  innerError:\n    request-id: 33333333-aaaa-bbbb-cccc-000000000003\nx-ms-request-id: 44444444-aaaa-bbbb-cccc-000000000004",
		},
		{
			name:   "no body, and the reason is in the challenge",
			status: http.StatusUnauthorized,
			headers: map[string]string{
				"WWW-Authenticate": `Bearer error="insufficient_claims", error_description="Continuous access evaluation resulted in challenge"`,
				"x-ms-request-id":  "11111111-aaaa-bbbb-cccc-000000000001",
			},
			expected: "unexpected status 401 (401 Unauthorized) received with no body\nWWW-Authenticate: Bearer error=\"insufficient_claims\", error_description=\"Continuous access evaluation resulted in challenge\"\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001",
		},
		{
			name:   "no body, and the reason is in the error code header",
			status: http.StatusConflict,
			headers: map[string]string{
				"x-ms-error-code": "ContainerBeingDeleted",
				"x-ms-request-id": "11111111-aaaa-bbbb-cccc-000000000001",
			},
			expected: "unexpected status 409 (409 Conflict) received with no body\nx-ms-error-code: ContainerBeingDeleted\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001",
		},
		{
			name:   "a body, so the headers that would repeat it are left out",
			status: http.StatusUnauthorized,
			headers: map[string]string{
				"Content-Type":     "application/json",
				"WWW-Authenticate": `Bearer authorization_uri="https://login.microsoftonline.com/000"`,
				"x-ms-error-code":  "InvalidAuthenticationToken",
			},
			body:     `{"error":{"code":"InvalidAuthenticationToken","message":"The access token is invalid."}}`,
			expected: "unexpected status 401 (401 Unauthorized) with error: InvalidAuthenticationToken: The access token is invalid.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, v := range tc.headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer server.Close()

			c := NewClient(server.URL, "testService", "v1.0")
			req, err := c.NewRequest(ctx, RequestOptions{
				ContentType:         "application/json",
				ExpectedStatusCodes: []int{http.StatusOK},
				HttpMethod:          http.MethodGet,
				Path:                "/test",
			})
			if err != nil {
				t.Fatalf("NewRequest error: %v", err)
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

// Callers read the response of a failed request for themselves to pull out what they need, as the AzureRM provider does when validating a
// resource before creating it, so the body has to come back from Execute unread however the error was built.
func TestExecute_UnexpectedStatus_LeavesTheBodyForTheCaller(t *testing.T) {
	testCases := []struct {
		name        string
		contentType string
		body        string
		errorParser ResponseErrorParser
	}{
		{
			name:        "an error that was parsed",
			contentType: "application/json",
			body:        `{"error":{"code":"InvalidTemplateDeployment","message":"The template deployment failed.","details":[{"code":"SkuNotAvailable","message":"The requested size is not available in location 'westus9'."}]}}`,
		},
		{
			name:        "JSON that is not an error",
			contentType: "application/json",
			body:        `{"status":"Failed","reason":"The requested size is not available in location 'westus9'."}`,
		},
		{
			name:        "a body that is not JSON",
			contentType: "text/plain",
			body:        "The requested size is not available in location 'westus9'.",
		},
		{
			name:        "a body that a custom error parser has read",
			contentType: "application/xml",
			body:        "<Error><Code>SkuNotAvailable</Code></Error>",
			errorParser: readingErrorParser{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer server.Close()

			c := NewClient(server.URL, "testService", "v1.0")
			req, err := c.NewRequest(ctx, RequestOptions{
				ContentType:         "application/json",
				ExpectedStatusCodes: []int{http.StatusOK},
				HttpMethod:          http.MethodPost,
				Path:                "/validate",
			})
			if err != nil {
				t.Fatalf("NewRequest error: %v", err)
			}
			req.CustomErrorParser = tc.errorParser

			resp, err := req.Execute(ctx)
			if err == nil {
				t.Fatal("expected an error for an unexpected status")
			}
			if resp == nil || resp.Response == nil {
				t.Fatal("expected the response to be returned alongside the error")
			}

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

// readingErrorParser reads the whole body to make its error and does not put it back.
type readingErrorParser struct{}

func (readingErrorParser) FromResponse(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("parsed %d bytes", len(body))
}

func TestExecute_ReturnsErrorFromRetryFunc(t *testing.T) {
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error":{"code":"MissingSubscriptionRegistration","message":"The subscription is not registered to use namespace 'Microsoft.Foo'."}}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")
	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	expected := errors.New("the Resource Provider was not registered")
	req.RetryFunc = func(_ *http.Response, _ *odata.OData) (bool, error) {
		return false, expected
	}

	resp, err := req.Execute(ctx)
	if !errors.Is(err, expected) {
		t.Fatalf("expected the error from the RetryFunc, got: %v", err)
	}
	if resp == nil || resp.Response == nil || resp.StatusCode != http.StatusConflict {
		t.Fatalf("expected the 409 response to be returned alongside the error, got: %+v", resp)
	}
}

func TestExecute_IgnoresErrorFromRetryFuncWhenTheStatusIsExpected(t *testing.T) {
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")
	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	req.RetryFunc = func(_ *http.Response, _ *odata.OData) (bool, error) {
		return false, errors.New("not relevant to a response that was wanted")
	}

	if _, err := req.Execute(ctx); err != nil {
		t.Fatalf("expected no error for an expected status, got: %v", err)
	}
}

func TestExecute_SaysWhatWasBeingRetriedWhenTheContextEnds(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"error":{"code":"SubscriptionRequestsThrottled","message":"Number of 'read' requests exceeded."}}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")
	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	// end the context once the response has arrived, which is during the 30 second wait to retry
	req.RetryFunc = func(_ *http.Response, _ *odata.OData) (bool, error) {
		time.AfterFunc(50*time.Millisecond, cancel)
		return true, nil
	}

	_, err = req.Execute(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected the context's error, got: %v", err)
	}

	expected := fmt.Sprintf("Get %q: context canceled (while retrying; the last response received was 429 Too Many Requests: SubscriptionRequestsThrottled: Number of 'read' requests exceeded.)", server.URL+"/test")
	if err.Error() != expected {
		t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
	}
}

func TestExecute_KeepsTheReasonAWriteGotNoResponse(t *testing.T) {
	// the retries this turns on for a write that got no response would hide the error being tested
	t.Setenv("ARM_FIVEPOINTZERO_BETA", "")

	ctx := context.Background()
	cause := errors.New("connect: connection refused")

	c := NewClient("http://localhost", "testService", "v1.0")
	c.SetTransport(&roundTripperMock{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			return nil, cause
		},
	})

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodPut,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	_, err = req.Execute(ctx)
	if !errors.Is(err, cause) {
		t.Fatalf("expected the error to wrap the cause, got: %v", err)
	}

	expected := `Put "http://localhost/test": HTTP response was nil; connection may have been reset: connect: connection refused`
	if err.Error() != expected {
		t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
	}
}

func TestExecute_ReportsTheReasonAReadGotNoResponse(t *testing.T) {
	// the retries this turns on for a request that got no response would hide the error being tested
	t.Setenv("ARM_FIVEPOINTZERO_BETA", "")

	ctx := context.Background()
	cause := errors.New("dial tcp 127.0.0.1:443: connect: connection refused")

	c := NewClient("http://localhost", "testService", "v1.0")
	c.SetTransport(&roundTripperMock{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			return nil, cause
		},
	})

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	_, err = req.Execute(ctx)
	if !errors.Is(err, cause) {
		t.Fatalf("expected the error to wrap the cause, got: %v", err)
	}

	expected := `Get "http://localhost/test": dial tcp 127.0.0.1:443: connect: connection refused`
	if err.Error() != expected {
		t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
	}
}

type failingReader struct {
	err error
}

func (r failingReader) Read(_ []byte) (int, error) {
	return 0, r.err
}

func TestExecute_KeepsTheReasonTheBodyCouldNotBeRead(t *testing.T) {
	ctx := context.Background()
	cause := errors.New("connection reset by peer")

	c := NewClient("http://localhost", "testService", "v1.0")
	c.SetTransport(&roundTripperMock{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				Status:     "400 Bad Request",
				StatusCode: http.StatusBadRequest,
				Header:     http.Header{"Content-Type": []string{"text/plain"}, "X-Ms-Request-Id": []string{"11111111-aaaa-bbbb-cccc-000000000001"}},
				Body:       io.NopCloser(failingReader{err: cause}),
				Request:    req,
			}, nil
		},
	})

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	_, err = req.Execute(ctx)
	if err == nil {
		t.Fatal("expected an error for an unexpected status")
	}

	expected := "unexpected status 400 (400 Bad Request), could not read response body: connection reset by peer\nx-ms-request-id: 11111111-aaaa-bbbb-cccc-000000000001"
	if err.Error() != expected {
		t.Fatalf("expected error:\n%s\n\ngot:\n%s", expected, err.Error())
	}
}

func TestUnmarshal_KeepsTheReasonTheBodyCouldNotBeRead(t *testing.T) {
	cause := errors.New("connection reset by peer")

	for _, contentType := range []string{"application/xml", "application/octet-stream"} {
		resp := &Response{
			Response: &http.Response{
				StatusCode:    http.StatusOK,
				ContentLength: -1,
				Header:        http.Header{"Content-Type": []string{contentType}},
				Body:          io.NopCloser(failingReader{err: cause}),
				Request:       &http.Request{Header: http.Header{}},
			},
		}

		var out []byte
		err := resp.Unmarshal(&out)
		if err == nil {
			t.Fatalf("%s: expected an error", contentType)
		}

		expected := "could not parse response body: connection reset by peer"
		if err.Error() != expected {
			t.Fatalf("%s: expected error %q, got %q", contentType, expected, err.Error())
		}
	}
}
