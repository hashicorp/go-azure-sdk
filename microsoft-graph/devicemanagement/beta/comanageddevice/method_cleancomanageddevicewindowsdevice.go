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

type CleanComanagedDeviceWindowsDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CleanComanagedDeviceWindowsDeviceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCleanComanagedDeviceWindowsDeviceOperationOptions() CleanComanagedDeviceWindowsDeviceOperationOptions {
	return CleanComanagedDeviceWindowsDeviceOperationOptions{}
}

func (o CleanComanagedDeviceWindowsDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CleanComanagedDeviceWindowsDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CleanComanagedDeviceWindowsDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CleanComanagedDeviceWindowsDevice - Invoke action cleanWindowsDevice. Clean Windows device
func (c ComanagedDeviceClient) CleanComanagedDeviceWindowsDevice(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input CleanComanagedDeviceWindowsDeviceRequest, options CleanComanagedDeviceWindowsDeviceOperationOptions) (result CleanComanagedDeviceWindowsDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/cleanWindowsDevice", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
