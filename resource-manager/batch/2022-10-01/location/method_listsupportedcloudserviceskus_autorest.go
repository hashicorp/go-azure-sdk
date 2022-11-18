package location

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

type ListSupportedCloudServiceSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SupportedSku

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSupportedCloudServiceSkusOperationResponse, error)
}

type ListSupportedCloudServiceSkusCompleteResult struct {
	Items []SupportedSku
}

func (r ListSupportedCloudServiceSkusOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSupportedCloudServiceSkusOperationResponse) LoadMore(ctx context.Context) (resp ListSupportedCloudServiceSkusOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListSupportedCloudServiceSkusOperationOptions struct {
	Filter     *string
	Maxresults *int64
}

func DefaultListSupportedCloudServiceSkusOperationOptions() ListSupportedCloudServiceSkusOperationOptions {
	return ListSupportedCloudServiceSkusOperationOptions{}
}

func (o ListSupportedCloudServiceSkusOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListSupportedCloudServiceSkusOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Maxresults != nil {
		out["maxresults"] = *o.Maxresults
	}

	return out
}

// ListSupportedCloudServiceSkus ...
func (c LocationClient) ListSupportedCloudServiceSkus(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions) (resp ListSupportedCloudServiceSkusOperationResponse, err error) {
	req, err := c.preparerForListSupportedCloudServiceSkus(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.LocationClient", "ListSupportedCloudServiceSkus", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.LocationClient", "ListSupportedCloudServiceSkus", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSupportedCloudServiceSkus(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.LocationClient", "ListSupportedCloudServiceSkus", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSupportedCloudServiceSkus prepares the ListSupportedCloudServiceSkus request.
func (c LocationClient) preparerForListSupportedCloudServiceSkus(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/cloudServiceSkus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSupportedCloudServiceSkusWithNextLink prepares the ListSupportedCloudServiceSkus request with the given nextLink token.
func (c LocationClient) preparerForListSupportedCloudServiceSkusWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSupportedCloudServiceSkus handles the response to the ListSupportedCloudServiceSkus request. The method always
// closes the http.Response Body.
func (c LocationClient) responderForListSupportedCloudServiceSkus(resp *http.Response) (result ListSupportedCloudServiceSkusOperationResponse, err error) {
	type page struct {
		Values   []SupportedSku `json:"value"`
		NextLink *string        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSupportedCloudServiceSkusOperationResponse, err error) {
			req, err := c.preparerForListSupportedCloudServiceSkusWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "location.LocationClient", "ListSupportedCloudServiceSkus", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "location.LocationClient", "ListSupportedCloudServiceSkus", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSupportedCloudServiceSkus(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "location.LocationClient", "ListSupportedCloudServiceSkus", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSupportedCloudServiceSkusComplete retrieves all of the results into a single object
func (c LocationClient) ListSupportedCloudServiceSkusComplete(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions) (ListSupportedCloudServiceSkusCompleteResult, error) {
	return c.ListSupportedCloudServiceSkusCompleteMatchingPredicate(ctx, id, options, SupportedSkuOperationPredicate{})
}

// ListSupportedCloudServiceSkusCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c LocationClient) ListSupportedCloudServiceSkusCompleteMatchingPredicate(ctx context.Context, id LocationId, options ListSupportedCloudServiceSkusOperationOptions, predicate SupportedSkuOperationPredicate) (resp ListSupportedCloudServiceSkusCompleteResult, err error) {
	items := make([]SupportedSku, 0)

	page, err := c.ListSupportedCloudServiceSkus(ctx, id, options)
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

	out := ListSupportedCloudServiceSkusCompleteResult{
		Items: items,
	}
	return out, nil
}
