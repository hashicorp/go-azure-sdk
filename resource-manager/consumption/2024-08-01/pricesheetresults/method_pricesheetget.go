package pricesheetresults

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PriceSheetResult
}

type PriceSheetGetOperationOptions struct {
	Expand *string
	Top    *int64
}

func DefaultPriceSheetGetOperationOptions() PriceSheetGetOperationOptions {
	return PriceSheetGetOperationOptions{}
}

func (o PriceSheetGetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PriceSheetGetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PriceSheetGetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// PriceSheetGet ...
func (c PriceSheetResultsClient) PriceSheetGet(ctx context.Context, id commonids.SubscriptionId, options PriceSheetGetOperationOptions) (result PriceSheetGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/pricesheets/default", id.ID()),
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

	var model PriceSheetResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
