package importedwindowsautopilotdeviceidentity

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

type DeleteImportedWindowsAutopilotDeviceIdentityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteImportedWindowsAutopilotDeviceIdentityOperationOptions() DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions {
	return DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions{}
}

func (o DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteImportedWindowsAutopilotDeviceIdentity - Delete importedWindowsAutopilotDeviceIdentity. Deletes a
// importedWindowsAutopilotDeviceIdentity.
func (c ImportedWindowsAutopilotDeviceIdentityClient) DeleteImportedWindowsAutopilotDeviceIdentity(ctx context.Context, id stable.DeviceManagementImportedWindowsAutopilotDeviceIdentityId, options DeleteImportedWindowsAutopilotDeviceIdentityOperationOptions) (result DeleteImportedWindowsAutopilotDeviceIdentityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
