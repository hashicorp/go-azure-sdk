package rolemanagementalertalertconfigurationalertdefinition

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

type GetRoleManagementAlertConfigurationAlertDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleManagementAlertDefinition
}

type GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetRoleManagementAlertConfigurationAlertDefinitionOperationOptions() GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions {
	return GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions{}
}

func (o GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions) ToOData() *odata.Query {
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

func (o GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRoleManagementAlertConfigurationAlertDefinition - Get alertDefinition from identityGovernance. The definition of
// the alert that contains its description, impact, and measures to mitigate or prevent it. Supports $expand.
func (c RoleManagementAlertAlertConfigurationAlertDefinitionClient) GetRoleManagementAlertConfigurationAlertDefinition(ctx context.Context, id beta.IdentityGovernanceRoleManagementAlertAlertConfigurationId, options GetRoleManagementAlertConfigurationAlertDefinitionOperationOptions) (result GetRoleManagementAlertConfigurationAlertDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/alertDefinition", id.ID()),
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

	var model beta.UnifiedRoleManagementAlertDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
