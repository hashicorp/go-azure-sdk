package integrationruntime

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SelfHostedIntegrationRuntimeNode
}

// NodesGet ...
func (c IntegrationRuntimeClient) NodesGet(ctx context.Context, id NodeId) (result NodesGetOperationResponse, err error) {
	req, err := c.preparerForNodesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNodesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNodesGet prepares the NodesGet request.
func (c IntegrationRuntimeClient) preparerForNodesGet(ctx context.Context, id NodeId) (*http.Request, error) {
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

// responderForNodesGet handles the response to the NodesGet request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForNodesGet(resp *http.Response) (result NodesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
