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

type ListRecommendedRulesForWebAppOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Recommendation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListRecommendedRulesForWebAppOperationResponse, error)
}

type ListRecommendedRulesForWebAppCompleteResult struct {
	Items []Recommendation
}

func (r ListRecommendedRulesForWebAppOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListRecommendedRulesForWebAppOperationResponse) LoadMore(ctx context.Context) (resp ListRecommendedRulesForWebAppOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListRecommendedRulesForWebAppOperationOptions struct {
	Featured *bool
	Filter   *string
}

func DefaultListRecommendedRulesForWebAppOperationOptions() ListRecommendedRulesForWebAppOperationOptions {
	return ListRecommendedRulesForWebAppOperationOptions{}
}

func (o ListRecommendedRulesForWebAppOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListRecommendedRulesForWebAppOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Featured != nil {
		out["featured"] = *o.Featured
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListRecommendedRulesForWebApp ...
func (c RecommendationsClient) ListRecommendedRulesForWebApp(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions) (resp ListRecommendedRulesForWebAppOperationResponse, err error) {
	req, err := c.preparerForListRecommendedRulesForWebApp(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListRecommendedRulesForWebApp", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListRecommendedRulesForWebApp", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListRecommendedRulesForWebApp(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListRecommendedRulesForWebApp", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListRecommendedRulesForWebApp prepares the ListRecommendedRulesForWebApp request.
func (c RecommendationsClient) preparerForListRecommendedRulesForWebApp(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/recommendations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListRecommendedRulesForWebAppWithNextLink prepares the ListRecommendedRulesForWebApp request with the given nextLink token.
func (c RecommendationsClient) preparerForListRecommendedRulesForWebAppWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListRecommendedRulesForWebApp handles the response to the ListRecommendedRulesForWebApp request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForListRecommendedRulesForWebApp(resp *http.Response) (result ListRecommendedRulesForWebAppOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListRecommendedRulesForWebAppOperationResponse, err error) {
			req, err := c.preparerForListRecommendedRulesForWebAppWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListRecommendedRulesForWebApp", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListRecommendedRulesForWebApp", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListRecommendedRulesForWebApp(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "ListRecommendedRulesForWebApp", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListRecommendedRulesForWebAppComplete retrieves all of the results into a single object
func (c RecommendationsClient) ListRecommendedRulesForWebAppComplete(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions) (ListRecommendedRulesForWebAppCompleteResult, error) {
	return c.ListRecommendedRulesForWebAppCompleteMatchingPredicate(ctx, id, options, RecommendationOperationPredicate{})
}

// ListRecommendedRulesForWebAppCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RecommendationsClient) ListRecommendedRulesForWebAppCompleteMatchingPredicate(ctx context.Context, id SiteId, options ListRecommendedRulesForWebAppOperationOptions, predicate RecommendationOperationPredicate) (resp ListRecommendedRulesForWebAppCompleteResult, err error) {
	items := make([]Recommendation, 0)

	page, err := c.ListRecommendedRulesForWebApp(ctx, id, options)
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

	out := ListRecommendedRulesForWebAppCompleteResult{
		Items: items,
	}
	return out, nil
}
