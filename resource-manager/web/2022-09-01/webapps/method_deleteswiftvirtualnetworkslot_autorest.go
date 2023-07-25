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

type DeleteSwiftVirtualNetworkSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteSwiftVirtualNetworkSlot ...
func (c WebAppsClient) DeleteSwiftVirtualNetworkSlot(ctx context.Context, id SlotId) (result DeleteSwiftVirtualNetworkSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteSwiftVirtualNetworkSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSwiftVirtualNetworkSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSwiftVirtualNetworkSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSwiftVirtualNetworkSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSwiftVirtualNetworkSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSwiftVirtualNetworkSlot prepares the DeleteSwiftVirtualNetworkSlot request.
func (c WebAppsClient) preparerForDeleteSwiftVirtualNetworkSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkConfig/virtualNetwork", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteSwiftVirtualNetworkSlot handles the response to the DeleteSwiftVirtualNetworkSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSwiftVirtualNetworkSlot(resp *http.Response) (result DeleteSwiftVirtualNetworkSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
