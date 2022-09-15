package loganalytics

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

type ExportRequestRateByIntervalOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ExportRequestRateByInterval ...
func (c LogAnalyticsClient) ExportRequestRateByInterval(ctx context.Context, id LocationId, input RequestRateByIntervalInput) (result ExportRequestRateByIntervalOperationResponse, err error) {
	req, err := c.preparerForExportRequestRateByInterval(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loganalytics.LogAnalyticsClient", "ExportRequestRateByInterval", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForExportRequestRateByInterval(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loganalytics.LogAnalyticsClient", "ExportRequestRateByInterval", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ExportRequestRateByIntervalThenPoll performs ExportRequestRateByInterval then polls until it's completed
func (c LogAnalyticsClient) ExportRequestRateByIntervalThenPoll(ctx context.Context, id LocationId, input RequestRateByIntervalInput) error {
	result, err := c.ExportRequestRateByInterval(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ExportRequestRateByInterval: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ExportRequestRateByInterval: %+v", err)
	}

	return nil
}

// preparerForExportRequestRateByInterval prepares the ExportRequestRateByInterval request.
func (c LogAnalyticsClient) preparerForExportRequestRateByInterval(ctx context.Context, id LocationId, input RequestRateByIntervalInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/logAnalytics/apiAccess/getRequestRateByInterval", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForExportRequestRateByInterval sends the ExportRequestRateByInterval request. The method will close the
// http.Response Body if it receives an error.
func (c LogAnalyticsClient) senderForExportRequestRateByInterval(ctx context.Context, req *http.Request) (future ExportRequestRateByIntervalOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
