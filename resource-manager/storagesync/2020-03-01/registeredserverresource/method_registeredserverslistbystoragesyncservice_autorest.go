package registeredserverresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredServersListByStorageSyncServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegisteredServerArray
}

// RegisteredServersListByStorageSyncService ...
func (c RegisteredServerResourceClient) RegisteredServersListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (result RegisteredServersListByStorageSyncServiceOperationResponse, err error) {
	req, err := c.preparerForRegisteredServersListByStorageSyncService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersListByStorageSyncService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegisteredServersListByStorageSyncService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersListByStorageSyncService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegisteredServersListByStorageSyncService prepares the RegisteredServersListByStorageSyncService request.
func (c RegisteredServerResourceClient) preparerForRegisteredServersListByStorageSyncService(ctx context.Context, id StorageSyncServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/registeredServers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegisteredServersListByStorageSyncService handles the response to the RegisteredServersListByStorageSyncService request. The method always
// closes the http.Response Body.
func (c RegisteredServerResourceClient) responderForRegisteredServersListByStorageSyncService(resp *http.Response) (result RegisteredServersListByStorageSyncServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
