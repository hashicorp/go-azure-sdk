package appserviceplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoutesForVnetOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VnetRoute
}

// ListRoutesForVnet ...
func (c AppServicePlansClient) ListRoutesForVnet(ctx context.Context, id ServerFarmVirtualNetworkConnectionId) (result ListRoutesForVnetOperationResponse, err error) {
	req, err := c.preparerForListRoutesForVnet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListRoutesForVnet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListRoutesForVnet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListRoutesForVnet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListRoutesForVnet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListRoutesForVnet prepares the ListRoutesForVnet request.
func (c AppServicePlansClient) preparerForListRoutesForVnet(ctx context.Context, id ServerFarmVirtualNetworkConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/routes", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListRoutesForVnet handles the response to the ListRoutesForVnet request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForListRoutesForVnet(resp *http.Response) (result ListRoutesForVnetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
