package recoverablesqlpools

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

type WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RecoverableSqlPool

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse, error)
}

type WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult struct {
	Items []RecoverableSqlPool
}

func (r WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsList ...
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx context.Context, id WorkspaceId) (resp WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceManagedSqlServerRecoverableSqlPoolsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsList prepares the WorkspaceManagedSqlServerRecoverableSqlPoolsList request.
func (c RecoverableSqlPoolsClient) preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recoverableSqlPools", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsListWithNextLink prepares the WorkspaceManagedSqlServerRecoverableSqlPoolsList request with the given nextLink token.
func (c RecoverableSqlPoolsClient) preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerRecoverableSqlPoolsList handles the response to the WorkspaceManagedSqlServerRecoverableSqlPoolsList request. The method always
// closes the http.Response Body.
func (c RecoverableSqlPoolsClient) responderForWorkspaceManagedSqlServerRecoverableSqlPoolsList(resp *http.Response) (result WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse, err error) {
	type page struct {
		Values   []RecoverableSqlPool `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceManagedSqlServerRecoverableSqlPoolsListOperationResponse, err error) {
			req, err := c.preparerForWorkspaceManagedSqlServerRecoverableSqlPoolsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceManagedSqlServerRecoverableSqlPoolsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recoverablesqlpools.RecoverableSqlPoolsClient", "WorkspaceManagedSqlServerRecoverableSqlPoolsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsListComplete retrieves all of the results into a single object
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteMatchingPredicate(ctx, id, RecoverableSqlPoolOperationPredicate{})
}

// WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RecoverableSqlPoolsClient) WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate RecoverableSqlPoolOperationPredicate) (resp WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult, err error) {
	items := make([]RecoverableSqlPool, 0)

	page, err := c.WorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx, id)
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

	out := WorkspaceManagedSqlServerRecoverableSqlPoolsListCompleteResult{
		Items: items,
	}
	return out, nil
}
