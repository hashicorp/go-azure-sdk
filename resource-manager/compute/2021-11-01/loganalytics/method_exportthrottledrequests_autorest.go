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

type ExportThrottledRequestsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ExportThrottledRequests ...
func (c LogAnalyticsClient) ExportThrottledRequests(ctx context.Context, id LocationId, input LogAnalyticsInputBase) (result ExportThrottledRequestsOperationResponse, err error) {
	req, err := c.preparerForExportThrottledRequests(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loganalytics.LogAnalyticsClient", "ExportThrottledRequests", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForExportThrottledRequests(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loganalytics.LogAnalyticsClient", "ExportThrottledRequests", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ExportThrottledRequestsThenPoll performs ExportThrottledRequests then polls until it's completed
func (c LogAnalyticsClient) ExportThrottledRequestsThenPoll(ctx context.Context, id LocationId, input LogAnalyticsInputBase) error {
	result, err := c.ExportThrottledRequests(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ExportThrottledRequests: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ExportThrottledRequests: %+v", err)
	}

	return nil
}

// preparerForExportThrottledRequests prepares the ExportThrottledRequests request.
func (c LogAnalyticsClient) preparerForExportThrottledRequests(ctx context.Context, id LocationId, input LogAnalyticsInputBase) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/logAnalytics/apiAccess/getThrottledRequests", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForExportThrottledRequests sends the ExportThrottledRequests request. The method will close the
// http.Response Body if it receives an error.
func (c LogAnalyticsClient) senderForExportThrottledRequests(ctx context.Context, req *http.Request) (future ExportThrottledRequestsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
