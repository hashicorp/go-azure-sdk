package backupusagesummariescrr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupUsageSummariesCRRListOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupManagementUsage
}

type BackupUsageSummariesCRRListOperationOptions struct {
	Filter *string
}

func DefaultBackupUsageSummariesCRRListOperationOptions() BackupUsageSummariesCRRListOperationOptions {
	return BackupUsageSummariesCRRListOperationOptions{}
}

func (o BackupUsageSummariesCRRListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o BackupUsageSummariesCRRListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// BackupUsageSummariesCRRList ...
func (c BackupUsageSummariesCRRClient) BackupUsageSummariesCRRList(ctx context.Context, id VaultId, options BackupUsageSummariesCRRListOperationOptions) (result BackupUsageSummariesCRRListOperationResponse, err error) {
	req, err := c.preparerForBackupUsageSummariesCRRList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupusagesummariescrr.BackupUsageSummariesCRRClient", "BackupUsageSummariesCRRList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupusagesummariescrr.BackupUsageSummariesCRRClient", "BackupUsageSummariesCRRList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupUsageSummariesCRRList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backupusagesummariescrr.BackupUsageSummariesCRRClient", "BackupUsageSummariesCRRList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupUsageSummariesCRRList prepares the BackupUsageSummariesCRRList request.
func (c BackupUsageSummariesCRRClient) preparerForBackupUsageSummariesCRRList(ctx context.Context, id VaultId, options BackupUsageSummariesCRRListOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/backupUsageSummaries", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupUsageSummariesCRRList handles the response to the BackupUsageSummariesCRRList request. The method always
// closes the http.Response Body.
func (c BackupUsageSummariesCRRClient) responderForBackupUsageSummariesCRRList(resp *http.Response) (result BackupUsageSummariesCRRListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
