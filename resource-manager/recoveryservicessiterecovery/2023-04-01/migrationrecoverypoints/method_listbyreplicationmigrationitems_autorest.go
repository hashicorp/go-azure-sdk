package migrationrecoverypoints

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

type ListByReplicationMigrationItemsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]MigrationRecoveryPoint

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByReplicationMigrationItemsOperationResponse, error)
}

type ListByReplicationMigrationItemsCompleteResult struct {
	Items []MigrationRecoveryPoint
}

func (r ListByReplicationMigrationItemsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByReplicationMigrationItemsOperationResponse) LoadMore(ctx context.Context) (resp ListByReplicationMigrationItemsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByReplicationMigrationItems ...
func (c MigrationRecoveryPointsClient) ListByReplicationMigrationItems(ctx context.Context, id ReplicationMigrationItemId) (resp ListByReplicationMigrationItemsOperationResponse, err error) {
	req, err := c.preparerForListByReplicationMigrationItems(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationrecoverypoints.MigrationRecoveryPointsClient", "ListByReplicationMigrationItems", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationrecoverypoints.MigrationRecoveryPointsClient", "ListByReplicationMigrationItems", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByReplicationMigrationItems(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrationrecoverypoints.MigrationRecoveryPointsClient", "ListByReplicationMigrationItems", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByReplicationMigrationItems prepares the ListByReplicationMigrationItems request.
func (c MigrationRecoveryPointsClient) preparerForListByReplicationMigrationItems(ctx context.Context, id ReplicationMigrationItemId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrationRecoveryPoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByReplicationMigrationItemsWithNextLink prepares the ListByReplicationMigrationItems request with the given nextLink token.
func (c MigrationRecoveryPointsClient) preparerForListByReplicationMigrationItemsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByReplicationMigrationItems handles the response to the ListByReplicationMigrationItems request. The method always
// closes the http.Response Body.
func (c MigrationRecoveryPointsClient) responderForListByReplicationMigrationItems(resp *http.Response) (result ListByReplicationMigrationItemsOperationResponse, err error) {
	type page struct {
		Values   []MigrationRecoveryPoint `json:"value"`
		NextLink *string                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByReplicationMigrationItemsOperationResponse, err error) {
			req, err := c.preparerForListByReplicationMigrationItemsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrationrecoverypoints.MigrationRecoveryPointsClient", "ListByReplicationMigrationItems", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrationrecoverypoints.MigrationRecoveryPointsClient", "ListByReplicationMigrationItems", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByReplicationMigrationItems(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "migrationrecoverypoints.MigrationRecoveryPointsClient", "ListByReplicationMigrationItems", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByReplicationMigrationItemsComplete retrieves all of the results into a single object
func (c MigrationRecoveryPointsClient) ListByReplicationMigrationItemsComplete(ctx context.Context, id ReplicationMigrationItemId) (ListByReplicationMigrationItemsCompleteResult, error) {
	return c.ListByReplicationMigrationItemsCompleteMatchingPredicate(ctx, id, MigrationRecoveryPointOperationPredicate{})
}

// ListByReplicationMigrationItemsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MigrationRecoveryPointsClient) ListByReplicationMigrationItemsCompleteMatchingPredicate(ctx context.Context, id ReplicationMigrationItemId, predicate MigrationRecoveryPointOperationPredicate) (resp ListByReplicationMigrationItemsCompleteResult, err error) {
	items := make([]MigrationRecoveryPoint, 0)

	page, err := c.ListByReplicationMigrationItems(ctx, id)
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

	out := ListByReplicationMigrationItemsCompleteResult{
		Items: items,
	}
	return out, nil
}
