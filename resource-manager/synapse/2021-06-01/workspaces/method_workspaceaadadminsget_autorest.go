package workspaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceAadAdminsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceAadAdminInfo
}

// WorkspaceAadAdminsGet ...
func (c WorkspacesClient) WorkspaceAadAdminsGet(ctx context.Context, id WorkspaceId) (result WorkspaceAadAdminsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceAadAdminsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceAadAdminsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceAadAdminsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceAadAdminsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceAadAdminsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceAadAdminsGet prepares the WorkspaceAadAdminsGet request.
func (c WorkspacesClient) preparerForWorkspaceAadAdminsGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/administrators/activeDirectory", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceAadAdminsGet handles the response to the WorkspaceAadAdminsGet request. The method always
// closes the http.Response Body.
func (c WorkspacesClient) responderForWorkspaceAadAdminsGet(resp *http.Response) (result WorkspaceAadAdminsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
