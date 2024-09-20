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

type ShutDownManagedDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ShutDownManagedDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultShutDownManagedDeviceOperationOptions() ShutDownManagedDeviceOperationOptions {
	return ShutDownManagedDeviceOperationOptions{}
}

func (o ShutDownManagedDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ShutDownManagedDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ShutDownManagedDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ShutDownManagedDevice - Invoke action shutDown. Shut down device
func (c ManagedDeviceClient) ShutDownManagedDevice(ctx context.Context, id stable.MeManagedDeviceId, options ShutDownManagedDeviceOperationOptions) (result ShutDownManagedDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/shutDown", id.ID()),
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
