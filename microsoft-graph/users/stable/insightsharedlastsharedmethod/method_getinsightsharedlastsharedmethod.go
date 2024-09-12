package insightsharedlastsharedmethod

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightSharedLastSharedMethodOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Entity
}

type GetInsightSharedLastSharedMethodOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetInsightSharedLastSharedMethodOperationOptions() GetInsightSharedLastSharedMethodOperationOptions {
	return GetInsightSharedLastSharedMethodOperationOptions{}
}

func (o GetInsightSharedLastSharedMethodOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetInsightSharedLastSharedMethodOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetInsightSharedLastSharedMethodOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetInsightSharedLastSharedMethod - Get lastSharedMethod from users
func (c InsightSharedLastSharedMethodClient) GetInsightSharedLastSharedMethod(ctx context.Context, id stable.UserIdInsightSharedId, options GetInsightSharedLastSharedMethodOperationOptions) (result GetInsightSharedLastSharedMethodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lastSharedMethod", id.ID()),
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
	model, err := stable.UnmarshalEntityImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
