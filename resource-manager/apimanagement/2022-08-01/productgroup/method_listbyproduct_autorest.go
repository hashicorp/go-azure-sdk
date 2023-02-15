package productgroup

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

type ListByProductOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GroupContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByProductOperationResponse, error)
}

type ListByProductCompleteResult struct {
	Items []GroupContract
}

func (r ListByProductOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByProductOperationResponse) LoadMore(ctx context.Context) (resp ListByProductOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByProductOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByProductOperationOptions() ListByProductOperationOptions {
	return ListByProductOperationOptions{}
}

func (o ListByProductOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByProductOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByProduct ...
func (c ProductGroupClient) ListByProduct(ctx context.Context, id ProductId, options ListByProductOperationOptions) (resp ListByProductOperationResponse, err error) {
	req, err := c.preparerForListByProduct(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "productgroup.ProductGroupClient", "ListByProduct", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "productgroup.ProductGroupClient", "ListByProduct", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByProduct(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "productgroup.ProductGroupClient", "ListByProduct", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByProduct prepares the ListByProduct request.
func (c ProductGroupClient) preparerForListByProduct(ctx context.Context, id ProductId, options ListByProductOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/groups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByProductWithNextLink prepares the ListByProduct request with the given nextLink token.
func (c ProductGroupClient) preparerForListByProductWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByProduct handles the response to the ListByProduct request. The method always
// closes the http.Response Body.
func (c ProductGroupClient) responderForListByProduct(resp *http.Response) (result ListByProductOperationResponse, err error) {
	type page struct {
		Values   []GroupContract `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByProductOperationResponse, err error) {
			req, err := c.preparerForListByProductWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "productgroup.ProductGroupClient", "ListByProduct", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "productgroup.ProductGroupClient", "ListByProduct", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByProduct(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "productgroup.ProductGroupClient", "ListByProduct", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByProductComplete retrieves all of the results into a single object
func (c ProductGroupClient) ListByProductComplete(ctx context.Context, id ProductId, options ListByProductOperationOptions) (ListByProductCompleteResult, error) {
	return c.ListByProductCompleteMatchingPredicate(ctx, id, options, GroupContractOperationPredicate{})
}

// ListByProductCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProductGroupClient) ListByProductCompleteMatchingPredicate(ctx context.Context, id ProductId, options ListByProductOperationOptions, predicate GroupContractOperationPredicate) (resp ListByProductCompleteResult, err error) {
	items := make([]GroupContract, 0)

	page, err := c.ListByProduct(ctx, id, options)
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

	out := ListByProductCompleteResult{
		Items: items,
	}
	return out, nil
}
