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

type EnableInteractiveQueryOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// EnableInteractiveQuery ...
func (c IntegrationRuntimeClient) EnableInteractiveQuery(ctx context.Context, id IntegrationRuntimeId) (result EnableInteractiveQueryOperationResponse, err error) {
	req, err := c.preparerForEnableInteractiveQuery(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "EnableInteractiveQuery", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForEnableInteractiveQuery(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "EnableInteractiveQuery", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// EnableInteractiveQueryThenPoll performs EnableInteractiveQuery then polls until it's completed
func (c IntegrationRuntimeClient) EnableInteractiveQueryThenPoll(ctx context.Context, id IntegrationRuntimeId) error {
	result, err := c.EnableInteractiveQuery(ctx, id)
	if err != nil {
		return fmt.Errorf("performing EnableInteractiveQuery: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after EnableInteractiveQuery: %+v", err)
	}

	return nil
}

// preparerForEnableInteractiveQuery prepares the EnableInteractiveQuery request.
func (c IntegrationRuntimeClient) preparerForEnableInteractiveQuery(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/enableInteractiveQuery", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForEnableInteractiveQuery sends the EnableInteractiveQuery request. The method will close the
// http.Response Body if it receives an error.
func (c IntegrationRuntimeClient) senderForEnableInteractiveQuery(ctx context.Context, req *http.Request) (future EnableInteractiveQueryOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
