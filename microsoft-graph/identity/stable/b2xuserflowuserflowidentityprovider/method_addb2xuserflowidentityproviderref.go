package b2xuserflowuserflowidentityprovider

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

type AddB2xUserFlowIdentityProviderRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddB2xUserFlowIdentityProviderRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddB2xUserFlowIdentityProviderRefOperationOptions() AddB2xUserFlowIdentityProviderRefOperationOptions {
	return AddB2xUserFlowIdentityProviderRefOperationOptions{}
}

func (o AddB2xUserFlowIdentityProviderRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddB2xUserFlowIdentityProviderRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddB2xUserFlowIdentityProviderRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddB2xUserFlowIdentityProviderRef - Create new navigation property ref to userFlowIdentityProviders for identity
func (c B2xUserFlowUserFlowIdentityProviderClient) AddB2xUserFlowIdentityProviderRef(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.ReferenceCreate, options AddB2xUserFlowIdentityProviderRefOperationOptions) (result AddB2xUserFlowIdentityProviderRefOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/userFlowIdentityProviders/$ref", id.ID()),
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
