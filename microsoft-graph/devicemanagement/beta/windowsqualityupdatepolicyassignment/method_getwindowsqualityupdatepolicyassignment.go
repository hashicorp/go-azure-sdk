package windowsqualityupdatepolicyassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWindowsQualityUpdatePolicyAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WindowsQualityUpdatePolicyAssignment
}

type GetWindowsQualityUpdatePolicyAssignmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetWindowsQualityUpdatePolicyAssignmentOperationOptions() GetWindowsQualityUpdatePolicyAssignmentOperationOptions {
	return GetWindowsQualityUpdatePolicyAssignmentOperationOptions{}
}

func (o GetWindowsQualityUpdatePolicyAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetWindowsQualityUpdatePolicyAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetWindowsQualityUpdatePolicyAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetWindowsQualityUpdatePolicyAssignment - Get assignments from deviceManagement. List of the groups this profile is
// assgined to.
func (c WindowsQualityUpdatePolicyAssignmentClient) GetWindowsQualityUpdatePolicyAssignment(ctx context.Context, id beta.DeviceManagementWindowsQualityUpdatePolicyIdAssignmentId, options GetWindowsQualityUpdatePolicyAssignmentOperationOptions) (result GetWindowsQualityUpdatePolicyAssignmentOperationResponse, err error) {
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

	var model beta.WindowsQualityUpdatePolicyAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
