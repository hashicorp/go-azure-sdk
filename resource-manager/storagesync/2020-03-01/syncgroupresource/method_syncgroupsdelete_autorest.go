package syncgroupresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// SyncGroupsDelete ...
func (c SyncGroupResourceClient) SyncGroupsDelete(ctx context.Context, id SyncGroupId) (result SyncGroupsDeleteOperationResponse, err error) {
	req, err := c.preparerForSyncGroupsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncGroupsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncGroupsDelete prepares the SyncGroupsDelete request.
func (c SyncGroupResourceClient) preparerForSyncGroupsDelete(ctx context.Context, id SyncGroupId) (*http.Request, error) {
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

// responderForSyncGroupsDelete handles the response to the SyncGroupsDelete request. The method always
// closes the http.Response Body.
func (c SyncGroupResourceClient) responderForSyncGroupsDelete(resp *http.Response) (result SyncGroupsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
