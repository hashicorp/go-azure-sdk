package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListWorkflowsConnectionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkflowEnvelope
}

// ListWorkflowsConnections ...
func (c WebAppsClient) ListWorkflowsConnections(ctx context.Context, id SiteId) (result ListWorkflowsConnectionsOperationResponse, err error) {
	req, err := c.preparerForListWorkflowsConnections(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflowsConnections", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflowsConnections", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListWorkflowsConnections(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListWorkflowsConnections", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListWorkflowsConnections prepares the ListWorkflowsConnections request.
func (c WebAppsClient) preparerForListWorkflowsConnections(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listWorkflowsConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListWorkflowsConnections handles the response to the ListWorkflowsConnections request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListWorkflowsConnections(resp *http.Response) (result ListWorkflowsConnectionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
