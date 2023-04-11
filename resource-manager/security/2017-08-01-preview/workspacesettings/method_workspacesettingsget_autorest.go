package workspacesettings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceSetting
}

// WorkspaceSettingsGet ...
func (c WorkspaceSettingsClient) WorkspaceSettingsGet(ctx context.Context, id WorkspaceSettingId) (result WorkspaceSettingsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSettingsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceSettingsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceSettingsGet prepares the WorkspaceSettingsGet request.
func (c WorkspaceSettingsClient) preparerForWorkspaceSettingsGet(ctx context.Context, id WorkspaceSettingId) (*http.Request, error) {
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

// responderForWorkspaceSettingsGet handles the response to the WorkspaceSettingsGet request. The method always
// closes the http.Response Body.
func (c WorkspaceSettingsClient) responderForWorkspaceSettingsGet(resp *http.Response) (result WorkspaceSettingsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
