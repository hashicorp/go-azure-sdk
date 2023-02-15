package apisbytag

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

type ApiListByTagsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TagResourceContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ApiListByTagsOperationResponse, error)
}

type ApiListByTagsCompleteResult struct {
	Items []TagResourceContract
}

func (r ApiListByTagsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ApiListByTagsOperationResponse) LoadMore(ctx context.Context) (resp ApiListByTagsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ApiListByTagsOperationOptions struct {
	Filter               *string
	IncludeNotTaggedApis *bool
	Skip                 *int64
	Top                  *int64
}

func DefaultApiListByTagsOperationOptions() ApiListByTagsOperationOptions {
	return ApiListByTagsOperationOptions{}
}

func (o ApiListByTagsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ApiListByTagsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.IncludeNotTaggedApis != nil {
		out["includeNotTaggedApis"] = *o.IncludeNotTaggedApis
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ApiListByTags ...
func (c ApisByTagClient) ApiListByTags(ctx context.Context, id ServiceId, options ApiListByTagsOperationOptions) (resp ApiListByTagsOperationResponse, err error) {
	req, err := c.preparerForApiListByTags(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apisbytag.ApisByTagClient", "ApiListByTags", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apisbytag.ApisByTagClient", "ApiListByTags", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForApiListByTags(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apisbytag.ApisByTagClient", "ApiListByTags", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForApiListByTags prepares the ApiListByTags request.
func (c ApisByTagClient) preparerForApiListByTags(ctx context.Context, id ServiceId, options ApiListByTagsOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/apisByTags", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForApiListByTagsWithNextLink prepares the ApiListByTags request with the given nextLink token.
func (c ApisByTagClient) preparerForApiListByTagsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForApiListByTags handles the response to the ApiListByTags request. The method always
// closes the http.Response Body.
func (c ApisByTagClient) responderForApiListByTags(resp *http.Response) (result ApiListByTagsOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ApiListByTagsOperationResponse, err error) {
			req, err := c.preparerForApiListByTagsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apisbytag.ApisByTagClient", "ApiListByTags", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "apisbytag.ApisByTagClient", "ApiListByTags", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForApiListByTags(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apisbytag.ApisByTagClient", "ApiListByTags", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ApiListByTagsComplete retrieves all of the results into a single object
func (c ApisByTagClient) ApiListByTagsComplete(ctx context.Context, id ServiceId, options ApiListByTagsOperationOptions) (ApiListByTagsCompleteResult, error) {
	return c.ApiListByTagsCompleteMatchingPredicate(ctx, id, options, TagResourceContractOperationPredicate{})
}

// ApiListByTagsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ApisByTagClient) ApiListByTagsCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ApiListByTagsOperationOptions, predicate TagResourceContractOperationPredicate) (resp ApiListByTagsCompleteResult, err error) {
	items := make([]TagResourceContract, 0)

	page, err := c.ApiListByTags(ctx, id, options)
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

	out := ApiListByTagsCompleteResult{
		Items: items,
	}
	return out, nil
}
