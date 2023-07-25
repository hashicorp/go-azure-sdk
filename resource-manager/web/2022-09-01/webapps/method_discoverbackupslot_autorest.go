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

type DiscoverBackupSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestoreRequest
}

// DiscoverBackupSlot ...
func (c WebAppsClient) DiscoverBackupSlot(ctx context.Context, id SlotId, input RestoreRequest) (result DiscoverBackupSlotOperationResponse, err error) {
	req, err := c.preparerForDiscoverBackupSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DiscoverBackupSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DiscoverBackupSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDiscoverBackupSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DiscoverBackupSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDiscoverBackupSlot prepares the DiscoverBackupSlot request.
func (c WebAppsClient) preparerForDiscoverBackupSlot(ctx context.Context, id SlotId, input RestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/discoverbackup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDiscoverBackupSlot handles the response to the DiscoverBackupSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDiscoverBackupSlot(resp *http.Response) (result DiscoverBackupSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
