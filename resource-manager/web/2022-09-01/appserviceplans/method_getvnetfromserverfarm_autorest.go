package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVnetFromServerFarmOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetInfoResource
}

// GetVnetFromServerFarm ...
func (c AppServicePlansClient) GetVnetFromServerFarm(ctx context.Context, id ServerFarmVirtualNetworkConnectionId) (result GetVnetFromServerFarmOperationResponse, err error) {
	req, err := c.preparerForGetVnetFromServerFarm(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetVnetFromServerFarm", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetVnetFromServerFarm", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVnetFromServerFarm(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetVnetFromServerFarm", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVnetFromServerFarm prepares the GetVnetFromServerFarm request.
func (c AppServicePlansClient) preparerForGetVnetFromServerFarm(ctx context.Context, id ServerFarmVirtualNetworkConnectionId) (*http.Request, error) {
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

// responderForGetVnetFromServerFarm handles the response to the GetVnetFromServerFarm request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForGetVnetFromServerFarm(resp *http.Response) (result GetVnetFromServerFarmOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
