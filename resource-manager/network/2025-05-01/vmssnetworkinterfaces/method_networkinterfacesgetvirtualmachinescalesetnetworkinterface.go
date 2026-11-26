package vmssnetworkinterfaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *NetworkInterface
}

type NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions struct {
	Expand *string
}

func DefaultNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions() NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions {
	return NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions{}
}

func (o NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

// NetworkInterfacesGetVirtualMachineScaleSetNetworkInterface ...
func (c VMSSNetworkInterfacesClient) NetworkInterfacesGetVirtualMachineScaleSetNetworkInterface(ctx context.Context, id commonids.VirtualMachineScaleSetNetworkInterfaceId, options NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions) (result NetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationResponse, err error) {
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

	var model NetworkInterface
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
