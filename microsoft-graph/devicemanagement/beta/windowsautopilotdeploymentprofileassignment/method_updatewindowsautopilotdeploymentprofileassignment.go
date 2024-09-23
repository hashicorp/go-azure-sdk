package windowsautopilotdeploymentprofileassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsAutopilotDeploymentProfileAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions() UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions {
	return UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions{}
}

func (o UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsAutopilotDeploymentProfileAssignment - Update the navigation property assignments in deviceManagement
func (c WindowsAutopilotDeploymentProfileAssignmentClient) UpdateWindowsAutopilotDeploymentProfileAssignment(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeploymentProfileIdAssignmentId, input beta.WindowsAutopilotDeploymentProfileAssignment, options UpdateWindowsAutopilotDeploymentProfileAssignmentOperationOptions) (result UpdateWindowsAutopilotDeploymentProfileAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
