package contenttypecontentitem

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentItemGetEntityTagOperationResponse struct {
	HttpResponse *http.Response
}

// ContentItemGetEntityTag ...
func (c ContentTypeContentItemClient) ContentItemGetEntityTag(ctx context.Context, id ContentItemId) (result ContentItemGetEntityTagOperationResponse, err error) {
	req, err := c.preparerForContentItemGetEntityTag(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemGetEntityTag", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemGetEntityTag", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentItemGetEntityTag(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemGetEntityTag", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentItemGetEntityTag prepares the ContentItemGetEntityTag request.
func (c ContentTypeContentItemClient) preparerForContentItemGetEntityTag(ctx context.Context, id ContentItemId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentItemGetEntityTag handles the response to the ContentItemGetEntityTag request. The method always
// closes the http.Response Body.
func (c ContentTypeContentItemClient) responderForContentItemGetEntityTag(resp *http.Response) (result ContentItemGetEntityTagOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
