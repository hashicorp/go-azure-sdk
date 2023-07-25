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

type RecoverSiteConfigurationSnapshotOperationResponse struct {
	HttpResponse *http.Response
}

// RecoverSiteConfigurationSnapshot ...
func (c WebAppsClient) RecoverSiteConfigurationSnapshot(ctx context.Context, id SnapshotId) (result RecoverSiteConfigurationSnapshotOperationResponse, err error) {
	req, err := c.preparerForRecoverSiteConfigurationSnapshot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RecoverSiteConfigurationSnapshot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RecoverSiteConfigurationSnapshot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecoverSiteConfigurationSnapshot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "RecoverSiteConfigurationSnapshot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecoverSiteConfigurationSnapshot prepares the RecoverSiteConfigurationSnapshot request.
func (c WebAppsClient) preparerForRecoverSiteConfigurationSnapshot(ctx context.Context, id SnapshotId) (*http.Request, error) {
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

// responderForRecoverSiteConfigurationSnapshot handles the response to the RecoverSiteConfigurationSnapshot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForRecoverSiteConfigurationSnapshot(resp *http.Response) (result RecoverSiteConfigurationSnapshotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
