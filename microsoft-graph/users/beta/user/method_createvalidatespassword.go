package user

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateValidatesPasswordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PasswordValidationInformation
}

type CreateValidatesPasswordOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateValidatesPasswordOperationOptions() CreateValidatesPasswordOperationOptions {
	return CreateValidatesPasswordOperationOptions{}
}

func (o CreateValidatesPasswordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateValidatesPasswordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateValidatesPasswordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateValidatesPassword - Invoke action validatePassword. Check a user's password against the organization's password
// validation policy and report whether the password is valid. Use this action to provide real-time feedback on password
// strength while the user types their password.
func (c UserClient) CreateValidatesPassword(ctx context.Context, input CreateValidatesPasswordRequest, options CreateValidatesPasswordOperationOptions) (result CreateValidatesPasswordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/users/validatePassword",
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

	var model beta.PasswordValidationInformation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
