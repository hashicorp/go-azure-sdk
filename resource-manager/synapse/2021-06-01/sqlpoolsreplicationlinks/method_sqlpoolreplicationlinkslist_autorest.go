package sqlpoolsreplicationlinks

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

type SqlPoolReplicationLinksListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ReplicationLink

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolReplicationLinksListOperationResponse, error)
}

type SqlPoolReplicationLinksListCompleteResult struct {
	Items []ReplicationLink
}

func (r SqlPoolReplicationLinksListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolReplicationLinksListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolReplicationLinksListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolReplicationLinksList ...
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksList(ctx context.Context, id SqlPoolId) (resp SqlPoolReplicationLinksListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolReplicationLinksList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolReplicationLinksList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolReplicationLinksList prepares the SqlPoolReplicationLinksList request.
func (c SqlPoolsReplicationLinksClient) preparerForSqlPoolReplicationLinksList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/replicationLinks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolReplicationLinksListWithNextLink prepares the SqlPoolReplicationLinksList request with the given nextLink token.
func (c SqlPoolsReplicationLinksClient) preparerForSqlPoolReplicationLinksListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolReplicationLinksList handles the response to the SqlPoolReplicationLinksList request. The method always
// closes the http.Response Body.
func (c SqlPoolsReplicationLinksClient) responderForSqlPoolReplicationLinksList(resp *http.Response) (result SqlPoolReplicationLinksListOperationResponse, err error) {
	type page struct {
		Values   []ReplicationLink `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolReplicationLinksListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolReplicationLinksListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolReplicationLinksList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolReplicationLinksListComplete retrieves all of the results into a single object
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksListComplete(ctx context.Context, id SqlPoolId) (SqlPoolReplicationLinksListCompleteResult, error) {
	return c.SqlPoolReplicationLinksListCompleteMatchingPredicate(ctx, id, ReplicationLinkOperationPredicate{})
}

// SqlPoolReplicationLinksListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate ReplicationLinkOperationPredicate) (resp SqlPoolReplicationLinksListCompleteResult, err error) {
	items := make([]ReplicationLink, 0)

	page, err := c.SqlPoolReplicationLinksList(ctx, id)
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

	out := SqlPoolReplicationLinksListCompleteResult{
		Items: items,
	}
	return out, nil
}
