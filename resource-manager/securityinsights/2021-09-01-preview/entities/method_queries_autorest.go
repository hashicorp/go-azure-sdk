package entities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueriesOperationResponse struct {
	HttpResponse *http.Response
	Model        *GetQueriesResponse
}

type QueriesOperationOptions struct {
	Kind *EntityItemQueryKind
}

func DefaultQueriesOperationOptions() QueriesOperationOptions {
	return QueriesOperationOptions{}
}

func (o QueriesOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o QueriesOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Kind != nil {
		out["kind"] = *o.Kind
	}

	return out
}

// Queries ...
func (c EntitiesClient) Queries(ctx context.Context, id EntitiesId, options QueriesOperationOptions) (result QueriesOperationResponse, err error) {
	req, err := c.preparerForQueries(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "Queries", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "Queries", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForQueries(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "Queries", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForQueries prepares the Queries request.
func (c EntitiesClient) preparerForQueries(ctx context.Context, id EntitiesId, options QueriesOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/queries", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForQueries handles the response to the Queries request. The method always
// closes the http.Response Body.
func (c EntitiesClient) responderForQueries(resp *http.Response) (result QueriesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
