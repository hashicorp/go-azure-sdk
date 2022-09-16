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

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, id WorkspaceId, input ExtendedServerBlobAuditingPolicy) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateThenPoll performs WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input ExtendedServerBlobAuditingPolicy) error {
	result, err := c.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate prepares the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate request.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, id WorkspaceId, input ExtendedServerBlobAuditingPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extendedAuditingSettings/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate sends the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceManagedSqlServerBlobAuditingClient) senderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
