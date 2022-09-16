package workspacemanagedsqlserversqlusages

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

type WorkspaceManagedSqlServerUsagesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ServerUsage

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceManagedSqlServerUsagesListOperationResponse, error)
}

type WorkspaceManagedSqlServerUsagesListCompleteResult struct {
	Items []ServerUsage
}

func (r WorkspaceManagedSqlServerUsagesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceManagedSqlServerUsagesListOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceManagedSqlServerUsagesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceManagedSqlServerUsagesList ...
func (c WorkspaceManagedSqlServerSqlUsagesClient) WorkspaceManagedSqlServerUsagesList(ctx context.Context, id WorkspaceId) (resp WorkspaceManagedSqlServerUsagesListOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerUsagesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient", "WorkspaceManagedSqlServerUsagesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient", "WorkspaceManagedSqlServerUsagesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceManagedSqlServerUsagesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient", "WorkspaceManagedSqlServerUsagesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceManagedSqlServerUsagesList prepares the WorkspaceManagedSqlServerUsagesList request.
func (c WorkspaceManagedSqlServerSqlUsagesClient) preparerForWorkspaceManagedSqlServerUsagesList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sqlUsages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceManagedSqlServerUsagesListWithNextLink prepares the WorkspaceManagedSqlServerUsagesList request with the given nextLink token.
func (c WorkspaceManagedSqlServerSqlUsagesClient) preparerForWorkspaceManagedSqlServerUsagesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerUsagesList handles the response to the WorkspaceManagedSqlServerUsagesList request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerSqlUsagesClient) responderForWorkspaceManagedSqlServerUsagesList(resp *http.Response) (result WorkspaceManagedSqlServerUsagesListOperationResponse, err error) {
	type page struct {
		Values   []ServerUsage `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceManagedSqlServerUsagesListOperationResponse, err error) {
			req, err := c.preparerForWorkspaceManagedSqlServerUsagesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient", "WorkspaceManagedSqlServerUsagesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient", "WorkspaceManagedSqlServerUsagesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceManagedSqlServerUsagesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient", "WorkspaceManagedSqlServerUsagesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceManagedSqlServerUsagesListComplete retrieves all of the results into a single object
func (c WorkspaceManagedSqlServerSqlUsagesClient) WorkspaceManagedSqlServerUsagesListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerUsagesListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerUsagesListCompleteMatchingPredicate(ctx, id, ServerUsageOperationPredicate{})
}

// WorkspaceManagedSqlServerUsagesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceManagedSqlServerSqlUsagesClient) WorkspaceManagedSqlServerUsagesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ServerUsageOperationPredicate) (resp WorkspaceManagedSqlServerUsagesListCompleteResult, err error) {
	items := make([]ServerUsage, 0)

	page, err := c.WorkspaceManagedSqlServerUsagesList(ctx, id)
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

	out := WorkspaceManagedSqlServerUsagesListCompleteResult{
		Items: items,
	}
	return out, nil
}
