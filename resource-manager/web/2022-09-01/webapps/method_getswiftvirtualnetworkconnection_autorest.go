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

type GetSwiftVirtualNetworkConnectionOperationResponse struct {
	HttpResponse *http.Response
	Model        *SwiftVirtualNetwork
}

// GetSwiftVirtualNetworkConnection ...
func (c WebAppsClient) GetSwiftVirtualNetworkConnection(ctx context.Context, id SiteId) (result GetSwiftVirtualNetworkConnectionOperationResponse, err error) {
	req, err := c.preparerForGetSwiftVirtualNetworkConnection(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSwiftVirtualNetworkConnection", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSwiftVirtualNetworkConnection", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSwiftVirtualNetworkConnection(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSwiftVirtualNetworkConnection", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSwiftVirtualNetworkConnection prepares the GetSwiftVirtualNetworkConnection request.
func (c WebAppsClient) preparerForGetSwiftVirtualNetworkConnection(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForGetSwiftVirtualNetworkConnection handles the response to the GetSwiftVirtualNetworkConnection request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSwiftVirtualNetworkConnection(resp *http.Response) (result GetSwiftVirtualNetworkConnectionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
