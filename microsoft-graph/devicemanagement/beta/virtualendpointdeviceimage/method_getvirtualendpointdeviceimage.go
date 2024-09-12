package virtualendpointdeviceimage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointDeviceImageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCDeviceImage
}

type GetVirtualEndpointDeviceImageOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetVirtualEndpointDeviceImageOperationOptions() GetVirtualEndpointDeviceImageOperationOptions {
	return GetVirtualEndpointDeviceImageOperationOptions{}
}

func (o GetVirtualEndpointDeviceImageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointDeviceImageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetVirtualEndpointDeviceImageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointDeviceImage - Get cloudPcDeviceImage. Read the properties and relationships of a specific
// cloudPcDeviceImage object.
func (c VirtualEndpointDeviceImageClient) GetVirtualEndpointDeviceImage(ctx context.Context, id beta.DeviceManagementVirtualEndpointDeviceImageId, options GetVirtualEndpointDeviceImageOperationOptions) (result GetVirtualEndpointDeviceImageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.CloudPCDeviceImage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
