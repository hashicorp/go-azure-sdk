package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetConfigurationSnapshotSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// GetConfigurationSnapshotSlot ...
func (c WebAppsClient) GetConfigurationSnapshotSlot(ctx context.Context, id WebSnapshotId) (result GetConfigurationSnapshotSlotOperationResponse, err error) {
	req, err := c.preparerForGetConfigurationSnapshotSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfigurationSnapshotSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfigurationSnapshotSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetConfigurationSnapshotSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfigurationSnapshotSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetConfigurationSnapshotSlot prepares the GetConfigurationSnapshotSlot request.
func (c WebAppsClient) preparerForGetConfigurationSnapshotSlot(ctx context.Context, id WebSnapshotId) (*http.Request, error) {
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

// responderForGetConfigurationSnapshotSlot handles the response to the GetConfigurationSnapshotSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetConfigurationSnapshotSlot(resp *http.Response) (result GetConfigurationSnapshotSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
