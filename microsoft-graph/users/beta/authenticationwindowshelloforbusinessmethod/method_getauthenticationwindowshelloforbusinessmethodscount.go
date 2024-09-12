package authenticationwindowshelloforbusinessmethod

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

type GetAuthenticationWindowsHelloForBusinessMethodsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions() GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions {
	return GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions{}
}

func (o GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationWindowsHelloForBusinessMethodsCount - Get the number of the resource
func (c AuthenticationWindowsHelloForBusinessMethodClient) GetAuthenticationWindowsHelloForBusinessMethodsCount(ctx context.Context, id beta.UserId, options GetAuthenticationWindowsHelloForBusinessMethodsCountOperationOptions) (result GetAuthenticationWindowsHelloForBusinessMethodsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/authentication/windowsHelloForBusinessMethods/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
