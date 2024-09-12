package windowsdriverupdateprofileassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWindowsDriverUpdateProfileAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WindowsDriverUpdateProfileAssignment
}

type GetWindowsDriverUpdateProfileAssignmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetWindowsDriverUpdateProfileAssignmentOperationOptions() GetWindowsDriverUpdateProfileAssignmentOperationOptions {
	return GetWindowsDriverUpdateProfileAssignmentOperationOptions{}
}

func (o GetWindowsDriverUpdateProfileAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetWindowsDriverUpdateProfileAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetWindowsDriverUpdateProfileAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetWindowsDriverUpdateProfileAssignment - Get assignments from deviceManagement. The list of group assignments of the
// profile.
func (c WindowsDriverUpdateProfileAssignmentClient) GetWindowsDriverUpdateProfileAssignment(ctx context.Context, id beta.DeviceManagementWindowsDriverUpdateProfileIdAssignmentId, options GetWindowsDriverUpdateProfileAssignmentOperationOptions) (result GetWindowsDriverUpdateProfileAssignmentOperationResponse, err error) {
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

	var model beta.WindowsDriverUpdateProfileAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
