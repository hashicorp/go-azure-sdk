package resourcemanager_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/client"
	"github.com/hashicorp/go-azure-sdk/client/resourcemanager"
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

	poller, err := resourcemanager.NewPollerFromResponse(context.TODO(), resp, resourcemanager.NewResourceManagerClient(""))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", *poller)
}
