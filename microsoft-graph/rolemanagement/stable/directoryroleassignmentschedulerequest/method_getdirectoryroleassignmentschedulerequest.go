package directoryroleassignmentschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDirectoryRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleAssignmentScheduleRequest
}

type GetDirectoryRoleAssignmentScheduleRequestOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDirectoryRoleAssignmentScheduleRequestOperationOptions() GetDirectoryRoleAssignmentScheduleRequestOperationOptions {
	return GetDirectoryRoleAssignmentScheduleRequestOperationOptions{}
}

func (o GetDirectoryRoleAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDirectoryRoleAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleAssignmentScheduleRequest - Get unifiedRoleAssignmentScheduleRequest. In PIM, read the details of a
// request for an active and persistent role assignment made through the unifiedRoleAssignmentScheduleRequest object.
func (c DirectoryRoleAssignmentScheduleRequestClient) GetDirectoryRoleAssignmentScheduleRequest(ctx context.Context, id stable.RoleManagementDirectoryRoleAssignmentScheduleRequestId, options GetDirectoryRoleAssignmentScheduleRequestOperationOptions) (result GetDirectoryRoleAssignmentScheduleRequestOperationResponse, err error) {
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

	var model stable.UnifiedRoleAssignmentScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
