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

type ListVnetConnectionsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VnetInfoResource
}

// ListVnetConnectionsSlot ...
func (c WebAppsClient) ListVnetConnectionsSlot(ctx context.Context, id SlotId) (result ListVnetConnectionsSlotOperationResponse, err error) {
	req, err := c.preparerForListVnetConnectionsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListVnetConnectionsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListVnetConnectionsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListVnetConnectionsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListVnetConnectionsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListVnetConnectionsSlot prepares the ListVnetConnectionsSlot request.
func (c WebAppsClient) preparerForListVnetConnectionsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/virtualNetworkConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListVnetConnectionsSlot handles the response to the ListVnetConnectionsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListVnetConnectionsSlot(resp *http.Response) (result ListVnetConnectionsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
