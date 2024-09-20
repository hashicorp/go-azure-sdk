package monthlyprintusagebyuser

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMonthlyPrintUsageByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateMonthlyPrintUsageByUserOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateMonthlyPrintUsageByUserOperationOptions() UpdateMonthlyPrintUsageByUserOperationOptions {
	return UpdateMonthlyPrintUsageByUserOperationOptions{}
}

func (o UpdateMonthlyPrintUsageByUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateMonthlyPrintUsageByUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateMonthlyPrintUsageByUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateMonthlyPrintUsageByUser - Update the navigation property monthlyPrintUsageByUser in reports
func (c MonthlyPrintUsageByUserClient) UpdateMonthlyPrintUsageByUser(ctx context.Context, id beta.ReportMonthlyPrintUsageByUserId, input beta.PrintUsageByUser, options UpdateMonthlyPrintUsageByUserOperationOptions) (result UpdateMonthlyPrintUsageByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
