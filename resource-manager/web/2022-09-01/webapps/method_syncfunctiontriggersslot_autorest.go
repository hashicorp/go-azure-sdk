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

type SyncFunctionTriggersSlotOperationResponse struct {
	HttpResponse *http.Response
}

// SyncFunctionTriggersSlot ...
func (c WebAppsClient) SyncFunctionTriggersSlot(ctx context.Context, id SlotId) (result SyncFunctionTriggersSlotOperationResponse, err error) {
	req, err := c.preparerForSyncFunctionTriggersSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctionTriggersSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctionTriggersSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncFunctionTriggersSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctionTriggersSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncFunctionTriggersSlot prepares the SyncFunctionTriggersSlot request.
func (c WebAppsClient) preparerForSyncFunctionTriggersSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/syncfunctiontriggers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncFunctionTriggersSlot handles the response to the SyncFunctionTriggersSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForSyncFunctionTriggersSlot(resp *http.Response) (result SyncFunctionTriggersSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
