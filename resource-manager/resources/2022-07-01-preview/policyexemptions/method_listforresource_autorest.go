package policyexemptions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListForResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PolicyExemption

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListForResourceOperationResponse, error)
}

type ListForResourceCompleteResult struct {
	Items []PolicyExemption
}

func (r ListForResourceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListForResourceOperationResponse) LoadMore(ctx context.Context) (resp ListForResourceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListForResourceOperationOptions struct {
	Filter *string
}

func DefaultListForResourceOperationOptions() ListForResourceOperationOptions {
	return ListForResourceOperationOptions{}
}

func (o ListForResourceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListForResourceOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListForResource ...
func (c PolicyExemptionsClient) ListForResource(ctx context.Context, id ScopedResourceId, options ListForResourceOperationOptions) (resp ListForResourceOperationResponse, err error) {
	req, err := c.preparerForListForResource(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForResource", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForResource", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListForResource(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForResource", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListForResource prepares the ListForResource request.
func (c PolicyExemptionsClient) preparerForListForResource(ctx context.Context, id ScopedResourceId, options ListForResourceOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Authorization/policyExemptions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListForResourceWithNextLink prepares the ListForResource request with the given nextLink token.
func (c PolicyExemptionsClient) preparerForListForResourceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListForResource handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (c PolicyExemptionsClient) responderForListForResource(resp *http.Response) (result ListForResourceOperationResponse, err error) {
	type page struct {
		Values   []PolicyExemption `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListForResourceOperationResponse, err error) {
			req, err := c.preparerForListForResourceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForResource", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForResource", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListForResource(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForResource", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListForResourceComplete retrieves all of the results into a single object
func (c PolicyExemptionsClient) ListForResourceComplete(ctx context.Context, id ScopedResourceId, options ListForResourceOperationOptions) (ListForResourceCompleteResult, error) {
	return c.ListForResourceCompleteMatchingPredicate(ctx, id, options, PolicyExemptionOperationPredicate{})
}

// ListForResourceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PolicyExemptionsClient) ListForResourceCompleteMatchingPredicate(ctx context.Context, id ScopedResourceId, options ListForResourceOperationOptions, predicate PolicyExemptionOperationPredicate) (resp ListForResourceCompleteResult, err error) {
	items := make([]PolicyExemption, 0)

	page, err := c.ListForResource(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListForResourceCompleteResult{
		Items: items,
	}
	return out, nil
}
