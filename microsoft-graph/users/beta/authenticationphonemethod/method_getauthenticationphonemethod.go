package authenticationphonemethod

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationPhoneMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PhoneAuthenticationMethod
}

type GetAuthenticationPhoneMethodOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationPhoneMethodOperationOptions() GetAuthenticationPhoneMethodOperationOptions {
	return GetAuthenticationPhoneMethodOperationOptions{}
}

func (o GetAuthenticationPhoneMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationPhoneMethodOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationPhoneMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationPhoneMethod - Get phoneMethods from users. Represents the phone registered to a user for
// authentication.
func (c AuthenticationPhoneMethodClient) GetAuthenticationPhoneMethod(ctx context.Context, id beta.UserIdAuthenticationPhoneMethodId, options GetAuthenticationPhoneMethodOperationOptions) (result GetAuthenticationPhoneMethodOperationResponse, err error) {
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

	var model beta.PhoneAuthenticationMethod
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
