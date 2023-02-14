package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccClient(t *testing.T) {
	test.AccTest(t)

	ctx := context.TODO()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.MicrosoftGraph
	endpoint, ok := api.Endpoint()
	if !ok {
		t.Fatalf("missing endpoint for microsoft graph for this environment")
	}
	conn.Authorize(ctx, t, api)

	c := NewClient(*endpoint, "example", "2020-01-01")
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
	req, err := c.NewRequest(ctx, reqOpts)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := req.ExecutePaged(ctx)
	if err != nil {
		t.Fatalf("ExecutePaged(): %v", err)
	}

	// test resp unmarshal as json
	resp.Header.Set("Content-Type", "application/json")
	var resJSON interface{}
	if err = resp.Unmarshal(&resJSON); err != nil {
		t.Fatal(err)
	}

	// unmarshal as *[]byte
	resp.Header.Set("Content-Type", "application/octet-stream")
	var bs []byte
	// pass slice itself will not work
	if err = resp.Unmarshal(bs); err == nil {
		t.Fatalf("response []byte should raise an error")
	}

	if err = resp.Unmarshal(&bs); err != nil {
		t.Fatal(err)
	}
	var bsPtr *[]byte
	if err = resp.Unmarshal(&bsPtr); err != nil {
		t.Fatal(err)
	}

	// unmarshal to XML should raise an error
	resp.Header.Set("Content-Type", "text/xml")
	var resXML interface{}
	if err = resp.Unmarshal(&resXML); err == nil {
		t.Fatal("should raise an error when unmarshalling to XML")
	}

}
