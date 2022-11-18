package pricesheets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetDownloadOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PriceSheetDownload ...
func (c PriceSheetsClient) PriceSheetDownload(ctx context.Context, id InvoiceId) (result PriceSheetDownloadOperationResponse, err error) {
	req, err := c.preparerForPriceSheetDownload(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheets.PriceSheetsClient", "PriceSheetDownload", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPriceSheetDownload(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheets.PriceSheetsClient", "PriceSheetDownload", result.HttpResponse, "Failure sending request")
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

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PriceSheetDownload: %+v", err)
	}

	return nil
}

// preparerForPriceSheetDownload prepares the PriceSheetDownload request.
func (c PriceSheetsClient) preparerForPriceSheetDownload(ctx context.Context, id InvoiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CostManagement/pricesheets/default/download", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForPriceSheetDownload sends the PriceSheetDownload request. The method will close the
// http.Response Body if it receives an error.
func (c PriceSheetsClient) senderForPriceSheetDownload(ctx context.Context, req *http.Request) (future PriceSheetDownloadOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
