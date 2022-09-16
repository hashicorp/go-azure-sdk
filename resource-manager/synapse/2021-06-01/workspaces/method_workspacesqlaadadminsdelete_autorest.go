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

type WorkspaceSqlAadAdminsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceSqlAadAdminsDelete ...
func (c WorkspacesClient) WorkspaceSqlAadAdminsDelete(ctx context.Context, id WorkspaceId) (result WorkspaceSqlAadAdminsDeleteOperationResponse, err error) {
	req, err := c.preparerForWorkspaceSqlAadAdminsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceSqlAadAdminsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaces.WorkspacesClient", "WorkspaceSqlAadAdminsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceSqlAadAdminsDeleteThenPoll performs WorkspaceSqlAadAdminsDelete then polls until it's completed
func (c WorkspacesClient) WorkspaceSqlAadAdminsDeleteThenPoll(ctx context.Context, id WorkspaceId) error {
	result, err := c.WorkspaceSqlAadAdminsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing WorkspaceSqlAadAdminsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceSqlAadAdminsDelete: %+v", err)
	}

	return nil
}

// preparerForWorkspaceSqlAadAdminsDelete prepares the WorkspaceSqlAadAdminsDelete request.
func (c WorkspacesClient) preparerForWorkspaceSqlAadAdminsDelete(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sqlAdministrators/activeDirectory", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceSqlAadAdminsDelete sends the WorkspaceSqlAadAdminsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspacesClient) senderForWorkspaceSqlAadAdminsDelete(ctx context.Context, req *http.Request) (future WorkspaceSqlAadAdminsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
