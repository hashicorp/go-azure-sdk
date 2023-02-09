package privatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByStorageSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResourceListResult
}

// ListByStorageSyncService ...
func (c PrivateLinkResourcesClient) ListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (result ListByStorageSyncServiceOperationResponse, err error) {
	req, err := c.preparerForListByStorageSyncService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresources.PrivateLinkResourcesClient", "ListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresources.PrivateLinkResourcesClient", "ListByStorageSyncService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByStorageSyncService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresources.PrivateLinkResourcesClient", "ListByStorageSyncService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByStorageSyncService prepares the ListByStorageSyncService request.
func (c PrivateLinkResourcesClient) preparerForListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateLinkResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByStorageSyncService handles the response to the ListByStorageSyncService request. The method always
// closes the http.Response Body.
func (c PrivateLinkResourcesClient) responderForListByStorageSyncService(resp *http.Response) (result ListByStorageSyncServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
