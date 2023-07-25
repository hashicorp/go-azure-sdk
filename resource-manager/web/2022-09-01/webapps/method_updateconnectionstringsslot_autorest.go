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

type UpdateConnectionStringsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConnectionStringDictionary
}

// UpdateConnectionStringsSlot ...
func (c WebAppsClient) UpdateConnectionStringsSlot(ctx context.Context, id SlotId, input ConnectionStringDictionary) (result UpdateConnectionStringsSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateConnectionStringsSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConnectionStringsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConnectionStringsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateConnectionStringsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConnectionStringsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateConnectionStringsSlot prepares the UpdateConnectionStringsSlot request.
func (c WebAppsClient) preparerForUpdateConnectionStringsSlot(ctx context.Context, id SlotId, input ConnectionStringDictionary) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/connectionStrings", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateConnectionStringsSlot handles the response to the UpdateConnectionStringsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateConnectionStringsSlot(resp *http.Response) (result UpdateConnectionStringsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
