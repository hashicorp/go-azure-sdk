package manageddevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RebootManagedDeviceNowOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RebootManagedDeviceNowOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRebootManagedDeviceNowOperationOptions() RebootManagedDeviceNowOperationOptions {
	return RebootManagedDeviceNowOperationOptions{}
}

func (o RebootManagedDeviceNowOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RebootManagedDeviceNowOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RebootManagedDeviceNowOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RebootManagedDeviceNow - Invoke action rebootNow. Reboot device
func (c ManagedDeviceClient) RebootManagedDeviceNow(ctx context.Context, id stable.MeManagedDeviceId, options RebootManagedDeviceNowOperationOptions) (result RebootManagedDeviceNowOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/rebootNow", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
