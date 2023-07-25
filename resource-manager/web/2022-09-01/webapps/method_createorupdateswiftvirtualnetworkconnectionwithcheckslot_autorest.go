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

type CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SwiftVirtualNetwork
}

// CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot ...
func (c WebAppsClient) CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot(ctx context.Context, id SlotId, input SwiftVirtualNetwork) (result CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot prepares the CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot(ctx context.Context, id SlotId, input SwiftVirtualNetwork) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkConfig/virtualNetwork", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot handles the response to the CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot(resp *http.Response) (result CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
