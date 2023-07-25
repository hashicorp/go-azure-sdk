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

type ListBackupStatusSecretsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupItem
}

// ListBackupStatusSecretsSlot ...
func (c WebAppsClient) ListBackupStatusSecretsSlot(ctx context.Context, id SlotBackupId, input BackupRequest) (result ListBackupStatusSecretsSlotOperationResponse, err error) {
	req, err := c.preparerForListBackupStatusSecretsSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackupStatusSecretsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackupStatusSecretsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListBackupStatusSecretsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBackupStatusSecretsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListBackupStatusSecretsSlot prepares the ListBackupStatusSecretsSlot request.
func (c WebAppsClient) preparerForListBackupStatusSecretsSlot(ctx context.Context, id SlotBackupId, input BackupRequest) (*http.Request, error) {
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

// responderForListBackupStatusSecretsSlot handles the response to the ListBackupStatusSecretsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListBackupStatusSecretsSlot(resp *http.Response) (result ListBackupStatusSecretsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
