package backuppolicy

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupPolicy
}

// BackupPoliciesGet ...
func (c BackupPolicyClient) BackupPoliciesGet(ctx context.Context, id BackupPoliciesId) (result BackupPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForBackupPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupPoliciesGet prepares the BackupPoliciesGet request.
func (c BackupPolicyClient) preparerForBackupPoliciesGet(ctx context.Context, id BackupPoliciesId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupPoliciesGet handles the response to the BackupPoliciesGet request. The method always
// closes the http.Response Body.
func (c BackupPolicyClient) responderForBackupPoliciesGet(resp *http.Response) (result BackupPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
