package dailyprintusagesummariesbyuser

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDailyPrintUsageSummariesByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDailyPrintUsageSummariesByUserOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDailyPrintUsageSummariesByUserOperationOptions() UpdateDailyPrintUsageSummariesByUserOperationOptions {
	return UpdateDailyPrintUsageSummariesByUserOperationOptions{}
}

func (o UpdateDailyPrintUsageSummariesByUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDailyPrintUsageSummariesByUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDailyPrintUsageSummariesByUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDailyPrintUsageSummariesByUser - Update the navigation property dailyPrintUsageSummariesByUser in reports
func (c DailyPrintUsageSummariesByUserClient) UpdateDailyPrintUsageSummariesByUser(ctx context.Context, id beta.ReportDailyPrintUsageSummariesByUserId, input beta.PrintUsageByUser, options UpdateDailyPrintUsageSummariesByUserOperationOptions) (result UpdateDailyPrintUsageSummariesByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
