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

type DeleteEnterpriseAppRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions() DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions {
	return DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions{}
}

func (o DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEnterpriseAppRoleAssignmentAppScope - Delete navigation property appScope for roleManagement
func (c EnterpriseAppRoleAssignmentAppScopeClient) DeleteEnterpriseAppRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentId, options DeleteEnterpriseAppRoleAssignmentAppScopeOperationOptions) (result DeleteEnterpriseAppRoleAssignmentAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
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
