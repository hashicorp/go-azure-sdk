package productsbytag

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

type ProductListByTagsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TagResourceContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ProductListByTagsOperationResponse, error)
}

type ProductListByTagsCompleteResult struct {
	Items []TagResourceContract
}

func (r ProductListByTagsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ProductListByTagsOperationResponse) LoadMore(ctx context.Context) (resp ProductListByTagsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ProductListByTagsOperationOptions struct {
	Filter                   *string
	IncludeNotTaggedProducts *bool
	Skip                     *int64
	Top                      *int64
}

func DefaultProductListByTagsOperationOptions() ProductListByTagsOperationOptions {
	return ProductListByTagsOperationOptions{}
}

func (o ProductListByTagsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ProductListByTagsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.IncludeNotTaggedProducts != nil {
		out["includeNotTaggedProducts"] = *o.IncludeNotTaggedProducts
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ProductListByTags ...
func (c ProductsByTagClient) ProductListByTags(ctx context.Context, id ServiceId, options ProductListByTagsOperationOptions) (resp ProductListByTagsOperationResponse, err error) {
	req, err := c.preparerForProductListByTags(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "productsbytag.ProductsByTagClient", "ProductListByTags", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "productsbytag.ProductsByTagClient", "ProductListByTags", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForProductListByTags(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "productsbytag.ProductsByTagClient", "ProductListByTags", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForProductListByTags prepares the ProductListByTags request.
func (c ProductsByTagClient) preparerForProductListByTags(ctx context.Context, id ServiceId, options ProductListByTagsOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/productsByTags", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForProductListByTagsWithNextLink prepares the ProductListByTags request with the given nextLink token.
func (c ProductsByTagClient) preparerForProductListByTagsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForProductListByTags handles the response to the ProductListByTags request. The method always
// closes the http.Response Body.
func (c ProductsByTagClient) responderForProductListByTags(resp *http.Response) (result ProductListByTagsOperationResponse, err error) {
	type page struct {
		Values   []TagResourceContract `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ProductListByTagsOperationResponse, err error) {
			req, err := c.preparerForProductListByTagsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "productsbytag.ProductsByTagClient", "ProductListByTags", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "productsbytag.ProductsByTagClient", "ProductListByTags", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForProductListByTags(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "productsbytag.ProductsByTagClient", "ProductListByTags", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ProductListByTagsComplete retrieves all of the results into a single object
func (c ProductsByTagClient) ProductListByTagsComplete(ctx context.Context, id ServiceId, options ProductListByTagsOperationOptions) (ProductListByTagsCompleteResult, error) {
	return c.ProductListByTagsCompleteMatchingPredicate(ctx, id, options, TagResourceContractOperationPredicate{})
}

// ProductListByTagsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProductsByTagClient) ProductListByTagsCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ProductListByTagsOperationOptions, predicate TagResourceContractOperationPredicate) (resp ProductListByTagsCompleteResult, err error) {
	items := make([]TagResourceContract, 0)

	page, err := c.ProductListByTags(ctx, id, options)
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

	out := ProductListByTagsCompleteResult{
		Items: items,
	}
	return out, nil
}
