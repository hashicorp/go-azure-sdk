package contenttypecontentitem

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentItemDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type ContentItemDeleteOperationOptions struct {
	IfMatch *string
}

func DefaultContentItemDeleteOperationOptions() ContentItemDeleteOperationOptions {
	return ContentItemDeleteOperationOptions{}
}

func (o ContentItemDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o ContentItemDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// ContentItemDelete ...
func (c ContentTypeContentItemClient) ContentItemDelete(ctx context.Context, id ContentItemId, options ContentItemDeleteOperationOptions) (result ContentItemDeleteOperationResponse, err error) {
	req, err := c.preparerForContentItemDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentItemDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentItemDelete prepares the ContentItemDelete request.
func (c ContentTypeContentItemClient) preparerForContentItemDelete(ctx context.Context, id ContentItemId, options ContentItemDeleteOperationOptions) (*http.Request, error) {
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

// responderForContentItemDelete handles the response to the ContentItemDelete request. The method always
// closes the http.Response Body.
func (c ContentTypeContentItemClient) responderForContentItemDelete(resp *http.Response) (result ContentItemDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
