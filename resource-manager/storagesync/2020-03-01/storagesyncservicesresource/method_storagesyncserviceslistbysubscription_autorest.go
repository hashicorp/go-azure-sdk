package storagesyncservicesresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageSyncServicesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *StorageSyncServiceArray
}

// StorageSyncServicesListBySubscription ...
func (c StorageSyncServicesResourceClient) StorageSyncServicesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result StorageSyncServicesListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForStorageSyncServicesListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesListBySubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesListBySubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStorageSyncServicesListBySubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesListBySubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStorageSyncServicesListBySubscription prepares the StorageSyncServicesListBySubscription request.
func (c StorageSyncServicesResourceClient) preparerForStorageSyncServicesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.StorageSync/storageSyncServices", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStorageSyncServicesListBySubscription handles the response to the StorageSyncServicesListBySubscription request. The method always
// closes the http.Response Body.
func (c StorageSyncServicesResourceClient) responderForStorageSyncServicesListBySubscription(resp *http.Response) (result StorageSyncServicesListBySubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
