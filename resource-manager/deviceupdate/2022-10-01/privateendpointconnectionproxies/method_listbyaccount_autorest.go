package privateendpointconnectionproxies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionProxyListResult
}

// ListByAccount ...
func (c PrivateEndpointConnectionProxiesClient) ListByAccount(ctx context.Context, id AccountId) (result ListByAccountOperationResponse, err error) {
	req, err := c.preparerForListByAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionproxies.PrivateEndpointConnectionProxiesClient", "ListByAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionproxies.PrivateEndpointConnectionProxiesClient", "ListByAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnectionproxies.PrivateEndpointConnectionProxiesClient", "ListByAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByAccount prepares the ListByAccount request.
func (c PrivateEndpointConnectionProxiesClient) preparerForListByAccount(ctx context.Context, id AccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateEndpointConnectionProxies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByAccount handles the response to the ListByAccount request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionProxiesClient) responderForListByAccount(resp *http.Response) (result ListByAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
