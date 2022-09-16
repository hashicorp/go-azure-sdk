package workspacemanagedsqlserverdedicatedsqlminimaltlssettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate ...
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate(ctx context.Context, id DedicatedSQLminimalTlsSettingId, input DedicatedSQLminimalTlsSettings) (result WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient", "WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdateThenPoll performs WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdateThenPoll(ctx context.Context, id DedicatedSQLminimalTlsSettingId, input DedicatedSQLminimalTlsSettings) error {
	result, err := c.WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate prepares the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate request.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) preparerForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate(ctx context.Context, id DedicatedSQLminimalTlsSettingId, input DedicatedSQLminimalTlsSettings) (*http.Request, error) {
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

// senderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate sends the WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient) senderForWorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdate(ctx context.Context, req *http.Request) (future WorkspaceManagedSqlServerDedicatedSQLMinimalTlsSettingsUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
