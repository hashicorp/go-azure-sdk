package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteContinuousWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteContinuousWebJobSlot ...
func (c WebAppsClient) DeleteContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (result DeleteContinuousWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteContinuousWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteContinuousWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteContinuousWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteContinuousWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteContinuousWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteContinuousWebJobSlot prepares the DeleteContinuousWebJobSlot request.
func (c WebAppsClient) preparerForDeleteContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (*http.Request, error) {
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

// responderForDeleteContinuousWebJobSlot handles the response to the DeleteContinuousWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteContinuousWebJobSlot(resp *http.Response) (result DeleteContinuousWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
