package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTriggeredWebJobHistoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *TriggeredJobHistory
}

// GetTriggeredWebJobHistory ...
func (c WebAppsClient) GetTriggeredWebJobHistory(ctx context.Context, id HistoryId) (result GetTriggeredWebJobHistoryOperationResponse, err error) {
	req, err := c.preparerForGetTriggeredWebJobHistory(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobHistory", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobHistory", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetTriggeredWebJobHistory(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobHistory", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetTriggeredWebJobHistory prepares the GetTriggeredWebJobHistory request.
func (c WebAppsClient) preparerForGetTriggeredWebJobHistory(ctx context.Context, id HistoryId) (*http.Request, error) {
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

// responderForGetTriggeredWebJobHistory handles the response to the GetTriggeredWebJobHistory request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetTriggeredWebJobHistory(resp *http.Response) (result GetTriggeredWebJobHistoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
