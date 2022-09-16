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

type WorkspaceAadAdminsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceAadAdminsCreateOrUpdate ...
func (c WorkspacesClient) WorkspaceAadAdminsCreateOrUpdate(ctx context.Context, id WorkspaceId, input WorkspaceAadAdminInfo) (result WorkspaceAadAdminsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceAadAdminsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceAadAdminsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceAadAdminsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceAadAdminsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceAadAdminsCreateOrUpdateThenPoll performs WorkspaceAadAdminsCreateOrUpdate then polls until it's completed
func (c WorkspacesClient) WorkspaceAadAdminsCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input WorkspaceAadAdminInfo) error {
	result, err := c.WorkspaceAadAdminsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceAadAdminsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceAadAdminsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceAadAdminsCreateOrUpdate prepares the WorkspaceAadAdminsCreateOrUpdate request.
func (c WorkspacesClient) preparerForWorkspaceAadAdminsCreateOrUpdate(ctx context.Context, id WorkspaceId, input WorkspaceAadAdminInfo) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/administrators/activeDirectory", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceAadAdminsCreateOrUpdate sends the WorkspaceAadAdminsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspacesClient) senderForWorkspaceAadAdminsCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceAadAdminsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
