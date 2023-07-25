package appserviceplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVnetRouteOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteVnetRoute ...
func (c AppServicePlansClient) DeleteVnetRoute(ctx context.Context, id RouteId) (result DeleteVnetRouteOperationResponse, err error) {
	req, err := c.preparerForDeleteVnetRoute(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "DeleteVnetRoute", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "DeleteVnetRoute", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteVnetRoute(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "DeleteVnetRoute", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteVnetRoute prepares the DeleteVnetRoute request.
func (c AppServicePlansClient) preparerForDeleteVnetRoute(ctx context.Context, id RouteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteVnetRoute handles the response to the DeleteVnetRoute request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForDeleteVnetRoute(resp *http.Response) (result DeleteVnetRouteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
