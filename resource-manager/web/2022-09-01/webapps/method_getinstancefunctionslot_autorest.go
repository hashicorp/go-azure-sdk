package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceFunctionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *FunctionEnvelope
}

// GetInstanceFunctionSlot ...
func (c WebAppsClient) GetInstanceFunctionSlot(ctx context.Context, id SlotFunctionId) (result GetInstanceFunctionSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceFunctionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceFunctionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceFunctionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceFunctionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceFunctionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceFunctionSlot prepares the GetInstanceFunctionSlot request.
func (c WebAppsClient) preparerForGetInstanceFunctionSlot(ctx context.Context, id SlotFunctionId) (*http.Request, error) {
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

// responderForGetInstanceFunctionSlot handles the response to the GetInstanceFunctionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceFunctionSlot(resp *http.Response) (result GetInstanceFunctionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
