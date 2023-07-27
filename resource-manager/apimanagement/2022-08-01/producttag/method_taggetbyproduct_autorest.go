package producttag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagGetByProductOperationResponse struct {
	HttpResponse *http.Response
	Model        *TagContract
}

// TagGetByProduct ...
func (c ProductTagClient) TagGetByProduct(ctx context.Context, id ProductTagId) (result TagGetByProductOperationResponse, err error) {
	req, err := c.preparerForTagGetByProduct(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagGetByProduct", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagGetByProduct", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagGetByProduct(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagGetByProduct", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagGetByProduct prepares the TagGetByProduct request.
func (c ProductTagClient) preparerForTagGetByProduct(ctx context.Context, id ProductTagId) (*http.Request, error) {
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

// responderForTagGetByProduct handles the response to the TagGetByProduct request. The method always
// closes the http.Response Body.
func (c ProductTagClient) responderForTagGetByProduct(resp *http.Response) (result TagGetByProductOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
