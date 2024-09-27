package comanageddevicewindowsprotectionstate

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

type UpdateComanagedDeviceWindowsProtectionStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateComanagedDeviceWindowsProtectionStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateComanagedDeviceWindowsProtectionStateOperationOptions() UpdateComanagedDeviceWindowsProtectionStateOperationOptions {
	return UpdateComanagedDeviceWindowsProtectionStateOperationOptions{}
}

func (o UpdateComanagedDeviceWindowsProtectionStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateComanagedDeviceWindowsProtectionStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateComanagedDeviceWindowsProtectionStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateComanagedDeviceWindowsProtectionState - Update the navigation property windowsProtectionState in
// deviceManagement
func (c ComanagedDeviceWindowsProtectionStateClient) UpdateComanagedDeviceWindowsProtectionState(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input beta.WindowsProtectionState, options UpdateComanagedDeviceWindowsProtectionStateOperationOptions) (result UpdateComanagedDeviceWindowsProtectionStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/windowsProtectionState", id.ID()),
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
