package lots

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

type ListByCustomerOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]LotSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByCustomerOperationResponse, error)
}

type ListByCustomerCompleteResult struct {
	Items []LotSummary
}

func (r ListByCustomerOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByCustomerOperationResponse) LoadMore(ctx context.Context) (resp ListByCustomerOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByCustomerOperationOptions struct {
	Filter *string
}

func DefaultListByCustomerOperationOptions() ListByCustomerOperationOptions {
	return ListByCustomerOperationOptions{}
}

func (o ListByCustomerOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByCustomerOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByCustomer ...
func (c LotsClient) ListByCustomer(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions) (resp ListByCustomerOperationResponse, err error) {
	req, err := c.preparerForListByCustomer(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByCustomer", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByCustomer", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByCustomer(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByCustomer", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByCustomer prepares the ListByCustomer request.
func (c LotsClient) preparerForListByCustomer(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/lots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByCustomerWithNextLink prepares the ListByCustomer request with the given nextLink token.
func (c LotsClient) preparerForListByCustomerWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByCustomer handles the response to the ListByCustomer request. The method always
// closes the http.Response Body.
func (c LotsClient) responderForListByCustomer(resp *http.Response) (result ListByCustomerOperationResponse, err error) {
	type page struct {
		Values   []LotSummary `json:"value"`
		NextLink *string      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByCustomerOperationResponse, err error) {
			req, err := c.preparerForListByCustomerWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByCustomer", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByCustomer", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByCustomer(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByCustomer", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByCustomerComplete retrieves all of the results into a single object
func (c LotsClient) ListByCustomerComplete(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions) (ListByCustomerCompleteResult, error) {
	return c.ListByCustomerCompleteMatchingPredicate(ctx, id, options, LotSummaryOperationPredicate{})
}

// ListByCustomerCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c LotsClient) ListByCustomerCompleteMatchingPredicate(ctx context.Context, id CustomerId, options ListByCustomerOperationOptions, predicate LotSummaryOperationPredicate) (resp ListByCustomerCompleteResult, err error) {
	items := make([]LotSummary, 0)

	page, err := c.ListByCustomer(ctx, id, options)
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

	out := ListByCustomerCompleteResult{
		Items: items,
	}
	return out, nil
}
