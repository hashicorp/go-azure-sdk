package sqlpoolsworkloadgroups

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

type SqlPoolWorkloadGroupListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadGroup

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolWorkloadGroupListOperationResponse, error)
}

type SqlPoolWorkloadGroupListCompleteResult struct {
	Items []WorkloadGroup
}

func (r SqlPoolWorkloadGroupListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolWorkloadGroupListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolWorkloadGroupListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolWorkloadGroupList ...
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupList(ctx context.Context, id SqlPoolId) (resp SqlPoolWorkloadGroupListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolWorkloadGroupList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolWorkloadGroupList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolWorkloadGroupList prepares the SqlPoolWorkloadGroupList request.
func (c SqlPoolsWorkloadGroupsClient) preparerForSqlPoolWorkloadGroupList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadGroups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolWorkloadGroupListWithNextLink prepares the SqlPoolWorkloadGroupList request with the given nextLink token.
func (c SqlPoolsWorkloadGroupsClient) preparerForSqlPoolWorkloadGroupListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForSqlPoolWorkloadGroupList handles the response to the SqlPoolWorkloadGroupList request. The method always
// closes the http.Response Body.
func (c SqlPoolsWorkloadGroupsClient) responderForSqlPoolWorkloadGroupList(resp *http.Response) (result SqlPoolWorkloadGroupListOperationResponse, err error) {
	type page struct {
		Values   []WorkloadGroup `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolWorkloadGroupListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolWorkloadGroupListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolWorkloadGroupList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient", "SqlPoolWorkloadGroupList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolWorkloadGroupListComplete retrieves all of the results into a single object
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupListComplete(ctx context.Context, id SqlPoolId) (SqlPoolWorkloadGroupListCompleteResult, error) {
	return c.SqlPoolWorkloadGroupListCompleteMatchingPredicate(ctx, id, WorkloadGroupOperationPredicate{})
}

// SqlPoolWorkloadGroupListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate WorkloadGroupOperationPredicate) (resp SqlPoolWorkloadGroupListCompleteResult, err error) {
	items := make([]WorkloadGroup, 0)

	page, err := c.SqlPoolWorkloadGroupList(ctx, id)
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

	out := SqlPoolWorkloadGroupListCompleteResult{
		Items: items,
	}
	return out, nil
}
