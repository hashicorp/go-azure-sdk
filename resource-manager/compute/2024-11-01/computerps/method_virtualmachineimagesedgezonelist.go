package computerps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineImagesEdgeZoneListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualMachineImageResource
}

type VirtualMachineImagesEdgeZoneListOperationOptions struct {
	Expand  *string
	Orderby *string
	Top     *int64
}

func DefaultVirtualMachineImagesEdgeZoneListOperationOptions() VirtualMachineImagesEdgeZoneListOperationOptions {
	return VirtualMachineImagesEdgeZoneListOperationOptions{}
}

func (o VirtualMachineImagesEdgeZoneListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o VirtualMachineImagesEdgeZoneListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o VirtualMachineImagesEdgeZoneListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// VirtualMachineImagesEdgeZoneList ...
func (c ComputeRPSClient) VirtualMachineImagesEdgeZoneList(ctx context.Context, id OfferSkuId, options VirtualMachineImagesEdgeZoneListOperationOptions) (result VirtualMachineImagesEdgeZoneListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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

	var model []VirtualMachineImageResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
