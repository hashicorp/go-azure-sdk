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

type SyncFunctionsOperationResponse struct {
	HttpResponse *http.Response
}

// SyncFunctions ...
func (c WebAppsClient) SyncFunctions(ctx context.Context, id SiteId) (result SyncFunctionsOperationResponse, err error) {
	req, err := c.preparerForSyncFunctions(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctions", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctions", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncFunctions(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncFunctions", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncFunctions prepares the SyncFunctions request.
func (c WebAppsClient) preparerForSyncFunctions(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/host/default/sync", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncFunctions handles the response to the SyncFunctions request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForSyncFunctions(resp *http.Response) (result SyncFunctionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
