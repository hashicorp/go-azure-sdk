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

type StorageSyncServicesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *StorageSyncServiceArray
}

// StorageSyncServicesListByResourceGroup ...
func (c StorageSyncServicesResourceClient) StorageSyncServicesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result StorageSyncServicesListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForStorageSyncServicesListByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesListByResourceGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesListByResourceGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStorageSyncServicesListByResourceGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesyncservicesresource.StorageSyncServicesResourceClient", "StorageSyncServicesListByResourceGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStorageSyncServicesListByResourceGroup prepares the StorageSyncServicesListByResourceGroup request.
func (c StorageSyncServicesResourceClient) preparerForStorageSyncServicesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
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

// responderForStorageSyncServicesListByResourceGroup handles the response to the StorageSyncServicesListByResourceGroup request. The method always
// closes the http.Response Body.
func (c StorageSyncServicesResourceClient) responderForStorageSyncServicesListByResourceGroup(resp *http.Response) (result StorageSyncServicesListByResourceGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
