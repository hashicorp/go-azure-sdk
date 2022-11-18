package reservationdetails

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationsDetailsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReservationDetail

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ReservationsDetailsListOperationResponse, error)
}

type ReservationsDetailsListCompleteResult struct {
	Items []ReservationDetail
}

func (r ReservationsDetailsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ReservationsDetailsListOperationResponse) LoadMore(ctx context.Context) (resp ReservationsDetailsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ReservationsDetailsListOperationOptions struct {
	EndDate            *string
	Filter             *string
	ReservationId      *string
	ReservationOrderId *string
	StartDate          *string
}

func DefaultReservationsDetailsListOperationOptions() ReservationsDetailsListOperationOptions {
	return ReservationsDetailsListOperationOptions{}
}

func (o ReservationsDetailsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ReservationsDetailsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndDate != nil {
		out["endDate"] = *o.EndDate
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.ReservationId != nil {
		out["reservationId"] = *o.ReservationId
	}

	if o.ReservationOrderId != nil {
		out["reservationOrderId"] = *o.ReservationOrderId
	}

	if o.StartDate != nil {
		out["startDate"] = *o.StartDate
	}

	return out
}

// ReservationsDetailsList ...
func (c ReservationDetailsClient) ReservationsDetailsList(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions) (resp ReservationsDetailsListOperationResponse, err error) {
	req, err := c.preparerForReservationsDetailsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForReservationsDetailsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForReservationsDetailsList prepares the ReservationsDetailsList request.
func (c ReservationDetailsClient) preparerForReservationsDetailsList(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions) (*http.Request, error) {
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

// preparerForReservationsDetailsListWithNextLink prepares the ReservationsDetailsList request with the given nextLink token.
func (c ReservationDetailsClient) preparerForReservationsDetailsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForReservationsDetailsList handles the response to the ReservationsDetailsList request. The method always
// closes the http.Response Body.
func (c ReservationDetailsClient) responderForReservationsDetailsList(resp *http.Response) (result ReservationsDetailsListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ReservationsDetailsListOperationResponse, err error) {
			req, err := c.preparerForReservationsDetailsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForReservationsDetailsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationdetails.ReservationDetailsClient", "ReservationsDetailsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ReservationsDetailsListComplete retrieves all of the results into a single object
func (c ReservationDetailsClient) ReservationsDetailsListComplete(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions) (ReservationsDetailsListCompleteResult, error) {
	return c.ReservationsDetailsListCompleteMatchingPredicate(ctx, id, options, ReservationDetailOperationPredicate{})
}

// ReservationsDetailsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReservationDetailsClient) ReservationsDetailsListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ReservationsDetailsListOperationOptions, predicate ReservationDetailOperationPredicate) (resp ReservationsDetailsListCompleteResult, err error) {
	items := make([]ReservationDetail, 0)

	page, err := c.ReservationsDetailsList(ctx, id, options)
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

	out := ReservationsDetailsListCompleteResult{
		Items: items,
	}
	return out, nil
}
