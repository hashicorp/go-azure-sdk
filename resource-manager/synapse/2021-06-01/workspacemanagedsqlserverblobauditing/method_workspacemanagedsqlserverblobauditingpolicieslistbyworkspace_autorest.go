package workspacemanagedsqlserverblobauditing

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

type WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ServerBlobAuditingPolicy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse, error)
}

type WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult struct {
	Items []ServerBlobAuditingPolicy
}

func (r WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx context.Context, id WorkspaceId) (resp WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace prepares the WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace request.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/auditingSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceWithNextLink prepares the WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace request with the given nextLink token.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace handles the response to the WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerBlobAuditingClient) responderForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(resp *http.Response) (result WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	type page struct {
		Values   []ServerBlobAuditingPolicy `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
			req, err := c.preparerForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceComplete retrieves all of the results into a single object
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult, error) {
	return c.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx, id, ServerBlobAuditingPolicyOperationPredicate{})
}

// WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ServerBlobAuditingPolicyOperationPredicate) (resp WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult, err error) {
	items := make([]ServerBlobAuditingPolicy, 0)

	page, err := c.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx, id)
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

	out := WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceCompleteResult{
		Items: items,
	}
	return out, nil
}
