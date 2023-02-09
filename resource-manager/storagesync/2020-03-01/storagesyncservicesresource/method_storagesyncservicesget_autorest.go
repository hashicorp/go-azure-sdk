package storagesyncservicesresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncServicesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *StorageSyncService
}

// StorageSyncServicesGet ...
func (c StorageSyncServicesResourceClient) StorageSyncServicesGet(ctx context.Context, id StorageSyncServiceId) (result StorageSyncServicesGetOperationResponse, err error) {
	req, err := c.preparerForStorageSyncServicesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStorageSyncServicesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStorageSyncServicesGet prepares the StorageSyncServicesGet request.
func (c StorageSyncServicesResourceClient) preparerForStorageSyncServicesGet(ctx context.Context, id StorageSyncServiceId) (*http.Request, error) {
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

// responderForStorageSyncServicesGet handles the response to the StorageSyncServicesGet request. The method always
// closes the http.Response Body.
func (c StorageSyncServicesResourceClient) responderForStorageSyncServicesGet(resp *http.Response) (result StorageSyncServicesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
