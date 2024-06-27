package transaction

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTransactionSummaryByInvoiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *TransactionSummary
}

type GetTransactionSummaryByInvoiceOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetTransactionSummaryByInvoiceOperationOptions() GetTransactionSummaryByInvoiceOperationOptions {
	return GetTransactionSummaryByInvoiceOperationOptions{}
}

func (o GetTransactionSummaryByInvoiceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTransactionSummaryByInvoiceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetTransactionSummaryByInvoiceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Search != nil {
		out.Append("search", fmt.Sprintf("%v", *o.Search))
	}
	return &out
}

// GetTransactionSummaryByInvoice ...
func (c TransactionClient) GetTransactionSummaryByInvoice(ctx context.Context, id BillingAccountInvoiceId, options GetTransactionSummaryByInvoiceOperationOptions) (result GetTransactionSummaryByInvoiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/transactionSummary", id.ID()),
		OptionsObject: options,
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

	var model TransactionSummary
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
