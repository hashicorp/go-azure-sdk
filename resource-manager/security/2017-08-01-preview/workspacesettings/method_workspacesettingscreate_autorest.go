package workspacesettings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkspaceSetting
}

// WorkspaceSettingsCreate ...
func (c WorkspaceSettingsClient) WorkspaceSettingsCreate(ctx context.Context, id WorkspaceSettingId, input WorkspaceSetting) (result WorkspaceSettingsCreateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSettingsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceSettingsCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacesettings.WorkspaceSettingsClient", "WorkspaceSettingsCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceSettingsCreate prepares the WorkspaceSettingsCreate request.
func (c WorkspaceSettingsClient) preparerForWorkspaceSettingsCreate(ctx context.Context, id WorkspaceSettingId, input WorkspaceSetting) (*http.Request, error) {
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

// responderForWorkspaceSettingsCreate handles the response to the WorkspaceSettingsCreate request. The method always
// closes the http.Response Body.
func (c WorkspaceSettingsClient) responderForWorkspaceSettingsCreate(resp *http.Response) (result WorkspaceSettingsCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
