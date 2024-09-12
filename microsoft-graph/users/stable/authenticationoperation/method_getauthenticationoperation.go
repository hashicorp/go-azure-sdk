package authenticationoperation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAuthenticationOperationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.LongRunningOperation
}

type GetAuthenticationOperationOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAuthenticationOperationOperationOptions() GetAuthenticationOperationOperationOptions {
	return GetAuthenticationOperationOperationOptions{}
}

func (o GetAuthenticationOperationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationOperationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAuthenticationOperationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationOperation - Get longRunningOperation. Read the properties and relationships of a
// longRunningOperation object. This API allows you to retrieve the details and status of the following long-running
// Microsoft Graph API operations. The possible states of the long-running operation are notStarted, running, succeeded,
// failed, unknownFutureValue where succeeded and failed are terminal states.
func (c AuthenticationOperationClient) GetAuthenticationOperation(ctx context.Context, id stable.UserIdAuthenticationOperationId, options GetAuthenticationOperationOperationOptions) (result GetAuthenticationOperationOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalLongRunningOperationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
