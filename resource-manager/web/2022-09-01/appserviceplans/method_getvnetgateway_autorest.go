package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVnetGatewayOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetGateway
}

// GetVnetGateway ...
func (c AppServicePlansClient) GetVnetGateway(ctx context.Context, id VirtualNetworkConnectionGatewayId) (result GetVnetGatewayOperationResponse, err error) {
	req, err := c.preparerForGetVnetGateway(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetVnetGateway", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetVnetGateway", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVnetGateway(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetVnetGateway", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVnetGateway prepares the GetVnetGateway request.
func (c AppServicePlansClient) preparerForGetVnetGateway(ctx context.Context, id VirtualNetworkConnectionGatewayId) (*http.Request, error) {
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

// responderForGetVnetGateway handles the response to the GetVnetGateway request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForGetVnetGateway(resp *http.Response) (result GetVnetGatewayOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
