package recoverypoints

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

type ListByReplicationProtectedItemsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RecoveryPoint

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByReplicationProtectedItemsOperationResponse, error)
}

type ListByReplicationProtectedItemsCompleteResult struct {
	Items []RecoveryPoint
}

func (r ListByReplicationProtectedItemsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByReplicationProtectedItemsOperationResponse) LoadMore(ctx context.Context) (resp ListByReplicationProtectedItemsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByReplicationProtectedItems ...
func (c RecoveryPointsClient) ListByReplicationProtectedItems(ctx context.Context, id ReplicationProtectedItemId) (resp ListByReplicationProtectedItemsOperationResponse, err error) {
	req, err := c.preparerForListByReplicationProtectedItems(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypoints.RecoveryPointsClient", "ListByReplicationProtectedItems", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypoints.RecoveryPointsClient", "ListByReplicationProtectedItems", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByReplicationProtectedItems(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypoints.RecoveryPointsClient", "ListByReplicationProtectedItems", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByReplicationProtectedItems prepares the ListByReplicationProtectedItems request.
func (c RecoveryPointsClient) preparerForListByReplicationProtectedItems(ctx context.Context, id ReplicationProtectedItemId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recoveryPoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByReplicationProtectedItemsWithNextLink prepares the ListByReplicationProtectedItems request with the given nextLink token.
func (c RecoveryPointsClient) preparerForListByReplicationProtectedItemsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByReplicationProtectedItems handles the response to the ListByReplicationProtectedItems request. The method always
// closes the http.Response Body.
func (c RecoveryPointsClient) responderForListByReplicationProtectedItems(resp *http.Response) (result ListByReplicationProtectedItemsOperationResponse, err error) {
	type page struct {
		Values   []RecoveryPoint `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByReplicationProtectedItemsOperationResponse, err error) {
			req, err := c.preparerForListByReplicationProtectedItemsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverypoints.RecoveryPointsClient", "ListByReplicationProtectedItems", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverypoints.RecoveryPointsClient", "ListByReplicationProtectedItems", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByReplicationProtectedItems(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverypoints.RecoveryPointsClient", "ListByReplicationProtectedItems", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByReplicationProtectedItemsComplete retrieves all of the results into a single object
func (c RecoveryPointsClient) ListByReplicationProtectedItemsComplete(ctx context.Context, id ReplicationProtectedItemId) (ListByReplicationProtectedItemsCompleteResult, error) {
	return c.ListByReplicationProtectedItemsCompleteMatchingPredicate(ctx, id, RecoveryPointOperationPredicate{})
}

// ListByReplicationProtectedItemsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RecoveryPointsClient) ListByReplicationProtectedItemsCompleteMatchingPredicate(ctx context.Context, id ReplicationProtectedItemId, predicate RecoveryPointOperationPredicate) (resp ListByReplicationProtectedItemsCompleteResult, err error) {
	items := make([]RecoveryPoint, 0)

	page, err := c.ListByReplicationProtectedItems(ctx, id)
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

	out := ListByReplicationProtectedItemsCompleteResult{
		Items: items,
	}
	return out, nil
}
