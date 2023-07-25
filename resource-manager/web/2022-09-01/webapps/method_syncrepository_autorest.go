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

type SyncRepositoryOperationResponse struct {
	HttpResponse *http.Response
}

// SyncRepository ...
func (c WebAppsClient) SyncRepository(ctx context.Context, id SiteId) (result SyncRepositoryOperationResponse, err error) {
	req, err := c.preparerForSyncRepository(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncRepository", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncRepository", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncRepository(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "SyncRepository", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncRepository prepares the SyncRepository request.
func (c WebAppsClient) preparerForSyncRepository(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sync", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncRepository handles the response to the SyncRepository request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForSyncRepository(resp *http.Response) (result SyncRepositoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
