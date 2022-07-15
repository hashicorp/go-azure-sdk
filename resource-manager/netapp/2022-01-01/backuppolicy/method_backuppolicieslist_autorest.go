package backuppolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupPoliciesList
}

// BackupPoliciesList ...
func (c BackupPolicyClient) BackupPoliciesList(ctx context.Context, id NetAppAccountId) (result BackupPoliciesListOperationResponse, err error) {
	req, err := c.preparerForBackupPoliciesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupPoliciesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backuppolicy.BackupPolicyClient", "BackupPoliciesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupPoliciesList prepares the BackupPoliciesList request.
func (c BackupPolicyClient) preparerForBackupPoliciesList(ctx context.Context, id NetAppAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupPolicies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupPoliciesList handles the response to the BackupPoliciesList request. The method always
// closes the http.Response Body.
func (c BackupPolicyClient) responderForBackupPoliciesList(resp *http.Response) (result BackupPoliciesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
