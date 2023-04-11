package workspacesettings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// WorkspaceSettingsDelete ...
func (c WorkspaceSettingsClient) WorkspaceSettingsDelete(ctx context.Context, id WorkspaceSettingId) (result WorkspaceSettingsDeleteOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSettingsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceSettingsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceSettingsDelete prepares the WorkspaceSettingsDelete request.
func (c WorkspaceSettingsClient) preparerForWorkspaceSettingsDelete(ctx context.Context, id WorkspaceSettingId) (*http.Request, error) {
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

// responderForWorkspaceSettingsDelete handles the response to the WorkspaceSettingsDelete request. The method always
// closes the http.Response Body.
func (c WorkspaceSettingsClient) responderForWorkspaceSettingsDelete(resp *http.Response) (result WorkspaceSettingsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
