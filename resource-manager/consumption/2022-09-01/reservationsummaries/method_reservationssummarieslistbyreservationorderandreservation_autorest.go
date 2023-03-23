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

type ReservationsSummariesListByReservationOrderAndReservationOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReservationSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ReservationsSummariesListByReservationOrderAndReservationOperationResponse, error)
}

type ReservationsSummariesListByReservationOrderAndReservationCompleteResult struct {
	Items []ReservationSummary
}

func (r ReservationsSummariesListByReservationOrderAndReservationOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ReservationsSummariesListByReservationOrderAndReservationOperationResponse) LoadMore(ctx context.Context) (resp ReservationsSummariesListByReservationOrderAndReservationOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ReservationsSummariesListByReservationOrderAndReservationOperationOptions struct {
	Filter *string
	Grain  *Datagrain
}

func DefaultReservationsSummariesListByReservationOrderAndReservationOperationOptions() ReservationsSummariesListByReservationOrderAndReservationOperationOptions {
	return ReservationsSummariesListByReservationOrderAndReservationOperationOptions{}
}

func (o ReservationsSummariesListByReservationOrderAndReservationOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ReservationsSummariesListByReservationOrderAndReservationOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Grain != nil {
		out["grain"] = *o.Grain
	}

	return out
}

// ReservationsSummariesListByReservationOrderAndReservation ...
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderAndReservation(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions) (resp ReservationsSummariesListByReservationOrderAndReservationOperationResponse, err error) {
	req, err := c.preparerForReservationsSummariesListByReservationOrderAndReservation(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrderAndReservation", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrderAndReservation", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForReservationsSummariesListByReservationOrderAndReservation(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrderAndReservation", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForReservationsSummariesListByReservationOrderAndReservation prepares the ReservationsSummariesListByReservationOrderAndReservation request.
func (c ReservationSummariesClient) preparerForReservationsSummariesListByReservationOrderAndReservation(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions) (*http.Request, error) {
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

// preparerForReservationsSummariesListByReservationOrderAndReservationWithNextLink prepares the ReservationsSummariesListByReservationOrderAndReservation request with the given nextLink token.
func (c ReservationSummariesClient) preparerForReservationsSummariesListByReservationOrderAndReservationWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForReservationsSummariesListByReservationOrderAndReservation handles the response to the ReservationsSummariesListByReservationOrderAndReservation request. The method always
// closes the http.Response Body.
func (c ReservationSummariesClient) responderForReservationsSummariesListByReservationOrderAndReservation(resp *http.Response) (result ReservationsSummariesListByReservationOrderAndReservationOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ReservationsSummariesListByReservationOrderAndReservationOperationResponse, err error) {
			req, err := c.preparerForReservationsSummariesListByReservationOrderAndReservationWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrderAndReservation", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrderAndReservation", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForReservationsSummariesListByReservationOrderAndReservation(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesListByReservationOrderAndReservation", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ReservationsSummariesListByReservationOrderAndReservationComplete retrieves all of the results into a single object
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderAndReservationComplete(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions) (ReservationsSummariesListByReservationOrderAndReservationCompleteResult, error) {
	return c.ReservationsSummariesListByReservationOrderAndReservationCompleteMatchingPredicate(ctx, id, options, ReservationSummaryOperationPredicate{})
}

// ReservationsSummariesListByReservationOrderAndReservationCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReservationSummariesClient) ReservationsSummariesListByReservationOrderAndReservationCompleteMatchingPredicate(ctx context.Context, id ReservationId, options ReservationsSummariesListByReservationOrderAndReservationOperationOptions, predicate ReservationSummaryOperationPredicate) (resp ReservationsSummariesListByReservationOrderAndReservationCompleteResult, err error) {
	items := make([]ReservationSummary, 0)

	page, err := c.ReservationsSummariesListByReservationOrderAndReservation(ctx, id, options)
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

	out := ReservationsSummariesListByReservationOrderAndReservationCompleteResult{
		Items: items,
	}
	return out, nil
}
