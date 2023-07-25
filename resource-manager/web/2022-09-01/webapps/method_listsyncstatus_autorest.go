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

type ListSyncStatusOperationResponse struct {
	HttpResponse *http.Response
}

// ListSyncStatus ...
func (c WebAppsClient) ListSyncStatus(ctx context.Context, id SiteId) (result ListSyncStatusOperationResponse, err error) {
	req, err := c.preparerForListSyncStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListSyncStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListSyncStatus prepares the ListSyncStatus request.
func (c WebAppsClient) preparerForListSyncStatus(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/host/default/listsyncstatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSyncStatus handles the response to the ListSyncStatus request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSyncStatus(resp *http.Response) (result ListSyncStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
