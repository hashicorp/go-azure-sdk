package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTriggeredWebJobHistorySlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *TriggeredJobHistory
}

// GetTriggeredWebJobHistorySlot ...
func (c WebAppsClient) GetTriggeredWebJobHistorySlot(ctx context.Context, id TriggeredWebJobHistoryId) (result GetTriggeredWebJobHistorySlotOperationResponse, err error) {
	req, err := c.preparerForGetTriggeredWebJobHistorySlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobHistorySlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobHistorySlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetTriggeredWebJobHistorySlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobHistorySlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetTriggeredWebJobHistorySlot prepares the GetTriggeredWebJobHistorySlot request.
func (c WebAppsClient) preparerForGetTriggeredWebJobHistorySlot(ctx context.Context, id TriggeredWebJobHistoryId) (*http.Request, error) {
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

// responderForGetTriggeredWebJobHistorySlot handles the response to the GetTriggeredWebJobHistorySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetTriggeredWebJobHistorySlot(resp *http.Response) (result GetTriggeredWebJobHistorySlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
