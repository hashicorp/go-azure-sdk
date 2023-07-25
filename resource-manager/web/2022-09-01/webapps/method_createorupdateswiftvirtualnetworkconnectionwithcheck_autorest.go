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

type CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckOperationResponse struct {
	HttpResponse *http.Response
	Model        *SwiftVirtualNetwork
}

// CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck ...
func (c WebAppsClient) CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck(ctx context.Context, id SiteId, input SwiftVirtualNetwork) (result CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck prepares the CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck request.
func (c WebAppsClient) preparerForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck(ctx context.Context, id SiteId, input SwiftVirtualNetwork) (*http.Request, error) {
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

// responderForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck handles the response to the CreateOrUpdateSwiftVirtualNetworkConnectionWithCheck request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck(resp *http.Response) (result CreateOrUpdateSwiftVirtualNetworkConnectionWithCheckOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
