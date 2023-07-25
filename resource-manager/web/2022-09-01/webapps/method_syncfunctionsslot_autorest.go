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

type SyncFunctionsSlotOperationResponse struct {
	HttpResponse *http.Response
}

// SyncFunctionsSlot ...
func (c WebAppsClient) SyncFunctionsSlot(ctx context.Context, id SlotId) (result SyncFunctionsSlotOperationResponse, err error) {
	req, err := c.preparerForSyncFunctionsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctionsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctionsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncFunctionsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctionsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncFunctionsSlot prepares the SyncFunctionsSlot request.
func (c WebAppsClient) preparerForSyncFunctionsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/host/default/sync", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncFunctionsSlot handles the response to the SyncFunctionsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForSyncFunctionsSlot(resp *http.Response) (result SyncFunctionsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
