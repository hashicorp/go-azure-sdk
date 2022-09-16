package workspacemanagedsqlserverblobauditing

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

type WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, id WorkspaceId, input ServerBlobAuditingPolicy) (result WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdateThenPoll performs WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input ServerBlobAuditingPolicy) error {
	result, err := c.WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate prepares the WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate request.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, id WorkspaceId, input ServerBlobAuditingPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/auditingSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate sends the WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceManagedSqlServerBlobAuditingClient) senderForWorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
