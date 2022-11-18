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

type PriceSheetDownloadByBillingProfileOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// PriceSheetDownloadByBillingProfile ...
func (c PriceSheetsClient) PriceSheetDownloadByBillingProfile(ctx context.Context, id BillingProfileId) (result PriceSheetDownloadByBillingProfileOperationResponse, err error) {
	req, err := c.preparerForPriceSheetDownloadByBillingProfile(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheets.PriceSheetsClient", "PriceSheetDownloadByBillingProfile", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForPriceSheetDownloadByBillingProfile(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheets.PriceSheetsClient", "PriceSheetDownloadByBillingProfile", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// PriceSheetDownloadByBillingProfileThenPoll performs PriceSheetDownloadByBillingProfile then polls until it's completed
func (c PriceSheetsClient) PriceSheetDownloadByBillingProfileThenPoll(ctx context.Context, id BillingProfileId) error {
	result, err := c.PriceSheetDownloadByBillingProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("performing PriceSheetDownloadByBillingProfile: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after PriceSheetDownloadByBillingProfile: %+v", err)
	}

	return nil
}

// preparerForPriceSheetDownloadByBillingProfile prepares the PriceSheetDownloadByBillingProfile request.
func (c PriceSheetsClient) preparerForPriceSheetDownloadByBillingProfile(ctx context.Context, id BillingProfileId) (*http.Request, error) {
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

// senderForPriceSheetDownloadByBillingProfile sends the PriceSheetDownloadByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (c PriceSheetsClient) senderForPriceSheetDownloadByBillingProfile(ctx context.Context, req *http.Request) (future PriceSheetDownloadByBillingProfileOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
