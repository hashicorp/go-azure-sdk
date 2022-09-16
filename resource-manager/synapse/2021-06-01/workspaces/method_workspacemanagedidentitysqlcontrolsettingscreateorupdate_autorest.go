package workspaces

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

type WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate ...
func (c WorkspacesClient) WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate(ctx context.Context, id WorkspaceId, input ManagedIdentitySqlControlSettingsModel) (result WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdateThenPoll performs WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate then polls until it's completed
func (c WorkspacesClient) WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input ManagedIdentitySqlControlSettingsModel) error {
	result, err := c.WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate prepares the WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate request.
func (c WorkspacesClient) preparerForWorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate(ctx context.Context, id WorkspaceId, input ManagedIdentitySqlControlSettingsModel) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/managedIdentitySqlControlSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate sends the WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspacesClient) senderForWorkspaceManagedIdentitySqlControlSettingsCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceManagedIdentitySqlControlSettingsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
