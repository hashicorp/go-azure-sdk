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

type DeleteSwiftVirtualNetworkOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteSwiftVirtualNetwork ...
func (c WebAppsClient) DeleteSwiftVirtualNetwork(ctx context.Context, id SiteId) (result DeleteSwiftVirtualNetworkOperationResponse, err error) {
	req, err := c.preparerForDeleteSwiftVirtualNetwork(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSwiftVirtualNetwork", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSwiftVirtualNetwork", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSwiftVirtualNetwork(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSwiftVirtualNetwork", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSwiftVirtualNetwork prepares the DeleteSwiftVirtualNetwork request.
func (c WebAppsClient) preparerForDeleteSwiftVirtualNetwork(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForDeleteSwiftVirtualNetwork handles the response to the DeleteSwiftVirtualNetwork request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSwiftVirtualNetwork(resp *http.Response) (result DeleteSwiftVirtualNetworkOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
