package userinsightmonthlysignup

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserInsightMonthlySignUpsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserInsightMonthlySignUpsCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetUserInsightMonthlySignUpsCountOperationOptions() GetUserInsightMonthlySignUpsCountOperationOptions {
	return GetUserInsightMonthlySignUpsCountOperationOptions{}
}

func (o GetUserInsightMonthlySignUpsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserInsightMonthlySignUpsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetUserInsightMonthlySignUpsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserInsightMonthlySignUpsCount - Get the number of the resource
func (c UserInsightMonthlySignUpClient) GetUserInsightMonthlySignUpsCount(ctx context.Context, options GetUserInsightMonthlySignUpsCountOperationOptions) (result GetUserInsightMonthlySignUpsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/userInsights/monthly/signUps/$count",
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
