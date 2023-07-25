package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBackupStatusSecretsOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupItem
}

// ListBackupStatusSecrets ...
func (c WebAppsClient) ListBackupStatusSecrets(ctx context.Context, id BackupId, input BackupRequest) (result ListBackupStatusSecretsOperationResponse, err error) {
	req, err := c.preparerForListBackupStatusSecrets(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackupStatusSecrets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackupStatusSecrets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListBackupStatusSecrets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackupStatusSecrets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListBackupStatusSecrets prepares the ListBackupStatusSecrets request.
func (c WebAppsClient) preparerForListBackupStatusSecrets(ctx context.Context, id BackupId, input BackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/list", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListBackupStatusSecrets handles the response to the ListBackupStatusSecrets request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListBackupStatusSecrets(resp *http.Response) (result ListBackupStatusSecretsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
