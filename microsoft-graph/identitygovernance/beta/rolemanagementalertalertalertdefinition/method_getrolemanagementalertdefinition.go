package rolemanagementalertalertalertdefinition

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

type GetRoleManagementAlertDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleManagementAlertDefinition
}

type GetRoleManagementAlertDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetRoleManagementAlertDefinitionOperationOptions() GetRoleManagementAlertDefinitionOperationOptions {
	return GetRoleManagementAlertDefinitionOperationOptions{}
}

func (o GetRoleManagementAlertDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRoleManagementAlertDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetRoleManagementAlertDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRoleManagementAlertDefinition - Get alertDefinition from identityGovernance. Contains the description, impact, and
// measures to mitigate or prevent the security alert from being triggered in your tenant. Supports $expand.
func (c RoleManagementAlertAlertAlertDefinitionClient) GetRoleManagementAlertDefinition(ctx context.Context, id beta.IdentityGovernanceRoleManagementAlertAlertId, options GetRoleManagementAlertDefinitionOperationOptions) (result GetRoleManagementAlertDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/alertDefinition", id.ID()),
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

	var model beta.UnifiedRoleManagementAlertDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
