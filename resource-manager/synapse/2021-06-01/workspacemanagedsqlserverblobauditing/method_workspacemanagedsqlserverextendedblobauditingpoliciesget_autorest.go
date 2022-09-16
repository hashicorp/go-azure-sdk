package workspacemanagedsqlserverblobauditing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExtendedServerBlobAuditingPolicy
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet prepares the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet request.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extendedAuditingSettings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet handles the response to the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerBlobAuditingClient) responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet(resp *http.Response) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
