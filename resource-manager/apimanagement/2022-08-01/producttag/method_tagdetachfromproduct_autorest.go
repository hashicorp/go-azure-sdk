package producttag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagDetachFromProductOperationResponse struct {
	HttpResponse *http.Response
}

// TagDetachFromProduct ...
func (c ProductTagClient) TagDetachFromProduct(ctx context.Context, id TagId) (result TagDetachFromProductOperationResponse, err error) {
	req, err := c.preparerForTagDetachFromProduct(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagDetachFromProduct", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagDetachFromProduct", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagDetachFromProduct(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagDetachFromProduct", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagDetachFromProduct prepares the TagDetachFromProduct request.
func (c ProductTagClient) preparerForTagDetachFromProduct(ctx context.Context, id TagId) (*http.Request, error) {
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

// responderForTagDetachFromProduct handles the response to the TagDetachFromProduct request. The method always
// closes the http.Response Body.
func (c ProductTagClient) responderForTagDetachFromProduct(resp *http.Response) (result TagDetachFromProductOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
