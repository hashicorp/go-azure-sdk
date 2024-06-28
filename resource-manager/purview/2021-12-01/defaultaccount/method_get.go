package defaultaccount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DefaultAccountPayload
}

type GetOperationOptions struct {
	Scope         *string
	ScopeTenantId *string
	ScopeType     *ScopeType
}

func DefaultGetOperationOptions() GetOperationOptions {
	return GetOperationOptions{}
}

func (o GetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Scope != nil {
		out.Append("scope", fmt.Sprintf("%v", *o.Scope))
	}
	if o.ScopeTenantId != nil {
		out.Append("scopeTenantId", fmt.Sprintf("%v", *o.ScopeTenantId))
	}
	if o.ScopeType != nil {
		out.Append("scopeType", fmt.Sprintf("%v", *o.ScopeType))
	}
	return &out
}

// Get ...
func (c DefaultAccountClient) Get(ctx context.Context, options GetOperationOptions) (result GetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,

		Path:          "/providers/Microsoft.Purview/getDefaultAccount",
		OptionsObject: options,
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

	var model DefaultAccountPayload
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
