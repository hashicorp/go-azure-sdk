package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteTriggeredWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteTriggeredWebJobSlot ...
func (c WebAppsClient) DeleteTriggeredWebJobSlot(ctx context.Context, id SlotTriggeredWebJobId) (result DeleteTriggeredWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteTriggeredWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteTriggeredWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteTriggeredWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteTriggeredWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteTriggeredWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteTriggeredWebJobSlot prepares the DeleteTriggeredWebJobSlot request.
func (c WebAppsClient) preparerForDeleteTriggeredWebJobSlot(ctx context.Context, id SlotTriggeredWebJobId) (*http.Request, error) {
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

// responderForDeleteTriggeredWebJobSlot handles the response to the DeleteTriggeredWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteTriggeredWebJobSlot(resp *http.Response) (result DeleteTriggeredWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
