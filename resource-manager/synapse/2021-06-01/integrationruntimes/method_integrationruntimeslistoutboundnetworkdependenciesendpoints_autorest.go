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

type IntegrationRuntimesListOutboundNetworkDependenciesEndpointsOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse
}

// IntegrationRuntimesListOutboundNetworkDependenciesEndpoints ...
func (c IntegrationRuntimesClient) IntegrationRuntimesListOutboundNetworkDependenciesEndpoints(ctx context.Context, id IntegrationRuntimeId) (result IntegrationRuntimesListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	req, err := c.preparerForIntegrationRuntimesListOutboundNetworkDependenciesEndpoints(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "IntegrationRuntimesListOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "IntegrationRuntimesListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIntegrationRuntimesListOutboundNetworkDependenciesEndpoints(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "IntegrationRuntimesListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIntegrationRuntimesListOutboundNetworkDependenciesEndpoints prepares the IntegrationRuntimesListOutboundNetworkDependenciesEndpoints request.
func (c IntegrationRuntimesClient) preparerForIntegrationRuntimesListOutboundNetworkDependenciesEndpoints(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
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

// responderForIntegrationRuntimesListOutboundNetworkDependenciesEndpoints handles the response to the IntegrationRuntimesListOutboundNetworkDependenciesEndpoints request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForIntegrationRuntimesListOutboundNetworkDependenciesEndpoints(resp *http.Response) (result IntegrationRuntimesListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
