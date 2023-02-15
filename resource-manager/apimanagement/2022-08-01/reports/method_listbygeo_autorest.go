package reports

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

type ListByGeoOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReportRecordContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByGeoOperationResponse, error)
}

type ListByGeoCompleteResult struct {
	Items []ReportRecordContract
}

func (r ListByGeoOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByGeoOperationResponse) LoadMore(ctx context.Context) (resp ListByGeoOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByGeoOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByGeoOperationOptions() ListByGeoOperationOptions {
	return ListByGeoOperationOptions{}
}

func (o ListByGeoOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByGeoOperationOptions) toQueryString() map[string]interface{} {
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

// ListByGeo ...
func (c ReportsClient) ListByGeo(ctx context.Context, id ServiceId, options ListByGeoOperationOptions) (resp ListByGeoOperationResponse, err error) {
	req, err := c.preparerForListByGeo(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByGeo", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByGeo", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByGeo(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByGeo", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByGeo prepares the ListByGeo request.
func (c ReportsClient) preparerForListByGeo(ctx context.Context, id ServiceId, options ListByGeoOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/reports/byGeo", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByGeoWithNextLink prepares the ListByGeo request with the given nextLink token.
func (c ReportsClient) preparerForListByGeoWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByGeo handles the response to the ListByGeo request. The method always
// closes the http.Response Body.
func (c ReportsClient) responderForListByGeo(resp *http.Response) (result ListByGeoOperationResponse, err error) {
	type page struct {
		Values   []ReportRecordContract `json:"value"`
		NextLink *string                `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByGeoOperationResponse, err error) {
			req, err := c.preparerForListByGeoWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByGeo", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByGeo", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByGeo(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByGeo", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByGeoComplete retrieves all of the results into a single object
func (c ReportsClient) ListByGeoComplete(ctx context.Context, id ServiceId, options ListByGeoOperationOptions) (ListByGeoCompleteResult, error) {
	return c.ListByGeoCompleteMatchingPredicate(ctx, id, options, ReportRecordContractOperationPredicate{})
}

// ListByGeoCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReportsClient) ListByGeoCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ListByGeoOperationOptions, predicate ReportRecordContractOperationPredicate) (resp ListByGeoCompleteResult, err error) {
	items := make([]ReportRecordContract, 0)

	page, err := c.ListByGeo(ctx, id, options)
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

	out := ListByGeoCompleteResult{
		Items: items,
	}
	return out, nil
}
