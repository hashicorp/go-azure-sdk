package userinsightmonthlysignup

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserInsightMonthlySignUpOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserSignUpMetric
}

type GetUserInsightMonthlySignUpOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetUserInsightMonthlySignUpOperationOptions() GetUserInsightMonthlySignUpOperationOptions {
	return GetUserInsightMonthlySignUpOperationOptions{}
}

func (o GetUserInsightMonthlySignUpOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserInsightMonthlySignUpOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserInsightMonthlySignUpOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserInsightMonthlySignUp - Get signUps from reports. Total sign-ups on apps registered in the tenant for a
// specified period.
func (c UserInsightMonthlySignUpClient) GetUserInsightMonthlySignUp(ctx context.Context, id beta.ReportUserInsightMonthlySignUpId, options GetUserInsightMonthlySignUpOperationOptions) (result GetUserInsightMonthlySignUpOperationResponse, err error) {
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

	var model beta.UserSignUpMetric
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
