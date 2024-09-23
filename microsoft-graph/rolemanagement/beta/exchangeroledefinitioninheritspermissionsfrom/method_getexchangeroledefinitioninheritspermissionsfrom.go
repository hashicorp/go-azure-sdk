package exchangeroledefinitioninheritspermissionsfrom

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetExchangeRoleDefinitionInheritsPermissionsFromOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleDefinition
}

type GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions() GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions {
	return GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions{}
}

func (o GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions) ToOData() *odata.Query {
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

func (o GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetExchangeRoleDefinitionInheritsPermissionsFrom - Get inheritsPermissionsFrom from roleManagement. Read-only
// collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles
// support this attribute.
func (c ExchangeRoleDefinitionInheritsPermissionsFromClient) GetExchangeRoleDefinitionInheritsPermissionsFrom(ctx context.Context, id beta.RoleManagementExchangeRoleDefinitionIdInheritsPermissionsFromId, options GetExchangeRoleDefinitionInheritsPermissionsFromOperationOptions) (result GetExchangeRoleDefinitionInheritsPermissionsFromOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
