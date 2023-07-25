package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVnetGatewayOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// UpdateVnetGateway ...
func (c AppServicePlansClient) UpdateVnetGateway(ctx context.Context, id VirtualNetworkConnectionGatewayId, input VnetGateway) (result UpdateVnetGatewayOperationResponse, err error) {
	req, err := c.preparerForUpdateVnetGateway(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "UpdateVnetGateway", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "UpdateVnetGateway", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateVnetGateway(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "UpdateVnetGateway", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateVnetGateway prepares the UpdateVnetGateway request.
func (c AppServicePlansClient) preparerForUpdateVnetGateway(ctx context.Context, id VirtualNetworkConnectionGatewayId, input VnetGateway) (*http.Request, error) {
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

// responderForUpdateVnetGateway handles the response to the UpdateVnetGateway request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForUpdateVnetGateway(resp *http.Response) (result UpdateVnetGatewayOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
