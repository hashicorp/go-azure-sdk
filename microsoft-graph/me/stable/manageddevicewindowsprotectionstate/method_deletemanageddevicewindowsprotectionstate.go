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

type DeleteManagedDeviceWindowsProtectionStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteManagedDeviceWindowsProtectionStateOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteManagedDeviceWindowsProtectionStateOperationOptions() DeleteManagedDeviceWindowsProtectionStateOperationOptions {
	return DeleteManagedDeviceWindowsProtectionStateOperationOptions{}
}

func (o DeleteManagedDeviceWindowsProtectionStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteManagedDeviceWindowsProtectionStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteManagedDeviceWindowsProtectionStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteManagedDeviceWindowsProtectionState - Delete navigation property windowsProtectionState for me
func (c ManagedDeviceWindowsProtectionStateClient) DeleteManagedDeviceWindowsProtectionState(ctx context.Context, id stable.MeManagedDeviceId, options DeleteManagedDeviceWindowsProtectionStateOperationOptions) (result DeleteManagedDeviceWindowsProtectionStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/windowsProtectionState", id.ID()),
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
