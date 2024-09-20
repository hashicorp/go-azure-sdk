package authenticationoperation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationOperationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.LongRunningOperation
}

type CreateAuthenticationOperationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAuthenticationOperationOperationOptions() CreateAuthenticationOperationOperationOptions {
	return CreateAuthenticationOperationOperationOptions{}
}

func (o CreateAuthenticationOperationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationOperationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationOperationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationOperation - Create new navigation property to operations for me
func (c AuthenticationOperationClient) CreateAuthenticationOperation(ctx context.Context, input beta.LongRunningOperation, options CreateAuthenticationOperationOperationOptions) (result CreateAuthenticationOperationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/authentication/operations",
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
	model, err := beta.UnmarshalLongRunningOperationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
