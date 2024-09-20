package enterpriseapproleeligibilityschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions() UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions {
	return UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions{}
}

func (o UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEnterpriseAppRoleEligibilityScheduleRequest - Update the navigation property roleEligibilityScheduleRequests in
// roleManagement
func (c EnterpriseAppRoleEligibilityScheduleRequestClient) UpdateEnterpriseAppRoleEligibilityScheduleRequest(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId, input beta.UnifiedRoleEligibilityScheduleRequest, options UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) (result UpdateEnterpriseAppRoleEligibilityScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
