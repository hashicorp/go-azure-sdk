package identityprovider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteIdentityProviderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteIdentityProviderOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteIdentityProviderOperationOptions() DeleteIdentityProviderOperationOptions {
	return DeleteIdentityProviderOperationOptions{}
}

func (o DeleteIdentityProviderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteIdentityProviderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteIdentityProviderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteIdentityProvider - Delete identityProvider. Delete an identity provider resource that is of the type specified
// by the id in the request. Among the types of providers derived from identityProviderBase, you can currently delete a
// socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently delete a
// socialIdentityProvider, or an appleManagedIdentityProvider resource.
func (c IdentityProviderClient) DeleteIdentityProvider(ctx context.Context, id stable.IdentityIdentityProviderId, options DeleteIdentityProviderOperationOptions) (result DeleteIdentityProviderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
