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

type RunTriggeredWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
}

// RunTriggeredWebJobSlot ...
func (c WebAppsClient) RunTriggeredWebJobSlot(ctx context.Context, id SlotTriggeredWebJobId) (result RunTriggeredWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForRunTriggeredWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RunTriggeredWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RunTriggeredWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRunTriggeredWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RunTriggeredWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRunTriggeredWebJobSlot prepares the RunTriggeredWebJobSlot request.
func (c WebAppsClient) preparerForRunTriggeredWebJobSlot(ctx context.Context, id SlotTriggeredWebJobId) (*http.Request, error) {
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

// responderForRunTriggeredWebJobSlot handles the response to the RunTriggeredWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForRunTriggeredWebJobSlot(resp *http.Response) (result RunTriggeredWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
