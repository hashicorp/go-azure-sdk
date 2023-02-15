package policyfragment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListReferencesOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceCollection
}

type ListReferencesOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListReferencesOperationOptions() ListReferencesOperationOptions {
	return ListReferencesOperationOptions{}
}

func (o ListReferencesOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListReferencesOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListReferences ...
func (c PolicyFragmentClient) ListReferences(ctx context.Context, id PolicyFragmentId, options ListReferencesOperationOptions) (result ListReferencesOperationResponse, err error) {
	req, err := c.preparerForListReferences(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyfragment.PolicyFragmentClient", "ListReferences", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyfragment.PolicyFragmentClient", "ListReferences", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListReferences(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyfragment.PolicyFragmentClient", "ListReferences", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListReferences prepares the ListReferences request.
func (c PolicyFragmentClient) preparerForListReferences(ctx context.Context, id PolicyFragmentId, options ListReferencesOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/listReferences", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListReferences handles the response to the ListReferences request. The method always
// closes the http.Response Body.
func (c PolicyFragmentClient) responderForListReferences(resp *http.Response) (result ListReferencesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
