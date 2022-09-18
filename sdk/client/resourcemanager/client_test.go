package resourcemanager_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
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
	_, err := client.NewGetRequest(context.TODO(), "/", "2022-02-22", options{})
	if err != nil {
		t.Fatal(err)
	}
}
