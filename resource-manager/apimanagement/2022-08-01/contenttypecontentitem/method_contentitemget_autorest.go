package contenttypecontentitem

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentItemGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContentItemContract
}

// ContentItemGet ...
func (c ContentTypeContentItemClient) ContentItemGet(ctx context.Context, id ContentItemId) (result ContentItemGetOperationResponse, err error) {
	req, err := c.preparerForContentItemGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentItemGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentItemGet prepares the ContentItemGet request.
func (c ContentTypeContentItemClient) preparerForContentItemGet(ctx context.Context, id ContentItemId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentItemGet handles the response to the ContentItemGet request. The method always
// closes the http.Response Body.
func (c ContentTypeContentItemClient) responderForContentItemGet(resp *http.Response) (result ContentItemGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
