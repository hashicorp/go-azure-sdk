package entitlementmanagementresourcenamespaceresourceactionauthenticationcontext

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

type GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AuthenticationContextClassReference
}

type GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions() GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions {
	return GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions{}
}

func (o GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContext - Get authenticationContext from
// roleManagement
func (c EntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient) GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContext(ctx context.Context, id beta.RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId, options GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationOptions) (result GetEntitlementManagementResourceNamespaceResourceActionAuthenticationContextOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/authenticationContext", id.ID()),
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

	var model beta.AuthenticationContextClassReference
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
