package dataplane

import (
	"net/http"
	"testing"
)

func TestCloneClient(t *testing.T) {
	originalEndpoint := "https://original.endpoint"
	newEndpoint := "https://new.endpoint"
	serviceName := "service"
	apiVersion := "2023-01-01"

	original, err := NewClient(originalEndpoint, serviceName, apiVersion)
	if err != nil {
		t.Fatalf("failed to create original client: %v", err)
	}

	// Add some middlewares to test deep copy
	original.AppendRequestMiddleware(func(request *http.Request) (*http.Request, error) {
		// ensure the `X-Correlation-ID` field is set
		if request.Header.Get("x-ms-correlation-request-id") == "" {
			request.Header.Add("x-ms-correlation-request-id", "test_id")
		}
		return request, nil
	})

	original.AppendResponseMiddleware(func(request *http.Request, response *http.Response) (*http.Response, error) {
		// ensure the `X-Correlation-ID` field is set
		if response.Header.Get("x-ms-correlation-request-id") == "" {
			response.Header.Add("x-ms-correlation-request-id", "test_id")
		}
		return response, nil
	})

	cloned := original.CloneClient(newEndpoint)

	if cloned == original {
		t.Fatal("cloned client should not be the same pointer as original")
	}

	if cloned.Client == original.Client {
		t.Fatal("cloned.Client should not be the same pointer as original.Client")
	}

	if cloned.BaseUri != newEndpoint {
		t.Errorf("expected BaseUri %s, got %s", newEndpoint, cloned.BaseUri)
	}

	if cloned.ApiVersion != original.ApiVersion {
		t.Errorf("expected ApiVersion %s, got %s", original.ApiVersion, cloned.ApiVersion)
	}

	if cloned.UserAgent != original.UserAgent {
		t.Errorf("expected UserAgent %s, got %s", original.UserAgent, cloned.UserAgent)
	}

	// Verify deep copy of middlewares
	if cloned.RequestMiddlewares == original.RequestMiddlewares {
		t.Error("RequestMiddlewares should be deep copied (different pointers)")
	}
	if len(*cloned.RequestMiddlewares) != len(*original.RequestMiddlewares) {
		t.Errorf("expected %d RequestMiddlewares, got %d", len(*original.RequestMiddlewares), len(*cloned.RequestMiddlewares))
	}

	if cloned.ResponseMiddlewares == original.ResponseMiddlewares {
		t.Error("ResponseMiddlewares should be deep copied (different pointers)")
	}

	if len(*cloned.ResponseMiddlewares) != len(*original.ResponseMiddlewares) {
		t.Errorf("expected %d ResponseMiddlewares, got %d", len(*original.ResponseMiddlewares), len(*cloned.ResponseMiddlewares))
	}
}
