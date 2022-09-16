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

type WorkspaceManagedIdentitySqlControlSettingsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedIdentitySqlControlSettingsModel
}

// WorkspaceManagedIdentitySqlControlSettingsGet ...
func (c WorkspacesClient) WorkspaceManagedIdentitySqlControlSettingsGet(ctx context.Context, id WorkspaceId) (result WorkspaceManagedIdentitySqlControlSettingsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedIdentitySqlControlSettingsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceManagedIdentitySqlControlSettingsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceManagedIdentitySqlControlSettingsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedIdentitySqlControlSettingsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceManagedIdentitySqlControlSettingsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedIdentitySqlControlSettingsGet prepares the WorkspaceManagedIdentitySqlControlSettingsGet request.
func (c WorkspacesClient) preparerForWorkspaceManagedIdentitySqlControlSettingsGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/managedIdentitySqlControlSettings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceManagedIdentitySqlControlSettingsGet handles the response to the WorkspaceManagedIdentitySqlControlSettingsGet request. The method always
// closes the http.Response Body.
func (c WorkspacesClient) responderForWorkspaceManagedIdentitySqlControlSettingsGet(resp *http.Response) (result WorkspaceManagedIdentitySqlControlSettingsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
