package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTriggeredWebJobOperationResponse struct {
	HttpResponse *http.Response
	Model        *TriggeredWebJob
}

// GetTriggeredWebJob ...
func (c WebAppsClient) GetTriggeredWebJob(ctx context.Context, id TriggeredWebJobId) (result GetTriggeredWebJobOperationResponse, err error) {
	req, err := c.preparerForGetTriggeredWebJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetTriggeredWebJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetTriggeredWebJob prepares the GetTriggeredWebJob request.
func (c WebAppsClient) preparerForGetTriggeredWebJob(ctx context.Context, id TriggeredWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetTriggeredWebJob handles the response to the GetTriggeredWebJob request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetTriggeredWebJob(resp *http.Response) (result GetTriggeredWebJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
