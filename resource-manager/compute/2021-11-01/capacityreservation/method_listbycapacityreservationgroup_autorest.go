package capacityreservation

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

type ListByCapacityReservationGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CapacityReservation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByCapacityReservationGroupOperationResponse, error)
}

type ListByCapacityReservationGroupCompleteResult struct {
	Items []CapacityReservation
}

func (r ListByCapacityReservationGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByCapacityReservationGroupOperationResponse) LoadMore(ctx context.Context) (resp ListByCapacityReservationGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByCapacityReservationGroup ...
func (c CapacityReservationClient) ListByCapacityReservationGroup(ctx context.Context, id CapacityReservationGroupId) (resp ListByCapacityReservationGroupOperationResponse, err error) {
	req, err := c.preparerForListByCapacityReservationGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "capacityreservation.CapacityReservationClient", "ListByCapacityReservationGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "capacityreservation.CapacityReservationClient", "ListByCapacityReservationGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByCapacityReservationGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "capacityreservation.CapacityReservationClient", "ListByCapacityReservationGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListByCapacityReservationGroupComplete retrieves all of the results into a single object
func (c CapacityReservationClient) ListByCapacityReservationGroupComplete(ctx context.Context, id CapacityReservationGroupId) (ListByCapacityReservationGroupCompleteResult, error) {
	return c.ListByCapacityReservationGroupCompleteMatchingPredicate(ctx, id, CapacityReservationOperationPredicate{})
}

// ListByCapacityReservationGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c CapacityReservationClient) ListByCapacityReservationGroupCompleteMatchingPredicate(ctx context.Context, id CapacityReservationGroupId, predicate CapacityReservationOperationPredicate) (resp ListByCapacityReservationGroupCompleteResult, err error) {
	items := make([]CapacityReservation, 0)

	page, err := c.ListByCapacityReservationGroup(ctx, id)
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

	out := ListByCapacityReservationGroupCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListByCapacityReservationGroup prepares the ListByCapacityReservationGroup request.
func (c CapacityReservationClient) preparerForListByCapacityReservationGroup(ctx context.Context, id CapacityReservationGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/capacityReservations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByCapacityReservationGroupWithNextLink prepares the ListByCapacityReservationGroup request with the given nextLink token.
func (c CapacityReservationClient) preparerForListByCapacityReservationGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByCapacityReservationGroup handles the response to the ListByCapacityReservationGroup request. The method always
// closes the http.Response Body.
func (c CapacityReservationClient) responderForListByCapacityReservationGroup(resp *http.Response) (result ListByCapacityReservationGroupOperationResponse, err error) {
	type page struct {
		Values   []CapacityReservation `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByCapacityReservationGroupOperationResponse, err error) {
			req, err := c.preparerForListByCapacityReservationGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "capacityreservation.CapacityReservationClient", "ListByCapacityReservationGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "capacityreservation.CapacityReservationClient", "ListByCapacityReservationGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByCapacityReservationGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "capacityreservation.CapacityReservationClient", "ListByCapacityReservationGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
