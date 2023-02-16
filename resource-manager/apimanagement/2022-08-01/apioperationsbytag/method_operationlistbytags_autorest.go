package apioperationsbytag

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

type OperationListByTagsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TagResourceContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (OperationListByTagsOperationResponse, error)
}

type OperationListByTagsCompleteResult struct {
	Items []TagResourceContract
}

func (r OperationListByTagsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r OperationListByTagsOperationResponse) LoadMore(ctx context.Context) (resp OperationListByTagsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type OperationListByTagsOperationOptions struct {
	Filter                     *string
	IncludeNotTaggedOperations *bool
	Skip                       *int64
	Top                        *int64
}

func DefaultOperationListByTagsOperationOptions() OperationListByTagsOperationOptions {
	return OperationListByTagsOperationOptions{}
}

func (o OperationListByTagsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o OperationListByTagsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.IncludeNotTaggedOperations != nil {
		out["includeNotTaggedOperations"] = *o.IncludeNotTaggedOperations
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// OperationListByTags ...
func (c ApiOperationsByTagClient) OperationListByTags(ctx context.Context, id ApiId, options OperationListByTagsOperationOptions) (resp OperationListByTagsOperationResponse, err error) {
	req, err := c.preparerForOperationListByTags(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationsbytag.ApiOperationsByTagClient", "OperationListByTags", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationsbytag.ApiOperationsByTagClient", "OperationListByTags", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForOperationListByTags(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationsbytag.ApiOperationsByTagClient", "OperationListByTags", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForOperationListByTags prepares the OperationListByTags request.
func (c ApiOperationsByTagClient) preparerForOperationListByTags(ctx context.Context, id ApiId, options OperationListByTagsOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/operationsByTags", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForOperationListByTagsWithNextLink prepares the OperationListByTags request with the given nextLink token.
func (c ApiOperationsByTagClient) preparerForOperationListByTagsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForOperationListByTags handles the response to the OperationListByTags request. The method always
// closes the http.Response Body.
func (c ApiOperationsByTagClient) responderForOperationListByTags(resp *http.Response) (result OperationListByTagsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result OperationListByTagsOperationResponse, err error) {
			req, err := c.preparerForOperationListByTagsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apioperationsbytag.ApiOperationsByTagClient", "OperationListByTags", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "apioperationsbytag.ApiOperationsByTagClient", "OperationListByTags", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForOperationListByTags(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apioperationsbytag.ApiOperationsByTagClient", "OperationListByTags", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// OperationListByTagsComplete retrieves all of the results into a single object
func (c ApiOperationsByTagClient) OperationListByTagsComplete(ctx context.Context, id ApiId, options OperationListByTagsOperationOptions) (OperationListByTagsCompleteResult, error) {
	return c.OperationListByTagsCompleteMatchingPredicate(ctx, id, options, TagResourceContractOperationPredicate{})
}

// OperationListByTagsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ApiOperationsByTagClient) OperationListByTagsCompleteMatchingPredicate(ctx context.Context, id ApiId, options OperationListByTagsOperationOptions, predicate TagResourceContractOperationPredicate) (resp OperationListByTagsCompleteResult, err error) {
	items := make([]TagResourceContract, 0)

	page, err := c.OperationListByTags(ctx, id, options)
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

	out := OperationListByTagsCompleteResult{
		Items: items,
	}
	return out, nil
}
