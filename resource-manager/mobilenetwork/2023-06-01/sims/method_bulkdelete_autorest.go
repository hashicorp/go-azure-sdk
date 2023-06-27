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

type BulkDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BulkDelete ...
func (c SIMsClient) BulkDelete(ctx context.Context, id SimGroupId, input SimDeleteList) (result BulkDeleteOperationResponse, err error) {
	req, err := c.preparerForBulkDelete(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "BulkDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBulkDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sims.SIMsClient", "BulkDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BulkDeleteThenPoll performs BulkDelete then polls until it's completed
func (c SIMsClient) BulkDeleteThenPoll(ctx context.Context, id SimGroupId, input SimDeleteList) error {
	result, err := c.BulkDelete(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BulkDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BulkDelete: %+v", err)
	}

	return nil
}

// preparerForBulkDelete prepares the BulkDelete request.
func (c SIMsClient) preparerForBulkDelete(ctx context.Context, id SimGroupId, input SimDeleteList) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deleteSims", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBulkDelete sends the BulkDelete request. The method will close the
// http.Response Body if it receives an error.
func (c SIMsClient) senderForBulkDelete(ctx context.Context, req *http.Request) (future BulkDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
