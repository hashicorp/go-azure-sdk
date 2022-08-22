package producttag

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

type TagListByProductOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TagContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (TagListByProductOperationResponse, error)
}

type TagListByProductCompleteResult struct {
	Items []TagContract
}

func (r TagListByProductOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r TagListByProductOperationResponse) LoadMore(ctx context.Context) (resp TagListByProductOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type TagListByProductOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultTagListByProductOperationOptions() TagListByProductOperationOptions {
	return TagListByProductOperationOptions{}
}

func (o TagListByProductOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o TagListByProductOperationOptions) toQueryString() map[string]interface{} {
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

// TagListByProduct ...
func (c ProductTagClient) TagListByProduct(ctx context.Context, id ProductId, options TagListByProductOperationOptions) (resp TagListByProductOperationResponse, err error) {
	req, err := c.preparerForTagListByProduct(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagListByProduct", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagListByProduct", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForTagListByProduct(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagListByProduct", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForTagListByProduct prepares the TagListByProduct request.
func (c ProductTagClient) preparerForTagListByProduct(ctx context.Context, id ProductId, options TagListByProductOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/tags", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForTagListByProductWithNextLink prepares the TagListByProduct request with the given nextLink token.
func (c ProductTagClient) preparerForTagListByProductWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForTagListByProduct handles the response to the TagListByProduct request. The method always
// closes the http.Response Body.
func (c ProductTagClient) responderForTagListByProduct(resp *http.Response) (result TagListByProductOperationResponse, err error) {
	type page struct {
		Values   []TagContract `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result TagListByProductOperationResponse, err error) {
			req, err := c.preparerForTagListByProductWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagListByProduct", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagListByProduct", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForTagListByProduct(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "producttag.ProductTagClient", "TagListByProduct", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// TagListByProductComplete retrieves all of the results into a single object
func (c ProductTagClient) TagListByProductComplete(ctx context.Context, id ProductId, options TagListByProductOperationOptions) (TagListByProductCompleteResult, error) {
	return c.TagListByProductCompleteMatchingPredicate(ctx, id, options, TagContractOperationPredicate{})
}

// TagListByProductCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProductTagClient) TagListByProductCompleteMatchingPredicate(ctx context.Context, id ProductId, options TagListByProductOperationOptions, predicate TagContractOperationPredicate) (resp TagListByProductCompleteResult, err error) {
	items := make([]TagContract, 0)

	page, err := c.TagListByProduct(ctx, id, options)
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

	out := TagListByProductCompleteResult{
		Items: items,
	}
	return out, nil
}
