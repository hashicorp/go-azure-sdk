// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalevaluation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

const missingToken403Payload = `{
  "error": {
    "code": "RequestDisallowedByPolicy",
    "message": "Resource 'myVM' was disallowed by policy.",
    "additionalInfo": [
      {
        "type": "PolicyViolation",
        "info": {
          "evaluationDetails": {
            "missingPolicyTokenDetails": [
              {
                "policyAssignmentId": "/subscriptions/12345/providers/Microsoft.Authorization/policyAssignments/ext-eval",
                "shouldDeny": true
              }
            ]
          }
        }
      }
    ]
  }
}`

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newHandlerTransport(handler http.HandlerFunc) (http.RoundTripper, *http.Client) {
	rt := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		rec := httptest.NewRecorder()
		handler(rec, req)
		resp := rec.Result()
		resp.Request = req
		return resp, nil
	})
	return rt, &http.Client{Transport: rt}
}

type trackingReadCloser struct {
	io.Reader
	closed bool
}

func (t *trackingReadCloser) Close() error {
	t.closed = true
	return nil
}

func TestPassThrough_GetAndHeadAndSuccessfulMutations(t *testing.T) {
	var armCalls atomic.Int32
	var tokenCalls atomic.Int32

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/acquirePolicyToken") {
			tokenCalls.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		armCalls.Add(1)
		if r.Method == http.MethodGet || r.Method == http.MethodHead {
			// Even if GET/HEAD returns 403 with missingPolicyTokenDetails, it must not be intercepted.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(missingToken403Payload))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	rt := NewRoundTripper(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	})

	for _, method := range []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete} {
		armURL := "https://management.azure.com/subscriptions/12345/resourceGroups/myRG?api-version=2021-04-01"
		req, err := http.NewRequest(method, armURL, nil)
		if err != nil {
			t.Fatalf("creating %s request: %v", method, err)
		}
		resp, err := rt.RoundTrip(req)
		if err != nil {
			t.Fatalf("%s RoundTrip failed: %v", method, err)
		}
		_ = resp.Body.Close()
	}

	if tokenCalls.Load() != 0 {
		t.Fatalf("expected 0 acquirePolicyToken calls, got %d", tokenCalls.Load())
	}
	if armCalls.Load() != 6 {
		t.Fatalf("expected 6 ARM calls, got %d", armCalls.Load())
	}
}

func TestPassThrough_403WithoutMissingPolicyTokenDetails(t *testing.T) {
	nonToken403 := `{"error":{"code":"RequestDisallowedByPolicy","message":"Static policy denial","additionalInfo":[{"type":"PolicyViolation","info":{"evaluationDetails":{}}}]}}`
	var tokenCalls atomic.Int32

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/acquirePolicyToken") {
			tokenCalls.Add(1)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(nonToken403))
	})

	rt := NewRoundTripper(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	})

	req, _ := http.NewRequest(http.MethodPut, "https://management.azure.com/subscriptions/12345/resourceGroups/myRG?api-version=2021-04-01", strings.NewReader(`{"location":"eastus"}`))
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("expected 403 Forbidden, got %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading 403 body: %v", err)
	}
	if string(body) != nonToken403 {
		t.Fatalf("expected preserved 403 body %q, got %q", nonToken403, string(body))
	}
	if tokenCalls.Load() != 0 {
		t.Fatalf("expected 0 acquirePolicyToken calls, got %d", tokenCalls.Load())
	}
}

func TestReactiveInterception_ByteExactContentAndVerbatimPoPToken(t *testing.T) {
	// Intentionally uncompacted JSON with extra whitespace, newlines, specific key ordering, and HTML characters (<, >, &)
	// that standard encoding/json.Marshal would compact and escape.
	rawTargetBody := "{\n  \"zebra\" :  \"a < b & c > d\" ,\n  \"alpha\": [ 1,   2, 3 ]\n}"
	expectedToken := "PoP eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.mock-signature"

	var armAttempt atomic.Int32
	var capturedTokenReqRaw []byte
	var capturedTokenReqHeaders http.Header
	var capturedTokenReqQuery string
	var capturedRetryBody []byte
	var capturedRetryTokenHeader string

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/subscriptions/12345/providers/Microsoft.Authorization/acquirePolicyToken" {
			capturedTokenReqQuery = r.URL.RawQuery
			capturedTokenReqHeaders = r.Header.Clone()
			capturedTokenReqRaw, _ = io.ReadAll(r.Body)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"result": "Succeeded",
				"token": "` + expectedToken + `",
				"tokenId": "tid-1",
				"expiration": "2026-11-12T00:00:00Z",
				"results": []
			}`))
			return
		}

		attempt := armAttempt.Add(1)
		if attempt == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(missingToken403Payload))
			return
		}

		capturedRetryTokenHeader = r.Header.Get(HeaderPolicyExternalEvaluations)
		capturedRetryBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"/subscriptions/12345/resourceGroups/myRG/providers/Microsoft.Compute/virtualMachines/myVM"}`))
	})

	rt := NewRoundTripper(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	})

	ctx := WithChangeReference(context.Background(), "ctx-change-ref-001")
	targetPathAndQuery := "/subscriptions/12345/resourceGroups/myRG/providers/Microsoft.Compute/virtualMachines/myVM?api-version=2023-03-01"
	bodyReader := &trackingReadCloser{Reader: strings.NewReader(rawTargetBody)}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, "https://management.azure.com"+targetPathAndQuery, bodyReader)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.Header.Set("Authorization", "Bearer arm-access-token")
	req.Header.Set("Content-Type", "application/json")

	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip failed: %v", err)
	}
	defer resp.Body.Close()

	if !bodyReader.closed {
		t.Errorf("expected original request body ReadCloser to be closed")
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 Created on retry, got %d", resp.StatusCode)
	}
	if armAttempt.Load() != 2 {
		t.Fatalf("expected 2 ARM attempts, got %d", armAttempt.Load())
	}

	// Verify acquirePolicyToken endpoint query and headers
	if capturedTokenReqQuery != "api-version=2025-11-01" {
		t.Errorf("expected api-version=2025-11-01, got %q", capturedTokenReqQuery)
	}
	if got := capturedTokenReqHeaders.Get(HeaderForceSync); got != "true" {
		t.Errorf("expected %s=true, got %q", HeaderForceSync, got)
	}
	if got := capturedTokenReqHeaders.Get("Content-Type"); got != "application/json" {
		t.Errorf("expected Content-Type=application/json, got %q", got)
	}
	if got := capturedTokenReqHeaders.Get("Authorization"); got != "Bearer arm-access-token" {
		t.Errorf("expected Authorization header to be forwarded, got %q", got)
	}

	// Verify byte-for-byte exactness of operation.content inside the raw acquirePolicyToken request payload
	expectedAcquireJSON := `{"operation":{"uri":"` + targetPathAndQuery + `","httpMethod":"PUT","content":` + rawTargetBody + `},"changeReference":"ctx-change-ref-001"}`
	if string(capturedTokenReqRaw) != expectedAcquireJSON {
		t.Errorf("acquirePolicyToken request body mismatch:\ngot:  %s\nwant: %s", string(capturedTokenReqRaw), expectedAcquireJSON)
	}

	// Verify byte-for-byte exactness of the retried ARM request body
	if string(capturedRetryBody) != rawTargetBody {
		t.Errorf("retried ARM request body mismatch:\ngot:  %q\nwant: %q", string(capturedRetryBody), rawTargetBody)
	}

	// Verify verbatim injection of the PoP token header
	if capturedRetryTokenHeader != expectedToken {
		t.Errorf("expected %s header %q, got %q", HeaderPolicyExternalEvaluations, expectedToken, capturedRetryTokenHeader)
	}
}

func TestReactiveInterception_ManagementGroupScopeAndDeleteOmitsContent(t *testing.T) {
	var capturedAcquirePath string
	var capturedAcquireRaw []byte
	var armCalls atomic.Int32

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/acquirePolicyToken") {
			capturedAcquirePath = r.URL.Path
			capturedAcquireRaw, _ = io.ReadAll(r.Body)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result":"Succeeded","token":"PoP mg-token"}`))
			return
		}

		if armCalls.Add(1) == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(missingToken403Payload))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	rt := NewRoundTripper(Options{
		Transport:       baseRT,
		HTTPClient:      httpClient,
		ChangeReference: "opt-change-ref",
	})

	mgTarget := "/providers/Microsoft.Management/managementGroups/my-mg-01/providers/Microsoft.Authorization/roleAssignments/ra-1?api-version=2022-04-01"
	req, _ := http.NewRequest(http.MethodDelete, "https://management.azure.com"+mgTarget, nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204 NoContent, got %d", resp.StatusCode)
	}

	expectedAcquirePath := "/providers/Microsoft.Management/managementGroups/my-mg-01/providers/Microsoft.Authorization/acquirePolicyToken"
	if capturedAcquirePath != expectedAcquirePath {
		t.Errorf("expected acquire path %q, got %q", expectedAcquirePath, capturedAcquirePath)
	}

	expectedBody := `{"operation":{"uri":"` + mgTarget + `","httpMethod":"DELETE"},"changeReference":"opt-change-ref"}`
	if string(capturedAcquireRaw) != expectedBody {
		t.Errorf("expected DELETE acquirePolicyToken payload without content:\ngot:  %s\nwant: %s", string(capturedAcquireRaw), expectedBody)
	}
}

func TestGuardedRetry_PreventsInfiniteLoop(t *testing.T) {
	var armCalls atomic.Int32
	var tokenCalls atomic.Int32

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/acquirePolicyToken") {
			tokenCalls.Add(1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result":"Succeeded","token":"PoP loop-test-token"}`))
			return
		}

		armCalls.Add(1)
		// Always return 403 with missingPolicyTokenDetails even on retry
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(missingToken403Payload))
	})

	rt := NewRoundTripper(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	})

	req, _ := http.NewRequest(http.MethodPut, "https://management.azure.com/subscriptions/12345/resourceGroups/myRG?api-version=2021-04-01", strings.NewReader(`{"location":"eastus"}`))
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("expected 403 Forbidden after single retry, got %d", resp.StatusCode)
	}
	if tokenCalls.Load() != 1 {
		t.Fatalf("expected exactly 1 acquirePolicyToken call, got %d", tokenCalls.Load())
	}
	if armCalls.Load() != 2 {
		t.Fatalf("expected exactly 2 ARM calls (initial + 1 retry), got %d", armCalls.Load())
	}
}

func TestAcquirePolicyToken_200WithoutTokenReturnsPolicyEvaluationError(t *testing.T) {
	var armCalls atomic.Int32

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/acquirePolicyToken") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"result": "Failed",
				"token": "",
				"results": [
					{
						"policyAssignmentId": "/subscriptions/12345/providers/Microsoft.Authorization/policyAssignments/pa1",
						"result": "Denied",
						"message": "External endpoint rejected deployment",
						"policyEvaluationDetails": {
							"reason": "Change ticket CR-999 is not approved"
						}
					}
				]
			}`))
			return
		}

		armCalls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(missingToken403Payload))
	})

	rt := NewRoundTripper(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	})

	req, _ := http.NewRequest(http.MethodPost, "https://management.azure.com/subscriptions/12345/resourceGroups/myRG/moveResources?api-version=2021-04-01", strings.NewReader(`{"resources":[]}`))
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatal("expected PolicyEvaluationError, got nil")
	}

	var evalErr *PolicyEvaluationError
	if !errors.As(err, &evalErr) {
		t.Fatalf("expected *PolicyEvaluationError, got %T: %v", err, err)
	}
	if !strings.Contains(evalErr.Error(), "External endpoint rejected deployment") || !strings.Contains(evalErr.Error(), "Change ticket CR-999 is not approved") {
		t.Errorf("unexpected error message: %s", evalErr.Error())
	}
	if armCalls.Load() != 1 {
		t.Errorf("expected ARM request not to be retried on policy denial, got %d ARM calls", armCalls.Load())
	}
}

func TestAcquirePolicyToken_HttpErrorAndRetryAfter(t *testing.T) {
	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/acquirePolicyToken") {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"error":{"code":"ServiceUnavailable","message":"Try again"}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(missingToken403Payload))
	})

	rt := NewRoundTripper(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	})

	req, _ := http.NewRequest(http.MethodPatch, "https://management.azure.com/subscriptions/12345/resourceGroups/myRG?api-version=2021-04-01", strings.NewReader(`{"tags":{}}`))
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatal("expected error when acquirePolicyToken returns 503")
	}
	if !strings.Contains(err.Error(), "503") || !strings.Contains(err.Error(), "retry-after: 1") {
		t.Errorf("expected status 503 and retry-after in error, got: %v", err)
	}
}

func TestClientIntegration_ResponseMiddleware(t *testing.T) {
	expectedToken := "PoP eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.client-middleware"
	rawPayload := []byte("{\n  \"location\":   \"westeurope\"\n}")
	var armCalls atomic.Int32
	var receivedContent string

	baseRT, httpClient := newHandlerTransport(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/acquirePolicyToken") {
			var reqObj struct {
				Operation struct {
					Content json.RawMessage `json:"content"`
				} `json:"operation"`
			}
			bodyBytes, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(bodyBytes, &reqObj)
			receivedContent = string(reqObj.Operation.Content)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"result":"Succeeded","token":"` + expectedToken + `"}`))
			return
		}

		if armCalls.Add(1) == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(missingToken403Payload))
			return
		}

		if got := r.Header.Get(HeaderPolicyExternalEvaluations); got != expectedToken {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"provisioningState":"Succeeded"}`))
	})

	c := client.NewClient("https://management.azure.com", "example", "2021-04-01")
	c.DisableRetries = true
	c.SetTransport(baseRT)
	c.AppendResponseMiddleware(NewMiddleware(Options{
		Transport:  baseRT,
		HTTPClient: httpClient,
	}))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := c.NewRequest(ctx, client.RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodPut,
		Path:                "/subscriptions/12345/resourceGroups/myRG",
	})
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(rawPayload))
	req.ContentLength = int64(len(rawPayload))

	resp, err := req.Execute(ctx)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}
	if receivedContent != string(rawPayload) {
		t.Errorf("expected byte-exact content %q, got %q", string(rawPayload), receivedContent)
	}
}
