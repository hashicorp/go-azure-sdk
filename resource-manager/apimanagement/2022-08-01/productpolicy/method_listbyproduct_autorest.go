package productpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByProductOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicyCollection
}

// ListByProduct ...
func (c ProductPolicyClient) ListByProduct(ctx context.Context, id ProductId) (result ListByProductOperationResponse, err error) {
	req, err := c.preparerForListByProduct(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "productpolicy.ProductPolicyClient", "ListByProduct", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "productpolicy.ProductPolicyClient", "ListByProduct", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByProduct(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "productpolicy.ProductPolicyClient", "ListByProduct", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByProduct prepares the ListByProduct request.
func (c ProductPolicyClient) preparerForListByProduct(ctx context.Context, id ProductId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/policies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByProduct handles the response to the ListByProduct request. The method always
// closes the http.Response Body.
func (c ProductPolicyClient) responderForListByProduct(resp *http.Response) (result ListByProductOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
