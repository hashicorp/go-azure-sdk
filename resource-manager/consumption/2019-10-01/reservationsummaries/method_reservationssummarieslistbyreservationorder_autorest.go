package reservationsummaries

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

type ReservationsSummariesListByReservationOrderOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReservationSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ReservationsSummariesListByReservationOrderOperationResponse, error)
}

type ReservationsSummariesListByReservationOrderCompleteResult struct {
	Items []ReservationSummary
}

func (r ReservationsSummariesListByReservationOrderOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ReservationsSummariesListByReservationOrderOperationResponse) LoadMore(ctx context.Context) (resp ReservationsSummariesListByReservationOrderOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ReservationsSummariesListByReservationOrderOperationOptions struct {
	Filter *string
	Grain  *Datagrain
}

func DefaultReservationsSummariesListByReservationOrderOperationOptions() ReservationsSummariesListByReservationOrderOperationOptions {
	return ReservationsSummariesListByReservationOrderOperationOptions{}
}

func (o ReservationsSummariesListByReservationOrderOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ReservationsSummariesListByReservationOrderOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Grain != nil {
		out["grain"] = *o.Grain
	}

	return out
}

// ReservationsSummariesListByReservationOrder ...
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrder(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions) (resp ReservationsSummariesListByReservationOrderOperationResponse, err error) {
	req, err := c.preparerForReservationsSummariesListByReservationOrder(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrder", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrder", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForReservationsSummariesListByReservationOrder(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrder", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForReservationsSummariesListByReservationOrder prepares the ReservationsSummariesListByReservationOrder request.
func (c ReservationSummariesClient) preparerForReservationsSummariesListByReservationOrder(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationSummaries", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForReservationsSummariesListByReservationOrderWithNextLink prepares the ReservationsSummariesListByReservationOrder request with the given nextLink token.
func (c ReservationSummariesClient) preparerForReservationsSummariesListByReservationOrderWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForReservationsSummariesListByReservationOrder handles the response to the ReservationsSummariesListByReservationOrder request. The method always
// closes the http.Response Body.
func (c ReservationSummariesClient) responderForReservationsSummariesListByReservationOrder(resp *http.Response) (result ReservationsSummariesListByReservationOrderOperationResponse, err error) {
	type page struct {
		Values   []ReservationSummary `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ReservationsSummariesListByReservationOrderOperationResponse, err error) {
			req, err := c.preparerForReservationsSummariesListByReservationOrderWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrder", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrder", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForReservationsSummariesListByReservationOrder(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrder", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ReservationsSummariesListByReservationOrderComplete retrieves all of the results into a single object
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderComplete(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions) (ReservationsSummariesListByReservationOrderCompleteResult, error) {
	return c.ReservationsSummariesListByReservationOrderCompleteMatchingPredicate(ctx, id, options, ReservationSummaryOperationPredicate{})
}

// ReservationsSummariesListByReservationOrderCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderCompleteMatchingPredicate(ctx context.Context, id ReservationOrderId, options ReservationsSummariesListByReservationOrderOperationOptions, predicate ReservationSummaryOperationPredicate) (resp ReservationsSummariesListByReservationOrderCompleteResult, err error) {
	items := make([]ReservationSummary, 0)

	page, err := c.ReservationsSummariesListByReservationOrder(ctx, id, options)
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

	out := ReservationsSummariesListByReservationOrderCompleteResult{
		Items: items,
	}
	return out, nil
}
