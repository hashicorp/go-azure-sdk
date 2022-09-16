package integrationruntime

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

type DisableInteractiveQueryOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DisableInteractiveQuery ...
func (c IntegrationRuntimeClient) DisableInteractiveQuery(ctx context.Context, id IntegrationRuntimeId) (result DisableInteractiveQueryOperationResponse, err error) {
	req, err := c.preparerForDisableInteractiveQuery(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "DisableInteractiveQuery", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDisableInteractiveQuery(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "DisableInteractiveQuery", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DisableInteractiveQueryThenPoll performs DisableInteractiveQuery then polls until it's completed
func (c IntegrationRuntimeClient) DisableInteractiveQueryThenPoll(ctx context.Context, id IntegrationRuntimeId) error {
	result, err := c.DisableInteractiveQuery(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DisableInteractiveQuery: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DisableInteractiveQuery: %+v", err)
	}

	return nil
}

// preparerForDisableInteractiveQuery prepares the DisableInteractiveQuery request.
func (c IntegrationRuntimeClient) preparerForDisableInteractiveQuery(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/disableInteractiveQuery", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDisableInteractiveQuery sends the DisableInteractiveQuery request. The method will close the
// http.Response Body if it receives an error.
func (c IntegrationRuntimeClient) senderForDisableInteractiveQuery(ctx context.Context, req *http.Request) (future DisableInteractiveQueryOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
