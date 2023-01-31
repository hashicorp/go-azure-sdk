package msgraph_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type options struct{}

func (o options) ToHeaders() *client.Headers {
	return nil
}

func (o options) ToOData() *odata.Query {
	return nil
}

func (o options) ToQuery() *client.QueryParams {
	return nil
}

func TestNewGetRequest(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
	defer cancel()

	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options{},
		Path:          "/applications",
	}
	localApi := environments.NewApiEndpoint("Example", "http://localhost", pointer.To("00000000-0000-000-0000-000000000000"))
	msGraphClient, err := msgraph.NewMsGraphClient(localApi, msgraph.Version10, "11111111-1111-1111-1111-111111111111")
	if err != nil {
		t.Fatalf("building client: %+v", err)
	}
	if _, err = msGraphClient.NewRequest(ctx, opts); err != nil {
		t.Fatalf("building new request: %+v", err)
	}
}
