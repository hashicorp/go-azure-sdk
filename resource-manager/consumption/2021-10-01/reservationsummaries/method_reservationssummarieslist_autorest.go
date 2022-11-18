package reservationsummaries

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

type ReservationsSummariesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReservationSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ReservationsSummariesListOperationResponse, error)
}

type ReservationsSummariesListCompleteResult struct {
	Items []ReservationSummary
}

func (r ReservationsSummariesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ReservationsSummariesListOperationResponse) LoadMore(ctx context.Context) (resp ReservationsSummariesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ReservationsSummariesListOperationOptions struct {
	EndDate            *string
	Filter             *string
	Grain              *Datagrain
	ReservationId      *string
	ReservationOrderId *string
	StartDate          *string
}

func DefaultReservationsSummariesListOperationOptions() ReservationsSummariesListOperationOptions {
	return ReservationsSummariesListOperationOptions{}
}

func (o ReservationsSummariesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ReservationsSummariesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndDate != nil {
		out["endDate"] = *o.EndDate
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Grain != nil {
		out["grain"] = *o.Grain
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

// ReservationsSummariesList ...
func (c ReservationSummariesClient) ReservationsSummariesList(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions) (resp ReservationsSummariesListOperationResponse, err error) {
	req, err := c.preparerForReservationsSummariesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForReservationsSummariesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForReservationsSummariesList prepares the ReservationsSummariesList request.
func (c ReservationSummariesClient) preparerForReservationsSummariesList(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions) (*http.Request, error) {
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

// preparerForReservationsSummariesListWithNextLink prepares the ReservationsSummariesList request with the given nextLink token.
func (c ReservationSummariesClient) preparerForReservationsSummariesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForReservationsSummariesList handles the response to the ReservationsSummariesList request. The method always
// closes the http.Response Body.
func (c ReservationSummariesClient) responderForReservationsSummariesList(resp *http.Response) (result ReservationsSummariesListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ReservationsSummariesListOperationResponse, err error) {
			req, err := c.preparerForReservationsSummariesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForReservationsSummariesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservationsummaries.ReservationSummariesClient", "ReservationsSummariesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ReservationsSummariesListComplete retrieves all of the results into a single object
func (c ReservationSummariesClient) ReservationsSummariesListComplete(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions) (ReservationsSummariesListCompleteResult, error) {
	return c.ReservationsSummariesListCompleteMatchingPredicate(ctx, id, options, ReservationSummaryOperationPredicate{})
}

// ReservationsSummariesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReservationSummariesClient) ReservationsSummariesListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, options ReservationsSummariesListOperationOptions, predicate ReservationSummaryOperationPredicate) (resp ReservationsSummariesListCompleteResult, err error) {
	items := make([]ReservationSummary, 0)

	page, err := c.ReservationsSummariesList(ctx, id, options)
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

	out := ReservationsSummariesListCompleteResult{
		Items: items,
	}
	return out, nil
}
