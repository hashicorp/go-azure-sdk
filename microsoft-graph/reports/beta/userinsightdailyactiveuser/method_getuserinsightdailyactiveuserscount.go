package userinsightdailyactiveuser

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserInsightDailyActiveUsersCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserInsightDailyActiveUsersCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetUserInsightDailyActiveUsersCountOperationOptions() GetUserInsightDailyActiveUsersCountOperationOptions {
	return GetUserInsightDailyActiveUsersCountOperationOptions{}
}

func (o GetUserInsightDailyActiveUsersCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserInsightDailyActiveUsersCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetUserInsightDailyActiveUsersCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserInsightDailyActiveUsersCount - Get the number of the resource
func (c UserInsightDailyActiveUserClient) GetUserInsightDailyActiveUsersCount(ctx context.Context, options GetUserInsightDailyActiveUsersCountOperationOptions) (result GetUserInsightDailyActiveUsersCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/userInsights/daily/activeUsers/$count",
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
