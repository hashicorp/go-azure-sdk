package graphqlapiresolver

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

type ListByApiOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ResolverContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByApiOperationResponse, error)
}

type ListByApiCompleteResult struct {
	Items []ResolverContract
}

func (r ListByApiOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByApiOperationResponse) LoadMore(ctx context.Context) (resp ListByApiOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByApiOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByApiOperationOptions() ListByApiOperationOptions {
	return ListByApiOperationOptions{}
}

func (o ListByApiOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByApiOperationOptions) toQueryString() map[string]interface{} {
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

// ListByApi ...
func (c GraphQLApiResolverClient) ListByApi(ctx context.Context, id ApiId, options ListByApiOperationOptions) (resp ListByApiOperationResponse, err error) {
	req, err := c.preparerForListByApi(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphqlapiresolver.GraphQLApiResolverClient", "ListByApi", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphqlapiresolver.GraphQLApiResolverClient", "ListByApi", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByApi(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphqlapiresolver.GraphQLApiResolverClient", "ListByApi", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByApi prepares the ListByApi request.
func (c GraphQLApiResolverClient) preparerForListByApi(ctx context.Context, id ApiId, options ListByApiOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/resolvers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByApiWithNextLink prepares the ListByApi request with the given nextLink token.
func (c GraphQLApiResolverClient) preparerForListByApiWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByApi handles the response to the ListByApi request. The method always
// closes the http.Response Body.
func (c GraphQLApiResolverClient) responderForListByApi(resp *http.Response) (result ListByApiOperationResponse, err error) {
	type page struct {
		Values   []ResolverContract `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByApiOperationResponse, err error) {
			req, err := c.preparerForListByApiWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "graphqlapiresolver.GraphQLApiResolverClient", "ListByApi", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "graphqlapiresolver.GraphQLApiResolverClient", "ListByApi", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByApi(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "graphqlapiresolver.GraphQLApiResolverClient", "ListByApi", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByApiComplete retrieves all of the results into a single object
func (c GraphQLApiResolverClient) ListByApiComplete(ctx context.Context, id ApiId, options ListByApiOperationOptions) (ListByApiCompleteResult, error) {
	return c.ListByApiCompleteMatchingPredicate(ctx, id, options, ResolverContractOperationPredicate{})
}

// ListByApiCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c GraphQLApiResolverClient) ListByApiCompleteMatchingPredicate(ctx context.Context, id ApiId, options ListByApiOperationOptions, predicate ResolverContractOperationPredicate) (resp ListByApiCompleteResult, err error) {
	items := make([]ResolverContract, 0)

	page, err := c.ListByApi(ctx, id, options)
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

	out := ListByApiCompleteResult{
		Items: items,
	}
	return out, nil
}
