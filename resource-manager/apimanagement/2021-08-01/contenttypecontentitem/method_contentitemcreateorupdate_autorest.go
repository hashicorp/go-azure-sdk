package contenttypecontentitem

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentItemCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContentItemContract
}

type ContentItemCreateOrUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultContentItemCreateOrUpdateOperationOptions() ContentItemCreateOrUpdateOperationOptions {
	return ContentItemCreateOrUpdateOperationOptions{}
}

func (o ContentItemCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o ContentItemCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// ContentItemCreateOrUpdate ...
func (c ContentTypeContentItemClient) ContentItemCreateOrUpdate(ctx context.Context, id ContentItemId, input ContentItemContract, options ContentItemCreateOrUpdateOperationOptions) (result ContentItemCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForContentItemCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentItemCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contenttypecontentitem.ContentTypeContentItemClient", "ContentItemCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentItemCreateOrUpdate prepares the ContentItemCreateOrUpdate request.
func (c ContentTypeContentItemClient) preparerForContentItemCreateOrUpdate(ctx context.Context, id ContentItemId, input ContentItemContract, options ContentItemCreateOrUpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentItemCreateOrUpdate handles the response to the ContentItemCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c ContentTypeContentItemClient) responderForContentItemCreateOrUpdate(resp *http.Response) (result ContentItemCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
