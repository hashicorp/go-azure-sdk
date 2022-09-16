package workspacemanagedsqlserver

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

type RecoverableSqlPoolsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RecoverableSqlPool

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RecoverableSqlPoolsListOperationResponse, error)
}

type RecoverableSqlPoolsListCompleteResult struct {
	Items []RecoverableSqlPool
}

func (r RecoverableSqlPoolsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RecoverableSqlPoolsListOperationResponse) LoadMore(ctx context.Context) (resp RecoverableSqlPoolsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// RecoverableSqlPoolsList ...
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsList(ctx context.Context, id WorkspaceId) (resp RecoverableSqlPoolsListOperationResponse, err error) {
	req, err := c.preparerForRecoverableSqlPoolsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRecoverableSqlPoolsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRecoverableSqlPoolsList prepares the RecoverableSqlPoolsList request.
func (c WorkspaceManagedSqlServerClient) preparerForRecoverableSqlPoolsList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
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

// preparerForRecoverableSqlPoolsListWithNextLink prepares the RecoverableSqlPoolsList request with the given nextLink token.
func (c WorkspaceManagedSqlServerClient) preparerForRecoverableSqlPoolsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRecoverableSqlPoolsList handles the response to the RecoverableSqlPoolsList request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerClient) responderForRecoverableSqlPoolsList(resp *http.Response) (result RecoverableSqlPoolsListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RecoverableSqlPoolsListOperationResponse, err error) {
			req, err := c.preparerForRecoverableSqlPoolsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRecoverableSqlPoolsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserver.WorkspaceManagedSqlServerClient", "RecoverableSqlPoolsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RecoverableSqlPoolsListComplete retrieves all of the results into a single object
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsListComplete(ctx context.Context, id WorkspaceId) (RecoverableSqlPoolsListCompleteResult, error) {
	return c.RecoverableSqlPoolsListCompleteMatchingPredicate(ctx, id, RecoverableSqlPoolOperationPredicate{})
}

// RecoverableSqlPoolsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceManagedSqlServerClient) RecoverableSqlPoolsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate RecoverableSqlPoolOperationPredicate) (resp RecoverableSqlPoolsListCompleteResult, err error) {
	items := make([]RecoverableSqlPool, 0)

	page, err := c.RecoverableSqlPoolsList(ctx, id)
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

	out := RecoverableSqlPoolsListCompleteResult{
		Items: items,
	}
	return out, nil
}
