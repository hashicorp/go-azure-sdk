package bookmark

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpandOperationResponse struct {
	HttpResponse *http.Response
	Model        *BookmarkExpandResponse
}

// Expand ...
func (c BookmarkClient) Expand(ctx context.Context, id BookmarkId, input BookmarkExpandParameters) (result ExpandOperationResponse, err error) {
	req, err := c.preparerForExpand(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bookmark.BookmarkClient", "Expand", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bookmark.BookmarkClient", "Expand", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExpand(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bookmark.BookmarkClient", "Expand", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExpand prepares the Expand request.
func (c BookmarkClient) preparerForExpand(ctx context.Context, id BookmarkId, input BookmarkExpandParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/expand", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForExpand handles the response to the Expand request. The method always
// closes the http.Response Body.
func (c BookmarkClient) responderForExpand(resp *http.Response) (result ExpandOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
