package entitlementmanagementtransitiveroleassignmentappscope

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

type GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.AppScope
}

type GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions() GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions {
	return GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions{}
}

func (o GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementTransitiveRoleAssignmentAppScope - Get appScope from roleManagement. Read-only property with
// details of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand for
// the entitlement provider only.
func (c EntitlementManagementTransitiveRoleAssignmentAppScopeClient) GetEntitlementManagementTransitiveRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementEntitlementManagementTransitiveRoleAssignmentId, options GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationOptions) (result GetEntitlementManagementTransitiveRoleAssignmentAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
		RetryFunc:     options.RetryFunc,
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
