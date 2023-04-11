package iotsecuritysolutionsanalytics

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

type IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]IoTSecurityAggregatedRecommendation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse, error)
}

type IoTSecuritySolutionsAnalyticsRecommendationsListCompleteResult struct {
	Items []IoTSecurityAggregatedRecommendation
}

func (r IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse) LoadMore(ctx context.Context) (resp IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions struct {
	Top *int64
}

func DefaultIoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions() IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions {
	return IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions{}
}

func (o IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// IoTSecuritySolutionsAnalyticsRecommendationsList ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsRecommendationsList(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions) (resp IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsRecommendationsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIoTSecuritySolutionsAnalyticsRecommendationsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForIoTSecuritySolutionsAnalyticsRecommendationsList prepares the IoTSecuritySolutionsAnalyticsRecommendationsList request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsRecommendationsList(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/analyticsModels/default/aggregatedRecommendations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForIoTSecuritySolutionsAnalyticsRecommendationsListWithNextLink prepares the IoTSecuritySolutionsAnalyticsRecommendationsList request with the given nextLink token.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsRecommendationsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIoTSecuritySolutionsAnalyticsRecommendationsList handles the response to the IoTSecuritySolutionsAnalyticsRecommendationsList request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsRecommendationsList(resp *http.Response) (result IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse, err error) {
	type page struct {
		Values   []IoTSecurityAggregatedRecommendation `json:"value"`
		NextLink *string                               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IoTSecuritySolutionsAnalyticsRecommendationsListOperationResponse, err error) {
			req, err := c.preparerForIoTSecuritySolutionsAnalyticsRecommendationsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIoTSecuritySolutionsAnalyticsRecommendationsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// IoTSecuritySolutionsAnalyticsRecommendationsListComplete retrieves all of the results into a single object
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsRecommendationsListComplete(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions) (IoTSecuritySolutionsAnalyticsRecommendationsListCompleteResult, error) {
	return c.IoTSecuritySolutionsAnalyticsRecommendationsListCompleteMatchingPredicate(ctx, id, options, IoTSecurityAggregatedRecommendationOperationPredicate{})
}

// IoTSecuritySolutionsAnalyticsRecommendationsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsRecommendationsListCompleteMatchingPredicate(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions, predicate IoTSecurityAggregatedRecommendationOperationPredicate) (resp IoTSecuritySolutionsAnalyticsRecommendationsListCompleteResult, err error) {
	items := make([]IoTSecurityAggregatedRecommendation, 0)

	page, err := c.IoTSecuritySolutionsAnalyticsRecommendationsList(ctx, id, options)
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

	out := IoTSecuritySolutionsAnalyticsRecommendationsListCompleteResult{
		Items: items,
	}
	return out, nil
}
