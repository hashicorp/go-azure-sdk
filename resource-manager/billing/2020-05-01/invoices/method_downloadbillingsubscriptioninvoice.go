package invoices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DownloadBillingSubscriptionInvoiceOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DownloadUrl
}

type DownloadBillingSubscriptionInvoiceOperationOptions struct {
	DownloadToken *string
}

func DefaultDownloadBillingSubscriptionInvoiceOperationOptions() DownloadBillingSubscriptionInvoiceOperationOptions {
	return DownloadBillingSubscriptionInvoiceOperationOptions{}
}

func (o DownloadBillingSubscriptionInvoiceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DownloadBillingSubscriptionInvoiceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DownloadBillingSubscriptionInvoiceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DownloadToken != nil {
		out.Append("downloadToken", fmt.Sprintf("%v", *o.DownloadToken))
	}
	return &out
}

// DownloadBillingSubscriptionInvoice ...
func (c InvoicesClient) DownloadBillingSubscriptionInvoice(ctx context.Context, id BillingSubscriptionInvoiceId, options DownloadBillingSubscriptionInvoiceOperationOptions) (result DownloadBillingSubscriptionInvoiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/download", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DownloadBillingSubscriptionInvoiceThenPoll performs DownloadBillingSubscriptionInvoice then polls until it's completed
func (c InvoicesClient) DownloadBillingSubscriptionInvoiceThenPoll(ctx context.Context, id BillingSubscriptionInvoiceId, options DownloadBillingSubscriptionInvoiceOperationOptions) error {
	result, err := c.DownloadBillingSubscriptionInvoice(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DownloadBillingSubscriptionInvoice: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadBillingSubscriptionInvoice: %+v", err)
	}

	return nil
}
