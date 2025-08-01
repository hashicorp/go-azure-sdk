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

type VirtualMachineExtensionImagesListVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualMachineExtensionImage
}

type VirtualMachineExtensionImagesListVersionsOperationOptions struct {
	Filter  *string
	Orderby *string
	Top     *int64
}

func DefaultVirtualMachineExtensionImagesListVersionsOperationOptions() VirtualMachineExtensionImagesListVersionsOperationOptions {
	return VirtualMachineExtensionImagesListVersionsOperationOptions{}
}

func (o VirtualMachineExtensionImagesListVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o VirtualMachineExtensionImagesListVersionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o VirtualMachineExtensionImagesListVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// VirtualMachineExtensionImagesListVersions ...
func (c ComputeRPSClient) VirtualMachineExtensionImagesListVersions(ctx context.Context, id TypeId, options VirtualMachineExtensionImagesListVersionsOperationOptions) (result VirtualMachineExtensionImagesListVersionsOperationResponse, err error) {
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

	var model []VirtualMachineExtensionImage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
