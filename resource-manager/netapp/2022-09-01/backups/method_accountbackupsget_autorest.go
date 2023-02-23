package backups

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountBackupsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Backup
}

// AccountBackupsGet ...
func (c BackupsClient) AccountBackupsGet(ctx context.Context, id AccountBackupId) (result AccountBackupsGetOperationResponse, err error) {
	req, err := c.preparerForAccountBackupsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "AccountBackupsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "AccountBackupsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAccountBackupsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.BackupsClient", "AccountBackupsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAccountBackupsGet prepares the AccountBackupsGet request.
func (c BackupsClient) preparerForAccountBackupsGet(ctx context.Context, id AccountBackupId) (*http.Request, error) {
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

// responderForAccountBackupsGet handles the response to the AccountBackupsGet request. The method always
// closes the http.Response Body.
func (c BackupsClient) responderForAccountBackupsGet(resp *http.Response) (result AccountBackupsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
