package exchangeroledefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetExchangeRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleDefinition
}

type GetExchangeRoleDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetExchangeRoleDefinitionOperationOptions() GetExchangeRoleDefinitionOperationOptions {
	return GetExchangeRoleDefinitionOperationOptions{}
}

func (o GetExchangeRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetExchangeRoleDefinitionOperationOptions) ToOData() *odata.Query {
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

func (o GetExchangeRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetExchangeRoleDefinition - Get unifiedRoleDefinition. Get the properties and relationships of a
// unifiedRoleDefinition object of an RBAC provider. The following RBAC providers are currently supported: - Cloud PC -
// device management (Intune) - directory (Microsoft Entra directory roles) - entitlement management (Microsoft Entra
// entitlement management) - Exchange Online
func (c ExchangeRoleDefinitionClient) GetExchangeRoleDefinition(ctx context.Context, id beta.RoleManagementExchangeRoleDefinitionId, options GetExchangeRoleDefinitionOperationOptions) (result GetExchangeRoleDefinitionOperationResponse, err error) {
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
