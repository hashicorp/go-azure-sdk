package apioperationpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByOperationOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicyCollection
}

// ListByOperation ...
func (c ApiOperationPolicyClient) ListByOperation(ctx context.Context, id OperationId) (result ListByOperationOperationResponse, err error) {
	req, err := c.preparerForListByOperation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationpolicy.ApiOperationPolicyClient", "ListByOperation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationpolicy.ApiOperationPolicyClient", "ListByOperation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByOperation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationpolicy.ApiOperationPolicyClient", "ListByOperation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByOperation prepares the ListByOperation request.
func (c ApiOperationPolicyClient) preparerForListByOperation(ctx context.Context, id OperationId) (*http.Request, error) {
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

// responderForListByOperation handles the response to the ListByOperation request. The method always
// closes the http.Response Body.
func (c ApiOperationPolicyClient) responderForListByOperation(resp *http.Response) (result ListByOperationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
