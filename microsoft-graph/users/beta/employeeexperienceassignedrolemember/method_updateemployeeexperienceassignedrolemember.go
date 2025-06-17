package employeeexperienceassignedrolemember

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEmployeeExperienceAssignedRoleMemberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEmployeeExperienceAssignedRoleMemberOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEmployeeExperienceAssignedRoleMemberOperationOptions() UpdateEmployeeExperienceAssignedRoleMemberOperationOptions {
	return UpdateEmployeeExperienceAssignedRoleMemberOperationOptions{}
}

func (o UpdateEmployeeExperienceAssignedRoleMemberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEmployeeExperienceAssignedRoleMemberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEmployeeExperienceAssignedRoleMemberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEmployeeExperienceAssignedRoleMember - Update the navigation property members in users
func (c EmployeeExperienceAssignedRoleMemberClient) UpdateEmployeeExperienceAssignedRoleMember(ctx context.Context, id beta.UserIdEmployeeExperienceAssignedRoleIdMemberId, input beta.EngagementRoleMember, options UpdateEmployeeExperienceAssignedRoleMemberOperationOptions) (result UpdateEmployeeExperienceAssignedRoleMemberOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
