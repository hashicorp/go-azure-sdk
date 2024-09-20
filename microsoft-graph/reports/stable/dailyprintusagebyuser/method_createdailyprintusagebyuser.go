package dailyprintusagebyuser

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDailyPrintUsageByUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrintUsageByUser
}

type CreateDailyPrintUsageByUserOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDailyPrintUsageByUserOperationOptions() CreateDailyPrintUsageByUserOperationOptions {
	return CreateDailyPrintUsageByUserOperationOptions{}
}

func (o CreateDailyPrintUsageByUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDailyPrintUsageByUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDailyPrintUsageByUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDailyPrintUsageByUser - Create new navigation property to dailyPrintUsageByUser for reports
func (c DailyPrintUsageByUserClient) CreateDailyPrintUsageByUser(ctx context.Context, input stable.PrintUsageByUser, options CreateDailyPrintUsageByUserOperationOptions) (result CreateDailyPrintUsageByUserOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/dailyPrintUsageByUser",
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

	var model stable.PrintUsageByUser
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
