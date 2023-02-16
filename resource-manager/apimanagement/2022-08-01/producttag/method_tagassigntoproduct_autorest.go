package producttag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagAssignToProductOperationResponse struct {
	HttpResponse *http.Response
	Model        *TagContract
}

// TagAssignToProduct ...
func (c ProductTagClient) TagAssignToProduct(ctx context.Context, id TagId) (result TagAssignToProductOperationResponse, err error) {
	req, err := c.preparerForTagAssignToProduct(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagAssignToProduct", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagAssignToProduct", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagAssignToProduct(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagAssignToProduct", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagAssignToProduct prepares the TagAssignToProduct request.
func (c ProductTagClient) preparerForTagAssignToProduct(ctx context.Context, id TagId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTagAssignToProduct handles the response to the TagAssignToProduct request. The method always
// closes the http.Response Body.
func (c ProductTagClient) responderForTagAssignToProduct(resp *http.Response) (result TagAssignToProductOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
