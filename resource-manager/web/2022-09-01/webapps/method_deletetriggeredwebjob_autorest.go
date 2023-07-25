package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteTriggeredWebJobOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteTriggeredWebJob ...
func (c WebAppsClient) DeleteTriggeredWebJob(ctx context.Context, id TriggeredWebJobId) (result DeleteTriggeredWebJobOperationResponse, err error) {
	req, err := c.preparerForDeleteTriggeredWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteTriggeredWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteTriggeredWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteTriggeredWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteTriggeredWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteTriggeredWebJob prepares the DeleteTriggeredWebJob request.
func (c WebAppsClient) preparerForDeleteTriggeredWebJob(ctx context.Context, id TriggeredWebJobId) (*http.Request, error) {
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

// responderForDeleteTriggeredWebJob handles the response to the DeleteTriggeredWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteTriggeredWebJob(resp *http.Response) (result DeleteTriggeredWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
