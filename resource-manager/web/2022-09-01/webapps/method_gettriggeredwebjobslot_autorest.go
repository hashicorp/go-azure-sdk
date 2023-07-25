package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTriggeredWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *TriggeredWebJob
}

// GetTriggeredWebJobSlot ...
func (c WebAppsClient) GetTriggeredWebJobSlot(ctx context.Context, id SlotTriggeredWebJobId) (result GetTriggeredWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForGetTriggeredWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetTriggeredWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetTriggeredWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetTriggeredWebJobSlot prepares the GetTriggeredWebJobSlot request.
func (c WebAppsClient) preparerForGetTriggeredWebJobSlot(ctx context.Context, id SlotTriggeredWebJobId) (*http.Request, error) {
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

// responderForGetTriggeredWebJobSlot handles the response to the GetTriggeredWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetTriggeredWebJobSlot(resp *http.Response) (result GetTriggeredWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
