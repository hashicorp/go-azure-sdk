package comanageddevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateComanagedDeviceReenableOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateComanagedDeviceReenableOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateComanagedDeviceReenableOperationOptions() CreateComanagedDeviceReenableOperationOptions {
	return CreateComanagedDeviceReenableOperationOptions{}
}

func (o CreateComanagedDeviceReenableOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceReenableOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceReenableOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceReenable - Invoke action reenable
func (c ComanagedDeviceClient) CreateComanagedDeviceReenable(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options CreateComanagedDeviceReenableOperationOptions) (result CreateComanagedDeviceReenableOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reenable", id.ID()),
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
