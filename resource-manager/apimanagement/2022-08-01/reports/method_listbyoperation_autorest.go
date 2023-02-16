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

type ListByOperationOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReportRecordContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByOperationOperationResponse, error)
}

type ListByOperationCompleteResult struct {
	Items []ReportRecordContract
}

func (r ListByOperationOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByOperationOperationResponse) LoadMore(ctx context.Context) (resp ListByOperationOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByOperationOperationOptions struct {
	Filter  *string
	Orderby *string
	Skip    *int64
	Top     *int64
}

func DefaultListByOperationOperationOptions() ListByOperationOperationOptions {
	return ListByOperationOperationOptions{}
}

func (o ListByOperationOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByOperationOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
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

// ListByOperation ...
func (c ReportsClient) ListByOperation(ctx context.Context, id ServiceId, options ListByOperationOperationOptions) (resp ListByOperationOperationResponse, err error) {
	req, err := c.preparerForListByOperation(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByOperation", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByOperation", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByOperation(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByOperation", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByOperation prepares the ListByOperation request.
func (c ReportsClient) preparerForListByOperation(ctx context.Context, id ServiceId, options ListByOperationOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/reports/byOperation", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByOperationWithNextLink prepares the ListByOperation request with the given nextLink token.
func (c ReportsClient) preparerForListByOperationWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByOperation handles the response to the ListByOperation request. The method always
// closes the http.Response Body.
func (c ReportsClient) responderForListByOperation(resp *http.Response) (result ListByOperationOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByOperationOperationResponse, err error) {
			req, err := c.preparerForListByOperationWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByOperation", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByOperation", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByOperation(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByOperation", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByOperationComplete retrieves all of the results into a single object
func (c ReportsClient) ListByOperationComplete(ctx context.Context, id ServiceId, options ListByOperationOperationOptions) (ListByOperationCompleteResult, error) {
	return c.ListByOperationCompleteMatchingPredicate(ctx, id, options, ReportRecordContractOperationPredicate{})
}

// ListByOperationCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReportsClient) ListByOperationCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ListByOperationOperationOptions, predicate ReportRecordContractOperationPredicate) (resp ListByOperationCompleteResult, err error) {
	items := make([]ReportRecordContract, 0)

	page, err := c.ListByOperation(ctx, id, options)
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

	out := ListByOperationCompleteResult{
		Items: items,
	}
	return out, nil
}
