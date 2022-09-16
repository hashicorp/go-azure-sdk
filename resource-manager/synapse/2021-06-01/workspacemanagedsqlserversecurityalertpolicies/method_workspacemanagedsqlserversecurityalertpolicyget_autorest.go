package workspacemanagedsqlserversecurityalertpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceManagedSqlServerSecurityAlertPolicyGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ServerSecurityAlertPolicy
}

// WorkspaceManagedSqlServerSecurityAlertPolicyGet ...
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) WorkspaceManagedSqlServerSecurityAlertPolicyGet(ctx context.Context, id WorkspaceId) (result WorkspaceManagedSqlServerSecurityAlertPolicyGetOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerSecurityAlertPolicyGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversecurityalertpolicies.WorkspaceManagedSqlServerSecurityAlertPoliciesClient", "WorkspaceManagedSqlServerSecurityAlertPolicyGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversecurityalertpolicies.WorkspaceManagedSqlServerSecurityAlertPoliciesClient", "WorkspaceManagedSqlServerSecurityAlertPolicyGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspaceManagedSqlServerSecurityAlertPolicyGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversecurityalertpolicies.WorkspaceManagedSqlServerSecurityAlertPoliciesClient", "WorkspaceManagedSqlServerSecurityAlertPolicyGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspaceManagedSqlServerSecurityAlertPolicyGet prepares the WorkspaceManagedSqlServerSecurityAlertPolicyGet request.
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) preparerForWorkspaceManagedSqlServerSecurityAlertPolicyGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/securityAlertPolicies/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspaceManagedSqlServerSecurityAlertPolicyGet handles the response to the WorkspaceManagedSqlServerSecurityAlertPolicyGet request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) responderForWorkspaceManagedSqlServerSecurityAlertPolicyGet(resp *http.Response) (result WorkspaceManagedSqlServerSecurityAlertPolicyGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
