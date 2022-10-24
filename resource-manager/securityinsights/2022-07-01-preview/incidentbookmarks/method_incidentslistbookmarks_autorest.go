package incidentbookmarks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentsListBookmarksOperationResponse struct {
	HttpResponse *http.Response
	Model        *IncidentBookmarkList
}

// IncidentsListBookmarks ...
func (c IncidentBookmarksClient) IncidentsListBookmarks(ctx context.Context, id IncidentId) (result IncidentsListBookmarksOperationResponse, err error) {
	req, err := c.preparerForIncidentsListBookmarks(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentbookmarks.IncidentBookmarksClient", "IncidentsListBookmarks", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentbookmarks.IncidentBookmarksClient", "IncidentsListBookmarks", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIncidentsListBookmarks(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentbookmarks.IncidentBookmarksClient", "IncidentsListBookmarks", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIncidentsListBookmarks prepares the IncidentsListBookmarks request.
func (c IncidentBookmarksClient) preparerForIncidentsListBookmarks(ctx context.Context, id IncidentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/bookmarks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIncidentsListBookmarks handles the response to the IncidentsListBookmarks request. The method always
// closes the http.Response Body.
func (c IncidentBookmarksClient) responderForIncidentsListBookmarks(resp *http.Response) (result IncidentsListBookmarksOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
