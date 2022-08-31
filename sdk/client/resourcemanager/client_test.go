package resourcemanager_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"testing"
)

type options struct{}

func (o options) ToHeaders() *resourcemanager.Headers {
	return nil
}

func (o options) ToOData() *odata.Query {
	return nil
}

func (o options) ToQuery() *resourcemanager.QueryParams {
	return nil
}

func TestNewGetRequest(t *testing.T) {
	client := resourcemanager.NewResourceManagerClient("http://localhost")
	req, err := client.NewGetRequest(context.TODO(), "/", "2022-02-22", options{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", *req)
}
