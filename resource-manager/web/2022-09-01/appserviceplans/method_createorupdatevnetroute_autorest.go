package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateVnetRouteOperationResponse struct {
	HttpResponse *http.Response
	Model        *VnetRoute
}

// CreateOrUpdateVnetRoute ...
func (c AppServicePlansClient) CreateOrUpdateVnetRoute(ctx context.Context, id RouteId, input VnetRoute) (result CreateOrUpdateVnetRouteOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateVnetRoute(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "CreateOrUpdateVnetRoute", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "CreateOrUpdateVnetRoute", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateVnetRoute(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "CreateOrUpdateVnetRoute", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateVnetRoute prepares the CreateOrUpdateVnetRoute request.
func (c AppServicePlansClient) preparerForCreateOrUpdateVnetRoute(ctx context.Context, id RouteId, input VnetRoute) (*http.Request, error) {
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

// responderForCreateOrUpdateVnetRoute handles the response to the CreateOrUpdateVnetRoute request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForCreateOrUpdateVnetRoute(resp *http.Response) (result CreateOrUpdateVnetRouteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
