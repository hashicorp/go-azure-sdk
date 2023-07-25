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

type GetSwiftVirtualNetworkConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SwiftVirtualNetwork
}

// GetSwiftVirtualNetworkConnectionSlot ...
func (c WebAppsClient) GetSwiftVirtualNetworkConnectionSlot(ctx context.Context, id SlotId) (result GetSwiftVirtualNetworkConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForGetSwiftVirtualNetworkConnectionSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSwiftVirtualNetworkConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSwiftVirtualNetworkConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSwiftVirtualNetworkConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSwiftVirtualNetworkConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSwiftVirtualNetworkConnectionSlot prepares the GetSwiftVirtualNetworkConnectionSlot request.
func (c WebAppsClient) preparerForGetSwiftVirtualNetworkConnectionSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkConfig/virtualNetwork", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSwiftVirtualNetworkConnectionSlot handles the response to the GetSwiftVirtualNetworkConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSwiftVirtualNetworkConnectionSlot(resp *http.Response) (result GetSwiftVirtualNetworkConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
