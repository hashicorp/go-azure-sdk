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

type ListConnectionStringsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConnectionStringDictionary
}

// ListConnectionStringsSlot ...
func (c WebAppsClient) ListConnectionStringsSlot(ctx context.Context, id SlotId) (result ListConnectionStringsSlotOperationResponse, err error) {
	req, err := c.preparerForListConnectionStringsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConnectionStringsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConnectionStringsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListConnectionStringsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConnectionStringsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListConnectionStringsSlot prepares the ListConnectionStringsSlot request.
func (c WebAppsClient) preparerForListConnectionStringsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/connectionStrings/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListConnectionStringsSlot handles the response to the ListConnectionStringsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListConnectionStringsSlot(resp *http.Response) (result ListConnectionStringsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
