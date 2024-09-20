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

type RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions() RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions {
	return RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions{}
}

func (o RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveComanagedDeviceFirmwareConfigurationInterfaceManagement - Invoke action
// removeDeviceFirmwareConfigurationInterfaceManagement. Remove device from Device Firmware Configuration Interface
// management
func (c ComanagedDeviceClient) RemoveComanagedDeviceFirmwareConfigurationInterfaceManagement(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions) (result RemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/removeDeviceFirmwareConfigurationInterfaceManagement", id.ID()),
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
