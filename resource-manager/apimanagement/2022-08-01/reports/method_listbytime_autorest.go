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

type ListByTimeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReportRecordContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByTimeOperationResponse, error)
}

type ListByTimeCompleteResult struct {
	Items []ReportRecordContract
}

func (r ListByTimeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByTimeOperationResponse) LoadMore(ctx context.Context) (resp ListByTimeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByTimeOperationOptions struct {
	Filter   *string
	Interval *string
	Orderby  *string
	Skip     *int64
	Top      *int64
}

func DefaultListByTimeOperationOptions() ListByTimeOperationOptions {
	return ListByTimeOperationOptions{}
}

func (o ListByTimeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByTimeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Interval != nil {
		out["interval"] = *o.Interval
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByTime ...
func (c ReportsClient) ListByTime(ctx context.Context, id ServiceId, options ListByTimeOperationOptions) (resp ListByTimeOperationResponse, err error) {
	req, err := c.preparerForListByTime(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByTime", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByTime", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByTime(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByTime", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByTime prepares the ListByTime request.
func (c ReportsClient) preparerForListByTime(ctx context.Context, id ServiceId, options ListByTimeOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/reports/byTime", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByTimeWithNextLink prepares the ListByTime request with the given nextLink token.
func (c ReportsClient) preparerForListByTimeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByTime handles the response to the ListByTime request. The method always
// closes the http.Response Body.
func (c ReportsClient) responderForListByTime(resp *http.Response) (result ListByTimeOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByTimeOperationResponse, err error) {
			req, err := c.preparerForListByTimeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByTime", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByTime", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByTime(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByTime", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByTimeComplete retrieves all of the results into a single object
func (c ReportsClient) ListByTimeComplete(ctx context.Context, id ServiceId, options ListByTimeOperationOptions) (ListByTimeCompleteResult, error) {
	return c.ListByTimeCompleteMatchingPredicate(ctx, id, options, ReportRecordContractOperationPredicate{})
}

// ListByTimeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReportsClient) ListByTimeCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ListByTimeOperationOptions, predicate ReportRecordContractOperationPredicate) (resp ListByTimeCompleteResult, err error) {
	items := make([]ReportRecordContract, 0)

	page, err := c.ListByTime(ctx, id, options)
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

	out := ListByTimeCompleteResult{
		Items: items,
	}
	return out, nil
}
