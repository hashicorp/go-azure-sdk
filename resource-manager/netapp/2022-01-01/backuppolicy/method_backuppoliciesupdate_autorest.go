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

type BackupPoliciesUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BackupPoliciesUpdate ...
func (c BackupPolicyClient) BackupPoliciesUpdate(ctx context.Context, id BackupPolicyId, input BackupPolicyPatch) (result BackupPoliciesUpdateOperationResponse, err error) {
	req, err := c.preparerForBackupPoliciesUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBackupPoliciesUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BackupPoliciesUpdateThenPoll performs BackupPoliciesUpdate then polls until it's completed
func (c BackupPolicyClient) BackupPoliciesUpdateThenPoll(ctx context.Context, id BackupPolicyId, input BackupPolicyPatch) error {
	result, err := c.BackupPoliciesUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BackupPoliciesUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BackupPoliciesUpdate: %+v", err)
	}

	return nil
}

// preparerForBackupPoliciesUpdate prepares the BackupPoliciesUpdate request.
func (c BackupPolicyClient) preparerForBackupPoliciesUpdate(ctx context.Context, id BackupPolicyId, input BackupPolicyPatch) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBackupPoliciesUpdate sends the BackupPoliciesUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c BackupPolicyClient) senderForBackupPoliciesUpdate(ctx context.Context, req *http.Request) (future BackupPoliciesUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
