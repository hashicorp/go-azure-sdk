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

type WorkspaceSqlAadAdminsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceSqlAadAdminsCreateOrUpdate ...
func (c WorkspacesClient) WorkspaceSqlAadAdminsCreateOrUpdate(ctx context.Context, id WorkspaceId, input WorkspaceAadAdminInfo) (result WorkspaceSqlAadAdminsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSqlAadAdminsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceSqlAadAdminsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceSqlAadAdminsCreateOrUpdateThenPoll performs WorkspaceSqlAadAdminsCreateOrUpdate then polls until it's completed
func (c WorkspacesClient) WorkspaceSqlAadAdminsCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input WorkspaceAadAdminInfo) error {
	result, err := c.WorkspaceSqlAadAdminsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceSqlAadAdminsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceSqlAadAdminsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceSqlAadAdminsCreateOrUpdate prepares the WorkspaceSqlAadAdminsCreateOrUpdate request.
func (c WorkspacesClient) preparerForWorkspaceSqlAadAdminsCreateOrUpdate(ctx context.Context, id WorkspaceId, input WorkspaceAadAdminInfo) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sqlAdministrators/activeDirectory", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceSqlAadAdminsCreateOrUpdate sends the WorkspaceSqlAadAdminsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspacesClient) senderForWorkspaceSqlAadAdminsCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceSqlAadAdminsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
