package backuppolicy

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

type BackupPoliciesDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BackupPoliciesDelete ...
func (c BackupPolicyClient) BackupPoliciesDelete(ctx context.Context, id BackupPolicyId) (result BackupPoliciesDeleteOperationResponse, err error) {
	req, err := c.preparerForBackupPoliciesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBackupPoliciesDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BackupPoliciesDeleteThenPoll performs BackupPoliciesDelete then polls until it's completed
func (c BackupPolicyClient) BackupPoliciesDeleteThenPoll(ctx context.Context, id BackupPolicyId) error {
	result, err := c.BackupPoliciesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing BackupPoliciesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BackupPoliciesDelete: %+v", err)
	}

	return nil
}

// preparerForBackupPoliciesDelete prepares the BackupPoliciesDelete request.
func (c BackupPolicyClient) preparerForBackupPoliciesDelete(ctx context.Context, id BackupPolicyId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBackupPoliciesDelete sends the BackupPoliciesDelete request. The method will close the
// http.Response Body if it receives an error.
func (c BackupPolicyClient) senderForBackupPoliciesDelete(ctx context.Context, req *http.Request) (future BackupPoliciesDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
