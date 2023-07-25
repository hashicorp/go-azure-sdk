package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteInstanceFunctionSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteInstanceFunctionSlot ...
func (c WebAppsClient) DeleteInstanceFunctionSlot(ctx context.Context, id SlotFunctionId) (result DeleteInstanceFunctionSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteInstanceFunctionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceFunctionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceFunctionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteInstanceFunctionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceFunctionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteInstanceFunctionSlot prepares the DeleteInstanceFunctionSlot request.
func (c WebAppsClient) preparerForDeleteInstanceFunctionSlot(ctx context.Context, id SlotFunctionId) (*http.Request, error) {
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

// responderForDeleteInstanceFunctionSlot handles the response to the DeleteInstanceFunctionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteInstanceFunctionSlot(resp *http.Response) (result DeleteInstanceFunctionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
