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

type ListByServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicyFragmentCollection
}

type ListByServiceOperationOptions struct {
	Filter  *string
	Orderby *string
	Skip    *int64
	Top     *int64
}

func DefaultListByServiceOperationOptions() ListByServiceOperationOptions {
	return ListByServiceOperationOptions{}
}

func (o ListByServiceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByServiceOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByService ...
func (c PolicyFragmentClient) ListByService(ctx context.Context, id ServiceId, options ListByServiceOperationOptions) (result ListByServiceOperationResponse, err error) {
	req, err := c.preparerForListByService(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyfragment.PolicyFragmentClient", "ListByService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyfragment.PolicyFragmentClient", "ListByService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyfragment.PolicyFragmentClient", "ListByService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByService prepares the ListByService request.
func (c PolicyFragmentClient) preparerForListByService(ctx context.Context, id ServiceId, options ListByServiceOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/policyFragments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByService handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (c PolicyFragmentClient) responderForListByService(resp *http.Response) (result ListByServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
