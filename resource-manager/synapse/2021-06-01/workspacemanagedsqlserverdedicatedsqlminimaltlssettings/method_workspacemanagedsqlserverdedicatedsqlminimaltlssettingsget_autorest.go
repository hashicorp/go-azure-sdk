package workspacemanagedsqlserverdedicatedsqlminimaltlssettings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DedicatedSQLminimalTlsSettings
}

// WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet ...
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet(ctx context.Context, id DedicatedSQLminimalTlsSettingId) (result WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet prepares the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet request.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet(ctx context.Context, id DedicatedSQLminimalTlsSettingId) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet handles the response to the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) responderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGet(resp *http.Response) (result WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
