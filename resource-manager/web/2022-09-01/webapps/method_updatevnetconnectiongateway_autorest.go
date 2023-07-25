package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVnetConnectionGatewayOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// UpdateVnetConnectionGateway ...
func (c WebAppsClient) UpdateVnetConnectionGateway(ctx context.Context, id GatewayId, input VnetGateway) (result UpdateVnetConnectionGatewayOperationResponse, err error) {
	req, err := c.preparerForUpdateVnetConnectionGateway(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionGateway", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionGateway", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateVnetConnectionGateway(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateVnetConnectionGateway", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateVnetConnectionGateway prepares the UpdateVnetConnectionGateway request.
func (c WebAppsClient) preparerForUpdateVnetConnectionGateway(ctx context.Context, id GatewayId, input VnetGateway) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateVnetConnectionGateway handles the response to the UpdateVnetConnectionGateway request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateVnetConnectionGateway(resp *http.Response) (result UpdateVnetConnectionGatewayOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
