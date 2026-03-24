package policyinvoke

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPolicyInvokeMiddleware(t *testing.T) {
	// Mock policy API server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(acquirePolicyTokenResponse{
			Token: "mock-policy-token",
		})
	}))
	defer ts.Close()

	// Create middleware with mocked client
	middleware := NewPolicyInvokeMiddleware(PolicyInvokeOptions{
		ChangeReference: "my-change-ref",
		HTTPClient:      ts.Client(),
	})

	// Original request to ARM
	// We use the mock server's URL but with a subscription path
	armUrl := ts.URL + "/subscriptions/12345/resourceGroups/myRG/providers/Microsoft.Compute/virtualMachines/myVM?api-version=2021-03-01"
	originalBody := `{"location":"eastus"}`
	req, _ := http.NewRequest(http.MethodPut, armUrl, bytes.NewBufferString(originalBody))
	req.Header.Set("Authorization", "Bearer original-token")
	req.Header.Set("Content-Type", "application/json")

	// Run middleware
	modifiedReq, err := middleware(req)
	if err != nil {
		t.Fatalf("middleware failed: %v", err)
	}

	// Verify header was set
	token := modifiedReq.Header.Get(HeaderPolicyExternalEvaluations)
	if token != "mock-policy-token" {
		t.Errorf("expected mock-policy-token, got %s", token)
	}

	// Verify original body was preserved
	body, _ := io.ReadAll(modifiedReq.Body)
	if string(body) != originalBody {
		t.Errorf("body was not preserved: %s", string(body))
	}
}

func TestPolicyInvokeMiddleware_NoSubscription(t *testing.T) {
	middleware := NewPolicyInvokeMiddleware(PolicyInvokeOptions{})

	// Request without subscription ID
	req, _ := http.NewRequest(http.MethodGet, "https://management.azure.com/providers/Microsoft.Resources/tenants?api-version=2020-01-01", nil)

	modifiedReq, err := middleware(req)
	if err != nil {
		t.Fatalf("middleware failed: %v", err)
	}

	// Should not have the policy header
	token := modifiedReq.Header.Get(HeaderPolicyExternalEvaluations)
	if token != "" {
		t.Errorf("expected empty token, got %s", token)
	}
}

func TestPolicyInvokeMiddleware_PolicyFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	middleware := NewPolicyInvokeMiddleware(PolicyInvokeOptions{
		HTTPClient: ts.Client(),
	})

	armUrl := ts.URL + "/subscriptions/12345/resourceGroups/myRG"
	req, _ := http.NewRequest(http.MethodGet, armUrl, nil)

	_, err := middleware(req)
	if err == nil {
		t.Fatal("expected error from middleware when policy API fails")
	}
}
