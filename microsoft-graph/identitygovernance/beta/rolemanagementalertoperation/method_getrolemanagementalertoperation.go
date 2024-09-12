package rolemanagementalertoperation

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

type GetRoleManagementAlertOperationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.LongRunningOperation
}

type GetRoleManagementAlertOperationOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetRoleManagementAlertOperationOperationOptions() GetRoleManagementAlertOperationOperationOptions {
	return GetRoleManagementAlertOperationOperationOptions{}
}

func (o GetRoleManagementAlertOperationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRoleManagementAlertOperationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetRoleManagementAlertOperationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRoleManagementAlertOperation - Get operations from identityGovernance. Represents operations on resources that
// take a long time to complete and can run in the background until completion.
func (c RoleManagementAlertOperationClient) GetRoleManagementAlertOperation(ctx context.Context, id beta.IdentityGovernanceRoleManagementAlertOperationId, options GetRoleManagementAlertOperationOperationOptions) (result GetRoleManagementAlertOperationOperationResponse, err error) {
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
	model, err := beta.UnmarshalLongRunningOperationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
