package recommendations

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

type ListHistoryForWebAppOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Recommendation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListHistoryForWebAppOperationResponse, error)
}

type ListHistoryForWebAppCompleteResult struct {
	Items []Recommendation
}

func (r ListHistoryForWebAppOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListHistoryForWebAppOperationResponse) LoadMore(ctx context.Context) (resp ListHistoryForWebAppOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListHistoryForWebAppOperationOptions struct {
	ExpiredOnly *bool
	Filter      *string
}

func DefaultListHistoryForWebAppOperationOptions() ListHistoryForWebAppOperationOptions {
	return ListHistoryForWebAppOperationOptions{}
}

func (o ListHistoryForWebAppOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListHistoryForWebAppOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ExpiredOnly != nil {
		out["expiredOnly"] = *o.ExpiredOnly
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListHistoryForWebApp ...
func (c RecommendationsClient) ListHistoryForWebApp(ctx context.Context, id SiteId, options ListHistoryForWebAppOperationOptions) (resp ListHistoryForWebAppOperationResponse, err error) {
	req, err := c.preparerForListHistoryForWebApp(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListHistoryForWebApp", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListHistoryForWebApp", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListHistoryForWebApp(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListHistoryForWebApp", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListHistoryForWebApp prepares the ListHistoryForWebApp request.
func (c RecommendationsClient) preparerForListHistoryForWebApp(ctx context.Context, id SiteId, options ListHistoryForWebAppOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/recommendationHistory", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListHistoryForWebAppWithNextLink prepares the ListHistoryForWebApp request with the given nextLink token.
func (c RecommendationsClient) preparerForListHistoryForWebAppWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListHistoryForWebApp handles the response to the ListHistoryForWebApp request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForListHistoryForWebApp(resp *http.Response) (result ListHistoryForWebAppOperationResponse, err error) {
	type page struct {
		Values   []Recommendation `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListHistoryForWebAppOperationResponse, err error) {
			req, err := c.preparerForListHistoryForWebAppWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListHistoryForWebApp", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListHistoryForWebApp", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListHistoryForWebApp(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListHistoryForWebApp", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListHistoryForWebAppComplete retrieves all of the results into a single object
func (c RecommendationsClient) ListHistoryForWebAppComplete(ctx context.Context, id SiteId, options ListHistoryForWebAppOperationOptions) (ListHistoryForWebAppCompleteResult, error) {
	return c.ListHistoryForWebAppCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// ListHistoryForWebAppCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RecommendationsClient) ListHistoryForWebAppCompleteMatchingPredicate(ctx context.Context, id SiteId, options ListHistoryForWebAppOperationOptions, predicate RecommendationOperationPredicate) (resp ListHistoryForWebAppCompleteResult, err error) {
	items := make([]Recommendation, 0)

	page, err := c.ListHistoryForWebApp(ctx, id, options)
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

	out := ListHistoryForWebAppCompleteResult{
		Items: items,
	}
	return out, nil
}
