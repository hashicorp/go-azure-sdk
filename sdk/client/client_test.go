package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/internal/test"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
)

func TestAccClient(t *testing.T) {
	test.AccTest(t)

	conn := test.NewConnection(t, auth.TokenVersion2)
	conn.Authorize(context.TODO(), conn.AuthConfig.Environment.MSGraph)

	c := NewClient(conn.AuthConfig.Environment.MSGraph.Endpoint)
	c.Authorizer = conn.Authorizer

	path := fmt.Sprintf("/v1.0/servicePrincipals/%s", conn.Claims.ObjectId)
	req, err := c.NewRequest(context.TODO(), http.MethodGet, "application/json", path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = req.ExecutePaged()
	if err != nil {
		t.Fatalf("ExecutePaged(): %v", err)
	}
}
