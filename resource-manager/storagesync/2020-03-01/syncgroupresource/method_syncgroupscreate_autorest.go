package syncgroupresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupsCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SyncGroup
}

// SyncGroupsCreate ...
func (c SyncGroupResourceClient) SyncGroupsCreate(ctx context.Context, id SyncGroupId, input SyncGroupCreateParameters) (result SyncGroupsCreateOperationResponse, err error) {
	req, err := c.preparerForSyncGroupsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncGroupsCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncGroupsCreate prepares the SyncGroupsCreate request.
func (c SyncGroupResourceClient) preparerForSyncGroupsCreate(ctx context.Context, id SyncGroupId, input SyncGroupCreateParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncGroupsCreate handles the response to the SyncGroupsCreate request. The method always
// closes the http.Response Body.
func (c SyncGroupResourceClient) responderForSyncGroupsCreate(resp *http.Response) (result SyncGroupsCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
