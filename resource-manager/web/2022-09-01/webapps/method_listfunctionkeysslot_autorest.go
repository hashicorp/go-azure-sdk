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

type ListFunctionKeysSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListFunctionKeysSlot ...
func (c WebAppsClient) ListFunctionKeysSlot(ctx context.Context, id SlotFunctionId) (result ListFunctionKeysSlotOperationResponse, err error) {
	req, err := c.preparerForListFunctionKeysSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionKeysSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionKeysSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListFunctionKeysSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionKeysSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListFunctionKeysSlot prepares the ListFunctionKeysSlot request.
func (c WebAppsClient) preparerForListFunctionKeysSlot(ctx context.Context, id SlotFunctionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListFunctionKeysSlot handles the response to the ListFunctionKeysSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListFunctionKeysSlot(resp *http.Response) (result ListFunctionKeysSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
