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

type WorkspaceManagedSqlServerBlobAuditingPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ServerBlobAuditingPolicy
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesGet ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesGet(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerBlobAuditingPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedSqlServerBlobAuditingPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesGet prepares the WorkspaceManagedSqlServerBlobAuditingPoliciesGet request.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/auditingSettings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceManagedSqlServerBlobAuditingPoliciesGet handles the response to the WorkspaceManagedSqlServerBlobAuditingPoliciesGet request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerBlobAuditingClient) responderForWorkspaceManagedSqlServerBlobAuditingPoliciesGet(resp *http.Response) (result WorkspaceManagedSqlServerBlobAuditingPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
