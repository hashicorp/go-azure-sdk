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

type DeleteBackupConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteBackupConfigurationSlot ...
func (c WebAppsClient) DeleteBackupConfigurationSlot(ctx context.Context, id SlotId) (result DeleteBackupConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteBackupConfigurationSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteBackupConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteBackupConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteBackupConfigurationSlot prepares the DeleteBackupConfigurationSlot request.
func (c WebAppsClient) preparerForDeleteBackupConfigurationSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/backup", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteBackupConfigurationSlot handles the response to the DeleteBackupConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteBackupConfigurationSlot(resp *http.Response) (result DeleteBackupConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
