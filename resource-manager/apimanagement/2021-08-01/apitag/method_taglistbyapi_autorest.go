package apitag

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

type TagListByApiOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TagContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (TagListByApiOperationResponse, error)
}

type TagListByApiCompleteResult struct {
	Items []TagContract
}

func (r TagListByApiOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r TagListByApiOperationResponse) LoadMore(ctx context.Context) (resp TagListByApiOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type TagListByApiOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultTagListByApiOperationOptions() TagListByApiOperationOptions {
	return TagListByApiOperationOptions{}
}

func (o TagListByApiOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o TagListByApiOperationOptions) toQueryString() map[string]interface{} {
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

// TagListByApi ...
func (c ApiTagClient) TagListByApi(ctx context.Context, id ApiId, options TagListByApiOperationOptions) (resp TagListByApiOperationResponse, err error) {
	req, err := c.preparerForTagListByApi(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagListByApi", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagListByApi", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForTagListByApi(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagListByApi", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForTagListByApi prepares the TagListByApi request.
func (c ApiTagClient) preparerForTagListByApi(ctx context.Context, id ApiId, options TagListByApiOperationOptions) (*http.Request, error) {
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

// preparerForTagListByApiWithNextLink prepares the TagListByApi request with the given nextLink token.
func (c ApiTagClient) preparerForTagListByApiWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForTagListByApi handles the response to the TagListByApi request. The method always
// closes the http.Response Body.
func (c ApiTagClient) responderForTagListByApi(resp *http.Response) (result TagListByApiOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result TagListByApiOperationResponse, err error) {
			req, err := c.preparerForTagListByApiWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagListByApi", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagListByApi", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForTagListByApi(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagListByApi", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// TagListByApiComplete retrieves all of the results into a single object
func (c ApiTagClient) TagListByApiComplete(ctx context.Context, id ApiId, options TagListByApiOperationOptions) (TagListByApiCompleteResult, error) {
	return c.TagListByApiCompleteMatchingPredicate(ctx, id, options, TagContractOperationPredicate{})
}

// TagListByApiCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ApiTagClient) TagListByApiCompleteMatchingPredicate(ctx context.Context, id ApiId, options TagListByApiOperationOptions, predicate TagContractOperationPredicate) (resp TagListByApiCompleteResult, err error) {
	items := make([]TagContract, 0)

	page, err := c.TagListByApi(ctx, id, options)
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

	out := TagListByApiCompleteResult{
		Items: items,
	}
	return out, nil
}
