package b2xuserflowidentityprovider

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2xUserFlowIdentityProviderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityProvider
}

type GetB2xUserFlowIdentityProviderOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetB2xUserFlowIdentityProviderOperationOptions() GetB2xUserFlowIdentityProviderOperationOptions {
	return GetB2xUserFlowIdentityProviderOperationOptions{}
}

func (o GetB2xUserFlowIdentityProviderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowIdentityProviderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetB2xUserFlowIdentityProviderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowIdentityProvider - Get identityProviders from identity. The identity providers included in the user
// flow.
func (c B2xUserFlowIdentityProviderClient) GetB2xUserFlowIdentityProvider(ctx context.Context, id stable.IdentityB2xUserFlowIdIdentityProviderId, options GetB2xUserFlowIdentityProviderOperationOptions) (result GetB2xUserFlowIdentityProviderOperationResponse, err error) {
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

	var model stable.IdentityProvider
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
