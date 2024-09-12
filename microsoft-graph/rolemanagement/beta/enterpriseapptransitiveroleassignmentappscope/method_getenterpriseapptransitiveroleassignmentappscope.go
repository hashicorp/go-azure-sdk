package enterpriseapptransitiveroleassignmentappscope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.AppScope
}

type GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions() GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions {
	return GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions{}
}

func (o GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEnterpriseAppTransitiveRoleAssignmentAppScope - Get appScope from roleManagement. Read-only property with details
// of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand for the
// entitlement provider only.
func (c EnterpriseAppTransitiveRoleAssignmentAppScopeClient) GetEnterpriseAppTransitiveRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementEnterpriseAppIdTransitiveRoleAssignmentId, options GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationOptions) (result GetEnterpriseAppTransitiveRoleAssignmentAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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
