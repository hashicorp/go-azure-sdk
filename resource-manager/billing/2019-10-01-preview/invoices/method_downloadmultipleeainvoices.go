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

type DownloadMultipleEAInvoicesOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// DownloadMultipleEAInvoices ...
func (c InvoicesClient) DownloadMultipleEAInvoices(ctx context.Context, id BillingAccountId, input []string) (result DownloadMultipleEAInvoicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/downloadDocuments", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DownloadMultipleEAInvoicesThenPoll performs DownloadMultipleEAInvoices then polls until it's completed
func (c InvoicesClient) DownloadMultipleEAInvoicesThenPoll(ctx context.Context, id BillingAccountId, input []string) error {
	result, err := c.DownloadMultipleEAInvoices(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing DownloadMultipleEAInvoices: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadMultipleEAInvoices: %+v", err)
	}

	return nil
}
