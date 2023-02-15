package apiversionsets

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiVersionSetDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type ApiVersionSetDeleteOperationOptions struct {
	IfMatch *string
}

func DefaultApiVersionSetDeleteOperationOptions() ApiVersionSetDeleteOperationOptions {
	return ApiVersionSetDeleteOperationOptions{}
}

func (o ApiVersionSetDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o ApiVersionSetDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// ApiVersionSetDelete ...
func (c ApiVersionSetsClient) ApiVersionSetDelete(ctx context.Context, id ApiVersionSetId, options ApiVersionSetDeleteOperationOptions) (result ApiVersionSetDeleteOperationResponse, err error) {
	req, err := c.preparerForApiVersionSetDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apiversionsets.ApiVersionSetsClient", "ApiVersionSetDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apiversionsets.ApiVersionSetsClient", "ApiVersionSetDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForApiVersionSetDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apiversionsets.ApiVersionSetsClient", "ApiVersionSetDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForApiVersionSetDelete prepares the ApiVersionSetDelete request.
func (c ApiVersionSetsClient) preparerForApiVersionSetDelete(ctx context.Context, id ApiVersionSetId, options ApiVersionSetDeleteOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForApiVersionSetDelete handles the response to the ApiVersionSetDelete request. The method always
// closes the http.Response Body.
func (c ApiVersionSetsClient) responderForApiVersionSetDelete(resp *http.Response) (result ApiVersionSetDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
