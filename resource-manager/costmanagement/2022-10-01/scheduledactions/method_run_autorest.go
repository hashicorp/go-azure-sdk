package scheduledactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunOperationResponse struct {
	HttpResponse *http.Response
}

// Run ...
func (c ScheduledActionsClient) Run(ctx context.Context, id ScheduledActionId) (result RunOperationResponse, err error) {
	req, err := c.preparerForRun(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "Run", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "Run", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRun(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "Run", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRun prepares the Run request.
func (c ScheduledActionsClient) preparerForRun(ctx context.Context, id ScheduledActionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/execute", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRun handles the response to the Run request. The method always
// closes the http.Response Body.
func (c ScheduledActionsClient) responderForRun(resp *http.Response) (result RunOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
