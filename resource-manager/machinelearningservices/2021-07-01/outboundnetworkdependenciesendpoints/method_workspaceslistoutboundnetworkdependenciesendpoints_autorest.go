package outboundnetworkdependenciesendpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacesListOutboundNetworkDependenciesEndpointsOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExternalFQDNResponse
}

// WorkspacesListOutboundNetworkDependenciesEndpoints ...
func (c OutboundNetworkDependenciesEndpointsClient) WorkspacesListOutboundNetworkDependenciesEndpoints(ctx context.Context, id WorkspaceId) (result WorkspacesListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	req, err := c.preparerForWorkspacesListOutboundNetworkDependenciesEndpoints(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "WorkspacesListOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "WorkspacesListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspacesListOutboundNetworkDependenciesEndpoints(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "WorkspacesListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspacesListOutboundNetworkDependenciesEndpoints prepares the WorkspacesListOutboundNetworkDependenciesEndpoints request.
func (c OutboundNetworkDependenciesEndpointsClient) preparerForWorkspacesListOutboundNetworkDependenciesEndpoints(ctx context.Context, id WorkspaceId) (*http.Request, error) {
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

// responderForWorkspacesListOutboundNetworkDependenciesEndpoints handles the response to the WorkspacesListOutboundNetworkDependenciesEndpoints request. The method always
// closes the http.Response Body.
func (c OutboundNetworkDependenciesEndpointsClient) responderForWorkspacesListOutboundNetworkDependenciesEndpoints(resp *http.Response) (result WorkspacesListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
