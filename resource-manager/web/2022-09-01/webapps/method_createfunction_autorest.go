package webapps

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

type CreateFunctionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateFunction ...
func (c WebAppsClient) CreateFunction(ctx context.Context, id FunctionId, input FunctionEnvelope) (result CreateFunctionOperationResponse, err error) {
	req, err := c.preparerForCreateFunction(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateFunction", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateFunction(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateFunction", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateFunctionThenPoll performs CreateFunction then polls until it's completed
func (c WebAppsClient) CreateFunctionThenPoll(ctx context.Context, id FunctionId, input FunctionEnvelope) error {
	result, err := c.CreateFunction(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateFunction: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateFunction: %+v", err)
	}

	return nil
}

// preparerForCreateFunction prepares the CreateFunction request.
func (c WebAppsClient) preparerForCreateFunction(ctx context.Context, id FunctionId, input FunctionEnvelope) (*http.Request, error) {
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

// senderForCreateFunction sends the CreateFunction request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForCreateFunction(ctx context.Context, req *http.Request) (future CreateFunctionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
