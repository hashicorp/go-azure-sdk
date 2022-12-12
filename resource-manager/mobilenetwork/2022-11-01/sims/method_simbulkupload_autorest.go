package sims

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

type SimBulkUploadOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// SimBulkUpload ...
func (c SIMsClient) SimBulkUpload(ctx context.Context, id SimGroupId, input SimUploadList) (result SimBulkUploadOperationResponse, err error) {
	req, err := c.preparerForSimBulkUpload(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "SimBulkUpload", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSimBulkUpload(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "SimBulkUpload", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SimBulkUploadThenPoll performs SimBulkUpload then polls until it's completed
func (c SIMsClient) SimBulkUploadThenPoll(ctx context.Context, id SimGroupId, input SimUploadList) error {
	result, err := c.SimBulkUpload(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SimBulkUpload: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SimBulkUpload: %+v", err)
	}

	return nil
}

// preparerForSimBulkUpload prepares the SimBulkUpload request.
func (c SIMsClient) preparerForSimBulkUpload(ctx context.Context, id SimGroupId, input SimUploadList) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/uploadSims", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSimBulkUpload sends the SimBulkUpload request. The method will close the
// http.Response Body if it receives an error.
func (c SIMsClient) senderForSimBulkUpload(ctx context.Context, req *http.Request) (future SimBulkUploadOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
