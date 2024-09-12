package userinsightdaily

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserInsightDailyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteUserInsightDailyOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteUserInsightDailyOperationOptions() DeleteUserInsightDailyOperationOptions {
	return DeleteUserInsightDailyOperationOptions{}
}

func (o DeleteUserInsightDailyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteUserInsightDailyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteUserInsightDailyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteUserInsightDaily - Delete navigation property daily for reports
func (c UserInsightDailyClient) DeleteUserInsightDaily(ctx context.Context, options DeleteUserInsightDailyOperationOptions) (result DeleteUserInsightDailyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/reports/userInsights/daily",
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
