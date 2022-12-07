package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayRouteConfigsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *GatewayRouteConfigResource
}

// GatewayRouteConfigsGet ...
func (c AppPlatformClient) GatewayRouteConfigsGet(ctx context.Context, id RouteConfigId) (result GatewayRouteConfigsGetOperationResponse, err error) {
	req, err := c.preparerForGatewayRouteConfigsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewayRouteConfigsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewayRouteConfigsGet prepares the GatewayRouteConfigsGet request.
func (c AppPlatformClient) preparerForGatewayRouteConfigsGet(ctx context.Context, id RouteConfigId) (*http.Request, error) {
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

// responderForGatewayRouteConfigsGet handles the response to the GatewayRouteConfigsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewayRouteConfigsGet(resp *http.Response) (result GatewayRouteConfigsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
