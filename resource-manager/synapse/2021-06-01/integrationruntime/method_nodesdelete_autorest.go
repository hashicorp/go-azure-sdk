package integrationruntime

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodesDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// NodesDelete ...
func (c IntegrationRuntimeClient) NodesDelete(ctx context.Context, id NodeId) (result NodesDeleteOperationResponse, err error) {
	req, err := c.preparerForNodesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNodesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNodesDelete prepares the NodesDelete request.
func (c IntegrationRuntimeClient) preparerForNodesDelete(ctx context.Context, id NodeId) (*http.Request, error) {
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

// responderForNodesDelete handles the response to the NodesDelete request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForNodesDelete(resp *http.Response) (result NodesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
