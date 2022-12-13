package resourcemanager_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
)

func TestNewPoller(t *testing.T) {
	testCases := []*client.Response{
		{
			Response: &http.Response{
				Status:     http.StatusText(http.StatusCreated),
				StatusCode: http.StatusCreated,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,

				Header: http.Header{
					http.CanonicalHeaderKey("Azure-AsyncOperation"): []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/5678"},
				},

				Request: nil,
			},
		},
		{
			Response: &http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,

				Header: http.Header{
					http.CanonicalHeaderKey("Location"): []string{"https://async-url-test.local/subscriptions/1234/providers/foo/operations/6789"},
				},

				Request: nil,
			},
		},
	}

	for _, resp := range testCases {
		localApi := environments.NewApiEndpoint("Example", "https://async-url-test.local", nil)
		client, err := resourcemanager.NewResourceManagerClient(localApi, "example", "2020-02-01")
		if err != nil {
			t.Fatalf("building client: %+v", err)
		}
		if _, err = resourcemanager.PollerFromResponse(resp, client); err != nil {
			t.Fatal(fmt.Errorf("building poller from response: %+v", err))
		}
	}
}
