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

type SyncRepositorySlotOperationResponse struct {
	HttpResponse *http.Response
}

// SyncRepositorySlot ...
func (c WebAppsClient) SyncRepositorySlot(ctx context.Context, id SlotId) (result SyncRepositorySlotOperationResponse, err error) {
	req, err := c.preparerForSyncRepositorySlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncRepositorySlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncRepositorySlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncRepositorySlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncRepositorySlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncRepositorySlot prepares the SyncRepositorySlot request.
func (c WebAppsClient) preparerForSyncRepositorySlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sync", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncRepositorySlot handles the response to the SyncRepositorySlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForSyncRepositorySlot(resp *http.Response) (result SyncRepositorySlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
