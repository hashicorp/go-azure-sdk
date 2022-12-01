package client

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
	"net/http"
	"testing"
)

func TestAccClient(t *testing.T) {
	test.AccTest(t)

	conn := test.NewConnection(t)
	conn.Authorize(context.TODO(), conn.AuthConfig.Environment.MSGraph)

	c := NewClient(conn.AuthConfig.Environment.MSGraph.Endpoint)
	c.Authorizer = conn.Authorizer

	path := fmt.Sprintf("/v1.0/servicePrincipals/%s", conn.Claims.ObjectId)
	reqOpts := RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: nil,
		Path:          path,
	}
	req, err := c.NewRequest(context.TODO(), reqOpts)
	if err != nil {
		t.Fatal(err)
	}

	_, err = req.ExecutePaged()
	if err != nil {
		t.Fatalf("ExecutePaged(): %v", err)
	}
}
