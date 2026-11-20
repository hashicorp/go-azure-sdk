package dataplane

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
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

	customTransport := &http.Transport{}
	original.SetTransport(customTransport)

	cloned := original.CloneClient(newEndpoint)

	if cloned == original {
		t.Fatal("cloned client should not be the same pointer as original")
	}

	if cloned.Transport != customTransport {
		t.Errorf("expected cloned Transport to preserve original Transport")
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

func TestCloneClient_Concurrent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	base, err := NewClient("https://placeholder.invalid", "service", "2023-01-01")
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	base.AuthorizeRequest = nil
	base.DisableRetries = true
	base.AppendRequestMiddleware(func(r *http.Request) (*http.Request, error) {
		r.Header.Set("X-Test-Middleware", "1")
		return r, nil
	})
	base.AppendResponseMiddleware(func(_ *http.Request, resp *http.Response) (*http.Response, error) {
		return resp, nil
	})

	const workers = 16
	const requestsPerWorker = 5
	var wg sync.WaitGroup
	errs := make(chan error, workers*requestsPerWorker)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cloned := base.CloneClient(srv.URL)
			for j := 0; j < requestsPerWorker; j++ {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				req, err := cloned.NewRequest(ctx, client.RequestOptions{
					ContentType:         "application/json",
					ExpectedStatusCodes: []int{http.StatusOK},
					HttpMethod:          http.MethodGet,
					Path:                "/test",
				})
				if err != nil {
					cancel()
					errs <- err
					return
				}
				resp, err := req.Execute(ctx)
				if err != nil {
					cancel()
					errs <- err
					return
				}
				var out map[string]bool
				err = resp.Unmarshal(&out)
				cancel()
				if err != nil {
					errs <- err
					return
				}
			}
		}()
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Fatalf("concurrent CloneClient/Execute failed: %v", err)
		}
	}
}
