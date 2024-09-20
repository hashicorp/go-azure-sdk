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

type LocateComanagedDeviceDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type LocateComanagedDeviceDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultLocateComanagedDeviceDeviceOperationOptions() LocateComanagedDeviceDeviceOperationOptions {
	return LocateComanagedDeviceDeviceOperationOptions{}
}

func (o LocateComanagedDeviceDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o LocateComanagedDeviceDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o LocateComanagedDeviceDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// LocateComanagedDeviceDevice - Invoke action locateDevice. Locate a device
func (c ComanagedDeviceClient) LocateComanagedDeviceDevice(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options LocateComanagedDeviceDeviceOperationOptions) (result LocateComanagedDeviceDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/locateDevice", id.ID()),
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
