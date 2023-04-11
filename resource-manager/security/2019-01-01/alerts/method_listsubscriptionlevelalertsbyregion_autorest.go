package alerts

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

type ListSubscriptionLevelAlertsByRegionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Alert

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSubscriptionLevelAlertsByRegionOperationResponse, error)
}

type ListSubscriptionLevelAlertsByRegionCompleteResult struct {
	Items []Alert
}

func (r ListSubscriptionLevelAlertsByRegionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSubscriptionLevelAlertsByRegionOperationResponse) LoadMore(ctx context.Context) (resp ListSubscriptionLevelAlertsByRegionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListSubscriptionLevelAlertsByRegionOperationOptions struct {
	AutoDismissRuleName *string
	Expand              *string
	Filter              *string
	Select              *string
}

func DefaultListSubscriptionLevelAlertsByRegionOperationOptions() ListSubscriptionLevelAlertsByRegionOperationOptions {
	return ListSubscriptionLevelAlertsByRegionOperationOptions{}
}

func (o ListSubscriptionLevelAlertsByRegionOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListSubscriptionLevelAlertsByRegionOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.AutoDismissRuleName != nil {
		out["autoDismissRuleName"] = *o.AutoDismissRuleName
	}

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Select != nil {
		out["$select"] = *o.Select
	}

	return out
}

// ListSubscriptionLevelAlertsByRegion ...
func (c AlertsClient) ListSubscriptionLevelAlertsByRegion(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions) (resp ListSubscriptionLevelAlertsByRegionOperationResponse, err error) {
	req, err := c.preparerForListSubscriptionLevelAlertsByRegion(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListSubscriptionLevelAlertsByRegion", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListSubscriptionLevelAlertsByRegion", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSubscriptionLevelAlertsByRegion(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListSubscriptionLevelAlertsByRegion", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSubscriptionLevelAlertsByRegion prepares the ListSubscriptionLevelAlertsByRegion request.
func (c AlertsClient) preparerForListSubscriptionLevelAlertsByRegion(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/alerts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSubscriptionLevelAlertsByRegionWithNextLink prepares the ListSubscriptionLevelAlertsByRegion request with the given nextLink token.
func (c AlertsClient) preparerForListSubscriptionLevelAlertsByRegionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSubscriptionLevelAlertsByRegion handles the response to the ListSubscriptionLevelAlertsByRegion request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForListSubscriptionLevelAlertsByRegion(resp *http.Response) (result ListSubscriptionLevelAlertsByRegionOperationResponse, err error) {
	type page struct {
		Values   []Alert `json:"value"`
		NextLink *string `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSubscriptionLevelAlertsByRegionOperationResponse, err error) {
			req, err := c.preparerForListSubscriptionLevelAlertsByRegionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListSubscriptionLevelAlertsByRegion", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListSubscriptionLevelAlertsByRegion", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSubscriptionLevelAlertsByRegion(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListSubscriptionLevelAlertsByRegion", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSubscriptionLevelAlertsByRegionComplete retrieves all of the results into a single object
func (c AlertsClient) ListSubscriptionLevelAlertsByRegionComplete(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions) (ListSubscriptionLevelAlertsByRegionCompleteResult, error) {
	return c.ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate(ctx, id, options, AlertOperationPredicate{})
}

// ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AlertsClient) ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions, predicate AlertOperationPredicate) (resp ListSubscriptionLevelAlertsByRegionCompleteResult, err error) {
	items := make([]Alert, 0)

	page, err := c.ListSubscriptionLevelAlertsByRegion(ctx, id, options)
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

	out := ListSubscriptionLevelAlertsByRegionCompleteResult{
		Items: items,
	}
	return out, nil
}
