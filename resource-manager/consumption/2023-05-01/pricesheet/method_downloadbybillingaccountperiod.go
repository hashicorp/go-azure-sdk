package pricesheet

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

type DownloadByBillingAccountPeriodOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *OperationStatus
}

// DownloadByBillingAccountPeriod ...
func (c PriceSheetClient) DownloadByBillingAccountPeriod(ctx context.Context, id BillingAccountBillingPeriodId) (result DownloadByBillingAccountPeriodOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Consumption/pricesheets/download", id.ID()),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DownloadByBillingAccountPeriodThenPoll performs DownloadByBillingAccountPeriod then polls until it's completed
func (c PriceSheetClient) DownloadByBillingAccountPeriodThenPoll(ctx context.Context, id BillingAccountBillingPeriodId) error {
	result, err := c.DownloadByBillingAccountPeriod(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DownloadByBillingAccountPeriod: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DownloadByBillingAccountPeriod: %+v", err)
	}

	return nil
}
