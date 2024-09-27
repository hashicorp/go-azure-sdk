package authenticationfido2method

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

type DeleteAuthenticationFido2MethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAuthenticationFido2MethodOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteAuthenticationFido2MethodOperationOptions() DeleteAuthenticationFido2MethodOperationOptions {
	return DeleteAuthenticationFido2MethodOperationOptions{}
}

func (o DeleteAuthenticationFido2MethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAuthenticationFido2MethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAuthenticationFido2MethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAuthenticationFido2Method - Delete navigation property fido2Methods for me
func (c AuthenticationFido2MethodClient) DeleteAuthenticationFido2Method(ctx context.Context, id stable.MeAuthenticationFido2MethodId, options DeleteAuthenticationFido2MethodOperationOptions) (result DeleteAuthenticationFido2MethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
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
