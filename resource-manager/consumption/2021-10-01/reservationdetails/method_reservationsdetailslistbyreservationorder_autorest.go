package reservationdetails

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

type ReservationsDetailsListByReservationOrderOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReservationDetail

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ReservationsDetailsListByReservationOrderOperationResponse, error)
}

type ReservationsDetailsListByReservationOrderCompleteResult struct {
	Items []ReservationDetail
}

func (r ReservationsDetailsListByReservationOrderOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ReservationsDetailsListByReservationOrderOperationResponse) LoadMore(ctx context.Context) (resp ReservationsDetailsListByReservationOrderOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ReservationsDetailsListByReservationOrderOperationOptions struct {
	Filter *string
}

func DefaultReservationsDetailsListByReservationOrderOperationOptions() ReservationsDetailsListByReservationOrderOperationOptions {
	return ReservationsDetailsListByReservationOrderOperationOptions{}
}

func (o ReservationsDetailsListByReservationOrderOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ReservationsDetailsListByReservationOrderOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ReservationsDetailsListByReservationOrder ...
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrder(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions) (resp ReservationsDetailsListByReservationOrderOperationResponse, err error) {
	req, err := c.preparerForReservationsDetailsListByReservationOrder(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsListByReservationOrder", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsListByReservationOrder", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForReservationsDetailsListByReservationOrder(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsListByReservationOrder", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForReservationsDetailsListByReservationOrder prepares the ReservationsDetailsListByReservationOrder request.
func (c ReservationDetailsClient) preparerForReservationsDetailsListByReservationOrder(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/reservationDetails", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForReservationsDetailsListByReservationOrderWithNextLink prepares the ReservationsDetailsListByReservationOrder request with the given nextLink token.
func (c ReservationDetailsClient) preparerForReservationsDetailsListByReservationOrderWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForReservationsDetailsListByReservationOrder handles the response to the ReservationsDetailsListByReservationOrder request. The method always
// closes the http.Response Body.
func (c ReservationDetailsClient) responderForReservationsDetailsListByReservationOrder(resp *http.Response) (result ReservationsDetailsListByReservationOrderOperationResponse, err error) {
	type page struct {
		Values   []ReservationDetail `json:"value"`
		NextLink *string             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ReservationsDetailsListByReservationOrderOperationResponse, err error) {
			req, err := c.preparerForReservationsDetailsListByReservationOrderWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsListByReservationOrder", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsListByReservationOrder", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForReservationsDetailsListByReservationOrder(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsListByReservationOrder", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ReservationsDetailsListByReservationOrderComplete retrieves all of the results into a single object
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderComplete(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions) (ReservationsDetailsListByReservationOrderCompleteResult, error) {
	return c.ReservationsDetailsListByReservationOrderCompleteMatchingPredicate(ctx, id, options, ReservationDetailOperationPredicate{})
}

// ReservationsDetailsListByReservationOrderCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReservationDetailsClient) ReservationsDetailsListByReservationOrderCompleteMatchingPredicate(ctx context.Context, id ReservationOrderId, options ReservationsDetailsListByReservationOrderOperationOptions, predicate ReservationDetailOperationPredicate) (resp ReservationsDetailsListByReservationOrderCompleteResult, err error) {
	items := make([]ReservationDetail, 0)

	page, err := c.ReservationsDetailsListByReservationOrder(ctx, id, options)
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

	out := ReservationsDetailsListByReservationOrderCompleteResult{
		Items: items,
	}
	return out, nil
}
