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

type ListResourceGroupLevelAlertsByRegionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Alert

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListResourceGroupLevelAlertsByRegionOperationResponse, error)
}

type ListResourceGroupLevelAlertsByRegionCompleteResult struct {
	Items []Alert
}

func (r ListResourceGroupLevelAlertsByRegionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListResourceGroupLevelAlertsByRegionOperationResponse) LoadMore(ctx context.Context) (resp ListResourceGroupLevelAlertsByRegionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListResourceGroupLevelAlertsByRegion ...
func (c AlertsClient) ListResourceGroupLevelAlertsByRegion(ctx context.Context, id ProviderLocationId) (resp ListResourceGroupLevelAlertsByRegionOperationResponse, err error) {
	req, err := c.preparerForListResourceGroupLevelAlertsByRegion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListResourceGroupLevelAlertsByRegion", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListResourceGroupLevelAlertsByRegion", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListResourceGroupLevelAlertsByRegion(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListResourceGroupLevelAlertsByRegion", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListResourceGroupLevelAlertsByRegion prepares the ListResourceGroupLevelAlertsByRegion request.
func (c AlertsClient) preparerForListResourceGroupLevelAlertsByRegion(ctx context.Context, id ProviderLocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/alerts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListResourceGroupLevelAlertsByRegionWithNextLink prepares the ListResourceGroupLevelAlertsByRegion request with the given nextLink token.
func (c AlertsClient) preparerForListResourceGroupLevelAlertsByRegionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListResourceGroupLevelAlertsByRegion handles the response to the ListResourceGroupLevelAlertsByRegion request. The method always
// closes the http.Response Body.
func (c AlertsClient) responderForListResourceGroupLevelAlertsByRegion(resp *http.Response) (result ListResourceGroupLevelAlertsByRegionOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListResourceGroupLevelAlertsByRegionOperationResponse, err error) {
			req, err := c.preparerForListResourceGroupLevelAlertsByRegionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListResourceGroupLevelAlertsByRegion", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListResourceGroupLevelAlertsByRegion", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListResourceGroupLevelAlertsByRegion(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "alerts.AlertsClient", "ListResourceGroupLevelAlertsByRegion", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListResourceGroupLevelAlertsByRegionComplete retrieves all of the results into a single object
func (c AlertsClient) ListResourceGroupLevelAlertsByRegionComplete(ctx context.Context, id ProviderLocationId) (ListResourceGroupLevelAlertsByRegionCompleteResult, error) {
	return c.ListResourceGroupLevelAlertsByRegionCompleteMatchingPredicate(ctx, id, AlertOperationPredicate{})
}

// ListResourceGroupLevelAlertsByRegionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AlertsClient) ListResourceGroupLevelAlertsByRegionCompleteMatchingPredicate(ctx context.Context, id ProviderLocationId, predicate AlertOperationPredicate) (resp ListResourceGroupLevelAlertsByRegionCompleteResult, err error) {
	items := make([]Alert, 0)

	page, err := c.ListResourceGroupLevelAlertsByRegion(ctx, id)
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

	out := ListResourceGroupLevelAlertsByRegionCompleteResult{
		Items: items,
	}
	return out, nil
}
