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

type RecoverSiteConfigurationSnapshotSlotOperationResponse struct {
	HttpResponse *http.Response
}

// RecoverSiteConfigurationSnapshotSlot ...
func (c WebAppsClient) RecoverSiteConfigurationSnapshotSlot(ctx context.Context, id WebSnapshotId) (result RecoverSiteConfigurationSnapshotSlotOperationResponse, err error) {
	req, err := c.preparerForRecoverSiteConfigurationSnapshotSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RecoverSiteConfigurationSnapshotSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RecoverSiteConfigurationSnapshotSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecoverSiteConfigurationSnapshotSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RecoverSiteConfigurationSnapshotSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecoverSiteConfigurationSnapshotSlot prepares the RecoverSiteConfigurationSnapshotSlot request.
func (c WebAppsClient) preparerForRecoverSiteConfigurationSnapshotSlot(ctx context.Context, id WebSnapshotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recover", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRecoverSiteConfigurationSnapshotSlot handles the response to the RecoverSiteConfigurationSnapshotSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForRecoverSiteConfigurationSnapshotSlot(resp *http.Response) (result RecoverSiteConfigurationSnapshotSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
