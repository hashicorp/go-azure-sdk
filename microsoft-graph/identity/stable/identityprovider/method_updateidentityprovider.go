package identityprovider

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateIdentityProviderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateIdentityProviderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateIdentityProviderOperationOptions() UpdateIdentityProviderOperationOptions {
	return UpdateIdentityProviderOperationOptions{}
}

func (o UpdateIdentityProviderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateIdentityProviderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateIdentityProviderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateIdentityProvider - Update identityProvider. Update the properties of the specified identity provider configured
// in the tenant. Among the types of providers derived from identityProviderBase, you can currently update a
// socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently update a
// socialIdentityProvider, or an appleManagedIdentityProvider resource.
func (c IdentityProviderClient) UpdateIdentityProvider(ctx context.Context, id stable.IdentityIdentityProviderId, input stable.IdentityProviderBase, options UpdateIdentityProviderOperationOptions) (result UpdateIdentityProviderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
