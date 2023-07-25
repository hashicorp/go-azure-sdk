package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateEndpointConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *RemotePrivateEndpointConnectionARMResource
}

// GetPrivateEndpointConnectionSlot ...
func (c WebAppsClient) GetPrivateEndpointConnectionSlot(ctx context.Context, id SlotPrivateEndpointConnectionId) (result GetPrivateEndpointConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForGetPrivateEndpointConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateEndpointConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateEndpointConnectionSlot prepares the GetPrivateEndpointConnectionSlot request.
func (c WebAppsClient) preparerForGetPrivateEndpointConnectionSlot(ctx context.Context, id SlotPrivateEndpointConnectionId) (*http.Request, error) {
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

// responderForGetPrivateEndpointConnectionSlot handles the response to the GetPrivateEndpointConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPrivateEndpointConnectionSlot(resp *http.Response) (result GetPrivateEndpointConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
