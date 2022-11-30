package restorepoints

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

type RestorePointsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestorePointsDelete ...
func (c RestorePointsClient) RestorePointsDelete(ctx context.Context, id RestorePointId) (result RestorePointsDeleteOperationResponse, err error) {
	req, err := c.preparerForRestorePointsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorepoints.RestorePointsClient", "RestorePointsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestorePointsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorepoints.RestorePointsClient", "RestorePointsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestorePointsDeleteThenPoll performs RestorePointsDelete then polls until it's completed
func (c RestorePointsClient) RestorePointsDeleteThenPoll(ctx context.Context, id RestorePointId) error {
	result, err := c.RestorePointsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RestorePointsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestorePointsDelete: %+v", err)
	}

	return nil
}

// preparerForRestorePointsDelete prepares the RestorePointsDelete request.
func (c RestorePointsClient) preparerForRestorePointsDelete(ctx context.Context, id RestorePointId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestorePointsDelete sends the RestorePointsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c RestorePointsClient) senderForRestorePointsDelete(ctx context.Context, req *http.Request) (future RestorePointsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
