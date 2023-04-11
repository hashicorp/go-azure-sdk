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

type IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]IoTSecurityAggregatedAlert

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse, error)
}

type IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteResult struct {
	Items []IoTSecurityAggregatedAlert
}

func (r IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse) LoadMore(ctx context.Context) (resp IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions struct {
	Top *int64
}

func DefaultIoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions() IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions {
	return IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions{}
}

func (o IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// IoTSecuritySolutionsAnalyticsAggregatedAlertsList ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsAggregatedAlertsList(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions) (resp IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIoTSecuritySolutionsAnalyticsAggregatedAlertsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertsList prepares the IoTSecuritySolutionsAnalyticsAggregatedAlertsList request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertsList(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/analyticsModels/default/aggregatedAlerts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertsListWithNextLink prepares the IoTSecuritySolutionsAnalyticsAggregatedAlertsList request with the given nextLink token.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIoTSecuritySolutionsAnalyticsAggregatedAlertsList handles the response to the IoTSecuritySolutionsAnalyticsAggregatedAlertsList request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsAggregatedAlertsList(resp *http.Response) (result IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse, err error) {
	type page struct {
		Values   []IoTSecurityAggregatedAlert `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationResponse, err error) {
			req, err := c.preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIoTSecuritySolutionsAnalyticsAggregatedAlertsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// IoTSecuritySolutionsAnalyticsAggregatedAlertsListComplete retrieves all of the results into a single object
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsAggregatedAlertsListComplete(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions) (IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteResult, error) {
	return c.IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteMatchingPredicate(ctx, id, options, IoTSecurityAggregatedAlertOperationPredicate{})
}

// IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteMatchingPredicate(ctx context.Context, id IotSecuritySolutionId, options IoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions, predicate IoTSecurityAggregatedAlertOperationPredicate) (resp IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteResult, err error) {
	items := make([]IoTSecurityAggregatedAlert, 0)

	page, err := c.IoTSecuritySolutionsAnalyticsAggregatedAlertsList(ctx, id, options)
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

	out := IoTSecuritySolutionsAnalyticsAggregatedAlertsListCompleteResult{
		Items: items,
	}
	return out, nil
}
