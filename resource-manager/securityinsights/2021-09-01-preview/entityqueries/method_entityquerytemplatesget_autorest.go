package entityqueries

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityQueryTemplatesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EntityQueryTemplate
}

// EntityQueryTemplatesGet ...
func (c EntityQueriesClient) EntityQueryTemplatesGet(ctx context.Context, id EntityQueryTemplateId) (result EntityQueryTemplatesGetOperationResponse, err error) {
	req, err := c.preparerForEntityQueryTemplatesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEntityQueryTemplatesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityqueries.EntityQueriesClient", "EntityQueryTemplatesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEntityQueryTemplatesGet prepares the EntityQueryTemplatesGet request.
func (c EntityQueriesClient) preparerForEntityQueryTemplatesGet(ctx context.Context, id EntityQueryTemplateId) (*http.Request, error) {
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

// responderForEntityQueryTemplatesGet handles the response to the EntityQueryTemplatesGet request. The method always
// closes the http.Response Body.
func (c EntityQueriesClient) responderForEntityQueryTemplatesGet(resp *http.Response) (result EntityQueryTemplatesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
