package enterpriseapptransitiveroleassignmentappscope

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

type UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions() UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions {
	return UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions{}
}

func (o UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEnterpriseAppTransitiveRoleAssignmentAppScope - Update the navigation property appScope in roleManagement
func (c EnterpriseAppTransitiveRoleAssignmentAppScopeClient) UpdateEnterpriseAppTransitiveRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId, input beta.AppScope, options UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) (result UpdateEnterpriseAppTransitiveRoleAssignmentAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
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
