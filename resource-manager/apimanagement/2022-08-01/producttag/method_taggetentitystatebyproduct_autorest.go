package producttag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagGetEntityStateByProductOperationResponse struct {
	HttpResponse *http.Response
}

// TagGetEntityStateByProduct ...
func (c ProductTagClient) TagGetEntityStateByProduct(ctx context.Context, id TagId) (result TagGetEntityStateByProductOperationResponse, err error) {
	req, err := c.preparerForTagGetEntityStateByProduct(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagGetEntityStateByProduct", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagGetEntityStateByProduct", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagGetEntityStateByProduct(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagGetEntityStateByProduct", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagGetEntityStateByProduct prepares the TagGetEntityStateByProduct request.
func (c ProductTagClient) preparerForTagGetEntityStateByProduct(ctx context.Context, id TagId) (*http.Request, error) {
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

// responderForTagGetEntityStateByProduct handles the response to the TagGetEntityStateByProduct request. The method always
// closes the http.Response Body.
func (c ProductTagClient) responderForTagGetEntityStateByProduct(resp *http.Response) (result TagGetEntityStateByProductOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
