package pricesheets

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

type PriceSheetDownloadOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DownloadUrl
}

// PriceSheetDownload ...
func (c PriceSheetsClient) PriceSheetDownload(ctx context.Context, id InvoiceId) (result PriceSheetDownloadOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.CostManagement/pricesheets/default/download", id.ID()),
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

// PriceSheetDownloadThenPoll performs PriceSheetDownload then polls until it's completed
func (c PriceSheetsClient) PriceSheetDownloadThenPoll(ctx context.Context, id InvoiceId) error {
	result, err := c.PriceSheetDownload(ctx, id)
	if err != nil {
		return fmt.Errorf("performing PriceSheetDownload: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after PriceSheetDownload: %+v", err)
	}

	return nil
}
