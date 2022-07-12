package compute

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

type RestorePointsCreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RestorePointsCreate ...
func (c ComputeClient) RestorePointsCreate(ctx context.Context, id RestorePointId, input RestorePoint) (result RestorePointsCreateOperationResponse, err error) {
	req, err := c.preparerForRestorePointsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ComputeClient", "RestorePointsCreate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRestorePointsCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ComputeClient", "RestorePointsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RestorePointsCreateThenPoll performs RestorePointsCreate then polls until it's completed
func (c ComputeClient) RestorePointsCreateThenPoll(ctx context.Context, id RestorePointId, input RestorePoint) error {
	result, err := c.RestorePointsCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestorePointsCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RestorePointsCreate: %+v", err)
	}

	return nil
}

// preparerForRestorePointsCreate prepares the RestorePointsCreate request.
func (c ComputeClient) preparerForRestorePointsCreate(ctx context.Context, id RestorePointId, input RestorePoint) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRestorePointsCreate sends the RestorePointsCreate request. The method will close the
// http.Response Body if it receives an error.
func (c ComputeClient) senderForRestorePointsCreate(ctx context.Context, req *http.Request) (future RestorePointsCreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
