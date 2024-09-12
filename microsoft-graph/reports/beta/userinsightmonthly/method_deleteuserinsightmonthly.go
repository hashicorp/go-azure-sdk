package userinsightmonthly

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserInsightMonthlyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteUserInsightMonthlyOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteUserInsightMonthlyOperationOptions() DeleteUserInsightMonthlyOperationOptions {
	return DeleteUserInsightMonthlyOperationOptions{}
}

func (o DeleteUserInsightMonthlyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteUserInsightMonthlyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteUserInsightMonthlyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteUserInsightMonthly - Delete navigation property monthly for reports
func (c UserInsightMonthlyClient) DeleteUserInsightMonthly(ctx context.Context, options DeleteUserInsightMonthlyOperationOptions) (result DeleteUserInsightMonthlyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/reports/userInsights/monthly",
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
