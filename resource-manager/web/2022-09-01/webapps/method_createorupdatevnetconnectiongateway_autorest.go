package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateVnetConnectionGatewayOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// CreateOrUpdateVnetConnectionGateway ...
func (c WebAppsClient) CreateOrUpdateVnetConnectionGateway(ctx context.Context, id GatewayId, input VnetGateway) (result CreateOrUpdateVnetConnectionGatewayOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateVnetConnectionGateway(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionGateway", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionGateway", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateVnetConnectionGateway(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateVnetConnectionGateway", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateVnetConnectionGateway prepares the CreateOrUpdateVnetConnectionGateway request.
func (c WebAppsClient) preparerForCreateOrUpdateVnetConnectionGateway(ctx context.Context, id GatewayId, input VnetGateway) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateVnetConnectionGateway handles the response to the CreateOrUpdateVnetConnectionGateway request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateVnetConnectionGateway(resp *http.Response) (result CreateOrUpdateVnetConnectionGatewayOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
