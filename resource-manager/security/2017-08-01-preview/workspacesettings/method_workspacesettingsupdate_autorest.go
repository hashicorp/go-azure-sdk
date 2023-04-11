package workspacesettings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceSetting
}

// WorkspaceSettingsUpdate ...
func (c WorkspaceSettingsClient) WorkspaceSettingsUpdate(ctx context.Context, id WorkspaceSettingId, input WorkspaceSetting) (result WorkspaceSettingsUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSettingsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceSettingsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceSettingsUpdate prepares the WorkspaceSettingsUpdate request.
func (c WorkspaceSettingsClient) preparerForWorkspaceSettingsUpdate(ctx context.Context, id WorkspaceSettingId, input WorkspaceSetting) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceSettingsUpdate handles the response to the WorkspaceSettingsUpdate request. The method always
// closes the http.Response Body.
func (c WorkspaceSettingsClient) responderForWorkspaceSettingsUpdate(resp *http.Response) (result WorkspaceSettingsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
