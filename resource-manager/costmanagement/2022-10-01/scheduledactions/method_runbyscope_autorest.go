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

type RunByScopeOperationResponse struct {
	HttpResponse *http.Response
}

// RunByScope ...
func (c ScheduledActionsClient) RunByScope(ctx context.Context, id ScopedScheduledActionId) (result RunByScopeOperationResponse, err error) {
	req, err := c.preparerForRunByScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "RunByScope", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "RunByScope", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRunByScope(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduledactions.ScheduledActionsClient", "RunByScope", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRunByScope prepares the RunByScope request.
func (c ScheduledActionsClient) preparerForRunByScope(ctx context.Context, id ScopedScheduledActionId) (*http.Request, error) {
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

// responderForRunByScope handles the response to the RunByScope request. The method always
// closes the http.Response Body.
func (c ScheduledActionsClient) responderForRunByScope(resp *http.Response) (result RunByScopeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
