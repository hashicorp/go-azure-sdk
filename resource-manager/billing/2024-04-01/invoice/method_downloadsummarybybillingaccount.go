package invoice

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

type DownloadSummaryByBillingAccountOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DocumentDownloadResult
}

// DownloadSummaryByBillingAccount ...
func (c InvoiceClient) DownloadSummaryByBillingAccount(ctx context.Context, id BillingAccountInvoiceId) (result DownloadSummaryByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,

		Path: fmt.Sprintf("%s/downloadSummary", id.ID()),
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

// DownloadSummaryByBillingAccountThenPoll performs DownloadSummaryByBillingAccount then polls until it's completed
func (c InvoiceClient) DownloadSummaryByBillingAccountThenPoll(ctx context.Context, id BillingAccountInvoiceId) error {
	result, err := c.DownloadSummaryByBillingAccount(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DownloadSummaryByBillingAccount: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadSummaryByBillingAccount: %+v", err)
	}

	return nil
}
