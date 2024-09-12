package devicemanagementroleassignmentappscope

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.AppScope
}

type GetDeviceManagementRoleAssignmentAppScopeOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceManagementRoleAssignmentAppScopeOperationOptions() GetDeviceManagementRoleAssignmentAppScopeOperationOptions {
	return GetDeviceManagementRoleAssignmentAppScopeOperationOptions{}
}

func (o GetDeviceManagementRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceManagementRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceManagementRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceManagementRoleAssignmentAppScope - Get appScopes from roleManagement. Read-only collection with details of
// the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
func (c DeviceManagementRoleAssignmentAppScopeClient) GetDeviceManagementRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentIdAppScopeId, options GetDeviceManagementRoleAssignmentAppScopeOperationOptions) (result GetDeviceManagementRoleAssignmentAppScopeOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalAppScopeImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
