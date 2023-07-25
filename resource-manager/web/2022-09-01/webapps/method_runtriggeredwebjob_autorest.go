package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunTriggeredWebJobOperationResponse struct {
	HttpResponse *http.Response
}

// RunTriggeredWebJob ...
func (c WebAppsClient) RunTriggeredWebJob(ctx context.Context, id TriggeredWebJobId) (result RunTriggeredWebJobOperationResponse, err error) {
	req, err := c.preparerForRunTriggeredWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RunTriggeredWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RunTriggeredWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRunTriggeredWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RunTriggeredWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRunTriggeredWebJob prepares the RunTriggeredWebJob request.
func (c WebAppsClient) preparerForRunTriggeredWebJob(ctx context.Context, id TriggeredWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/run", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRunTriggeredWebJob handles the response to the RunTriggeredWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForRunTriggeredWebJob(resp *http.Response) (result RunTriggeredWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
