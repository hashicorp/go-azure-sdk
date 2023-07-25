package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVnetConnectionGatewayOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// GetVnetConnectionGateway ...
func (c WebAppsClient) GetVnetConnectionGateway(ctx context.Context, id GatewayId) (result GetVnetConnectionGatewayOperationResponse, err error) {
	req, err := c.preparerForGetVnetConnectionGateway(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionGateway", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionGateway", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVnetConnectionGateway(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetVnetConnectionGateway", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVnetConnectionGateway prepares the GetVnetConnectionGateway request.
func (c WebAppsClient) preparerForGetVnetConnectionGateway(ctx context.Context, id GatewayId) (*http.Request, error) {
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

// responderForGetVnetConnectionGateway handles the response to the GetVnetConnectionGateway request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetVnetConnectionGateway(resp *http.Response) (result GetVnetConnectionGatewayOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
