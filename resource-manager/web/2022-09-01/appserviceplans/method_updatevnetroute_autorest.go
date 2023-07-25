package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVnetRouteOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetRoute
}

// UpdateVnetRoute ...
func (c AppServicePlansClient) UpdateVnetRoute(ctx context.Context, id RouteId, input VnetRoute) (result UpdateVnetRouteOperationResponse, err error) {
	req, err := c.preparerForUpdateVnetRoute(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "UpdateVnetRoute", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "UpdateVnetRoute", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateVnetRoute(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "UpdateVnetRoute", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateVnetRoute prepares the UpdateVnetRoute request.
func (c AppServicePlansClient) preparerForUpdateVnetRoute(ctx context.Context, id RouteId, input VnetRoute) (*http.Request, error) {
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

// responderForUpdateVnetRoute handles the response to the UpdateVnetRoute request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForUpdateVnetRoute(resp *http.Response) (result UpdateVnetRouteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
