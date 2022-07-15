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

type BackupPoliciesCreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BackupPoliciesCreate ...
func (c BackupPolicyClient) BackupPoliciesCreate(ctx context.Context, id BackupPoliciesId, input BackupPolicy) (result BackupPoliciesCreateOperationResponse, err error) {
	req, err := c.preparerForBackupPoliciesCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesCreate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBackupPoliciesCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesCreate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BackupPoliciesCreateThenPoll performs BackupPoliciesCreate then polls until it's completed
func (c BackupPolicyClient) BackupPoliciesCreateThenPoll(ctx context.Context, id BackupPoliciesId, input BackupPolicy) error {
	result, err := c.BackupPoliciesCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BackupPoliciesCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BackupPoliciesCreate: %+v", err)
	}

	return nil
}

// preparerForBackupPoliciesCreate prepares the BackupPoliciesCreate request.
func (c BackupPolicyClient) preparerForBackupPoliciesCreate(ctx context.Context, id BackupPoliciesId, input BackupPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForBackupPoliciesCreate sends the BackupPoliciesCreate request. The method will close the
// http.Response Body if it receives an error.
func (c BackupPolicyClient) senderForBackupPoliciesCreate(ctx context.Context, req *http.Request) (future BackupPoliciesCreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
