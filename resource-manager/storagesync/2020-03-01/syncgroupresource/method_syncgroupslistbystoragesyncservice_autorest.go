package syncgroupresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupsListByStorageSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *SyncGroupArray
}

// SyncGroupsListByStorageSyncService ...
func (c SyncGroupResourceClient) SyncGroupsListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (result SyncGroupsListByStorageSyncServiceOperationResponse, err error) {
	req, err := c.preparerForSyncGroupsListByStorageSyncService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsListByStorageSyncService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncGroupsListByStorageSyncService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsListByStorageSyncService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncGroupsListByStorageSyncService prepares the SyncGroupsListByStorageSyncService request.
func (c SyncGroupResourceClient) preparerForSyncGroupsListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/syncGroups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncGroupsListByStorageSyncService handles the response to the SyncGroupsListByStorageSyncService request. The method always
// closes the http.Response Body.
func (c SyncGroupResourceClient) responderForSyncGroupsListByStorageSyncService(resp *http.Response) (result SyncGroupsListByStorageSyncServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
