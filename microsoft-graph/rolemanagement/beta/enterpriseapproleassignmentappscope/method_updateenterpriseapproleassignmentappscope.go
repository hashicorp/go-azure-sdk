package enterpriseapproleassignmentappscope

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

type UpdateEnterpriseAppRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions() UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions {
	return UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions{}
}

func (o UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEnterpriseAppRoleAssignmentAppScope - Update the navigation property appScope in roleManagement
func (c EnterpriseAppRoleAssignmentAppScopeClient) UpdateEnterpriseAppRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentId, input beta.AppScope, options UpdateEnterpriseAppRoleAssignmentAppScopeOperationOptions) (result UpdateEnterpriseAppRoleAssignmentAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
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
