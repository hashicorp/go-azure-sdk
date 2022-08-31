package resourcemanager_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	resourcemanager2 "github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"net/http"
	"testing"
)

func TestNewPoller(t *testing.T) {
	resp := &client.Response{
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
	}

	poller, err := resourcemanager2.NewPollerFromResponse(context.TODO(), resp, resourcemanager2.NewResourceManagerClient(""))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", *poller)
}
