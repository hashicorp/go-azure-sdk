package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteBackupSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteBackupSlot ...
func (c WebAppsClient) DeleteBackupSlot(ctx context.Context, id SlotBackupId) (result DeleteBackupSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteBackupSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteBackupSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteBackupSlot prepares the DeleteBackupSlot request.
func (c WebAppsClient) preparerForDeleteBackupSlot(ctx context.Context, id SlotBackupId) (*http.Request, error) {
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

// responderForDeleteBackupSlot handles the response to the DeleteBackupSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteBackupSlot(resp *http.Response) (result DeleteBackupSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
