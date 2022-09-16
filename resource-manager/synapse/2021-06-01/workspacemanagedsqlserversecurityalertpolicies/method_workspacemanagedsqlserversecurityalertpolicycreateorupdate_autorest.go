package workspacemanagedsqlserversecurityalertpolicies

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

type WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate ...
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate(ctx context.Context, id WorkspaceId, input ServerSecurityAlertPolicy) (result WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversecurityalertpolicies.WorkspaceManagedSqlServerSecurityAlertPoliciesClient", "WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversecurityalertpolicies.WorkspaceManagedSqlServerSecurityAlertPoliciesClient", "WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdateThenPoll performs WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate then polls until it's completed
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdateThenPoll(ctx context.Context, id WorkspaceId, input ServerSecurityAlertPolicy) error {
	result, err := c.WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForWorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate prepares the WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate request.
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) preparerForWorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate(ctx context.Context, id WorkspaceId, input ServerSecurityAlertPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/securityAlertPolicies/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate sends the WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceManagedSqlServerSecurityAlertPoliciesClient) senderForWorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate(ctx context.Context, req *http.Request) (future WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
