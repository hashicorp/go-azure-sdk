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

type WorkspaceSqlAadAdminsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceAadAdminInfo
}

// WorkspaceSqlAadAdminsGet ...
func (c WorkspacesClient) WorkspaceSqlAadAdminsGet(ctx context.Context, id WorkspaceId) (result WorkspaceSqlAadAdminsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSqlAadAdminsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceSqlAadAdminsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceSqlAadAdminsGet prepares the WorkspaceSqlAadAdminsGet request.
func (c WorkspacesClient) preparerForWorkspaceSqlAadAdminsGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sqlAdministrators/activeDirectory", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceSqlAadAdminsGet handles the response to the WorkspaceSqlAadAdminsGet request. The method always
// closes the http.Response Body.
func (c WorkspacesClient) responderForWorkspaceSqlAadAdminsGet(resp *http.Response) (result WorkspaceSqlAadAdminsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
