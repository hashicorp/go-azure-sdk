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

type BulkUploadOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BulkUpload ...
func (c SIMsClient) BulkUpload(ctx context.Context, id SimGroupId, input SimUploadList) (result BulkUploadOperationResponse, err error) {
	req, err := c.preparerForBulkUpload(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "BulkUpload", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBulkUpload(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "BulkUpload", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BulkUploadThenPoll performs BulkUpload then polls until it's completed
func (c SIMsClient) BulkUploadThenPoll(ctx context.Context, id SimGroupId, input SimUploadList) error {
	result, err := c.BulkUpload(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BulkUpload: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BulkUpload: %+v", err)
	}

	return nil
}

// preparerForBulkUpload prepares the BulkUpload request.
func (c SIMsClient) preparerForBulkUpload(ctx context.Context, id SimGroupId, input SimUploadList) (*http.Request, error) {
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

// senderForBulkUpload sends the BulkUpload request. The method will close the
// http.Response Body if it receives an error.
func (c SIMsClient) senderForBulkUpload(ctx context.Context, req *http.Request) (future BulkUploadOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
