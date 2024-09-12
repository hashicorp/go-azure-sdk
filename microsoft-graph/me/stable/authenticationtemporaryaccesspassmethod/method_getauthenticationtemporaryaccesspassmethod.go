package authenticationtemporaryaccesspassmethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationTemporaryAccessPassMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TemporaryAccessPassAuthenticationMethod
}

type GetAuthenticationTemporaryAccessPassMethodOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAuthenticationTemporaryAccessPassMethodOperationOptions() GetAuthenticationTemporaryAccessPassMethodOperationOptions {
	return GetAuthenticationTemporaryAccessPassMethodOperationOptions{}
}

func (o GetAuthenticationTemporaryAccessPassMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationTemporaryAccessPassMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAuthenticationTemporaryAccessPassMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationTemporaryAccessPassMethod - Get temporaryAccessPassMethods from me. Represents a Temporary Access
// Pass registered to a user for authentication through time-limited passcodes.
func (c AuthenticationTemporaryAccessPassMethodClient) GetAuthenticationTemporaryAccessPassMethod(ctx context.Context, id stable.MeAuthenticationTemporaryAccessPassMethodId, options GetAuthenticationTemporaryAccessPassMethodOperationOptions) (result GetAuthenticationTemporaryAccessPassMethodOperationResponse, err error) {
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

	var model stable.TemporaryAccessPassAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
