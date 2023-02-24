package integrationruntimes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOutboundNetworkDependenciesEndpointsOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse
}

// ListOutboundNetworkDependenciesEndpoints ...
func (c IntegrationRuntimesClient) ListOutboundNetworkDependenciesEndpoints(ctx context.Context, id IntegrationRuntimeId) (result ListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	req, err := c.preparerForListOutboundNetworkDependenciesEndpoints(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "ListOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "ListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListOutboundNetworkDependenciesEndpoints(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "ListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListOutboundNetworkDependenciesEndpoints prepares the ListOutboundNetworkDependenciesEndpoints request.
func (c IntegrationRuntimesClient) preparerForListOutboundNetworkDependenciesEndpoints(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/outboundNetworkDependenciesEndpoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListOutboundNetworkDependenciesEndpoints handles the response to the ListOutboundNetworkDependenciesEndpoints request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForListOutboundNetworkDependenciesEndpoints(resp *http.Response) (result ListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
