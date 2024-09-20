package userinsightdaily

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserInsightDailyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserInsightDailyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateUserInsightDailyOperationOptions() UpdateUserInsightDailyOperationOptions {
	return UpdateUserInsightDailyOperationOptions{}
}

func (o UpdateUserInsightDailyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserInsightDailyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserInsightDailyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserInsightDaily - Update the navigation property daily in reports
func (c UserInsightDailyClient) UpdateUserInsightDaily(ctx context.Context, input beta.DailyUserInsightMetricsRoot, options UpdateUserInsightDailyOperationOptions) (result UpdateUserInsightDailyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/reports/userInsights/daily",
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

	return
}
