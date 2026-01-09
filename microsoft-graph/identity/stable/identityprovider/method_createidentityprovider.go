package identityprovider

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateIdentityProviderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.IdentityProviderBase
}

type CreateIdentityProviderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateIdentityProviderOperationOptions() CreateIdentityProviderOperationOptions {
	return CreateIdentityProviderOperationOptions{}
}

func (o CreateIdentityProviderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateIdentityProviderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateIdentityProviderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateIdentityProvider - Create identityProvider. Create an identity provider object that is of the type specified in
// the request body. Among the types of providers derived from identityProviderBase, you can currently create a
// socialIdentityProvider resource in Microsoft Entra ID. In Azure AD B2C, this operation can currently create a
// socialIdentityProvider, or an appleManagedIdentityProvider resource.
func (c IdentityProviderClient) CreateIdentityProvider(ctx context.Context, input stable.IdentityProviderBase, options CreateIdentityProviderOperationOptions) (result CreateIdentityProviderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identity/identityProviders",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalIdentityProviderBaseImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
