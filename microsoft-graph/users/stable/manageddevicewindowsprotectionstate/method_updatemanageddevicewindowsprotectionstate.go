package manageddevicewindowsprotectionstate

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

type UpdateManagedDeviceWindowsProtectionStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateManagedDeviceWindowsProtectionStateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateManagedDeviceWindowsProtectionStateOperationOptions() UpdateManagedDeviceWindowsProtectionStateOperationOptions {
	return UpdateManagedDeviceWindowsProtectionStateOperationOptions{}
}

func (o UpdateManagedDeviceWindowsProtectionStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateManagedDeviceWindowsProtectionStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateManagedDeviceWindowsProtectionStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateManagedDeviceWindowsProtectionState - Update the navigation property windowsProtectionState in users
func (c ManagedDeviceWindowsProtectionStateClient) UpdateManagedDeviceWindowsProtectionState(ctx context.Context, id stable.UserIdManagedDeviceId, input stable.WindowsProtectionState, options UpdateManagedDeviceWindowsProtectionStateOperationOptions) (result UpdateManagedDeviceWindowsProtectionStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/windowsProtectionState", id.ID()),
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
