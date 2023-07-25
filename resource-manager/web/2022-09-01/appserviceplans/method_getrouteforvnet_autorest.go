package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRouteForVnetOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VnetRoute
}

// GetRouteForVnet ...
func (c AppServicePlansClient) GetRouteForVnet(ctx context.Context, id RouteId) (result GetRouteForVnetOperationResponse, err error) {
	req, err := c.preparerForGetRouteForVnet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetRouteForVnet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetRouteForVnet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRouteForVnet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "GetRouteForVnet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRouteForVnet prepares the GetRouteForVnet request.
func (c AppServicePlansClient) preparerForGetRouteForVnet(ctx context.Context, id RouteId) (*http.Request, error) {
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

// responderForGetRouteForVnet handles the response to the GetRouteForVnet request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForGetRouteForVnet(resp *http.Response) (result GetRouteForVnetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
