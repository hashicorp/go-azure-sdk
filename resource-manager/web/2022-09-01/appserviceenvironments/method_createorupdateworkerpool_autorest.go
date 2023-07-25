package appserviceenvironments

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

type CreateOrUpdateWorkerPoolOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateWorkerPool ...
func (c AppServiceEnvironmentsClient) CreateOrUpdateWorkerPool(ctx context.Context, id WorkerPoolId, input WorkerPoolResource) (result CreateOrUpdateWorkerPoolOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateWorkerPool(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "CreateOrUpdateWorkerPool", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateWorkerPool(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "CreateOrUpdateWorkerPool", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateWorkerPoolThenPoll performs CreateOrUpdateWorkerPool then polls until it's completed
func (c AppServiceEnvironmentsClient) CreateOrUpdateWorkerPoolThenPoll(ctx context.Context, id WorkerPoolId, input WorkerPoolResource) error {
	result, err := c.CreateOrUpdateWorkerPool(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateWorkerPool: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateWorkerPool: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateWorkerPool prepares the CreateOrUpdateWorkerPool request.
func (c AppServiceEnvironmentsClient) preparerForCreateOrUpdateWorkerPool(ctx context.Context, id WorkerPoolId, input WorkerPoolResource) (*http.Request, error) {
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

// senderForCreateOrUpdateWorkerPool sends the CreateOrUpdateWorkerPool request. The method will close the
// http.Response Body if it receives an error.
func (c AppServiceEnvironmentsClient) senderForCreateOrUpdateWorkerPool(ctx context.Context, req *http.Request) (future CreateOrUpdateWorkerPoolOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
