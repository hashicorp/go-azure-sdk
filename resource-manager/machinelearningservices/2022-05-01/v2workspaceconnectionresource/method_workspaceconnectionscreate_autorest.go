package v2workspaceconnectionresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionsCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceConnectionPropertiesV2BasicResource
}

// WorkspaceConnectionsCreate ...
func (c V2WorkspaceConnectionResourceClient) WorkspaceConnectionsCreate(ctx context.Context, id ConnectionId, input WorkspaceConnectionPropertiesV2BasicResource) (result WorkspaceConnectionsCreateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceConnectionsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceConnectionsCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient", "WorkspaceConnectionsCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceConnectionsCreate prepares the WorkspaceConnectionsCreate request.
func (c V2WorkspaceConnectionResourceClient) preparerForWorkspaceConnectionsCreate(ctx context.Context, id ConnectionId, input WorkspaceConnectionPropertiesV2BasicResource) (*http.Request, error) {
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

// responderForWorkspaceConnectionsCreate handles the response to the WorkspaceConnectionsCreate request. The method always
// closes the http.Response Body.
func (c V2WorkspaceConnectionResourceClient) responderForWorkspaceConnectionsCreate(resp *http.Response) (result WorkspaceConnectionsCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
