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

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ExtendedServerBlobAuditingPolicy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse, error)
}

type WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult struct {
	Items []ExtendedServerBlobAuditingPolicy
}

func (r WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace ...
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx context.Context, id WorkspaceId) (resp WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace prepares the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace request.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extendedAuditingSettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceWithNextLink prepares the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace request with the given nextLink token.
func (c WorkspaceManagedSqlServerBlobAuditingClient) preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace handles the response to the WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerBlobAuditingClient) responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(resp *http.Response) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
	type page struct {
		Values   []ExtendedServerBlobAuditingPolicy `json:"value"`
		NextLink *string                            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceOperationResponse, err error) {
			req, err := c.preparerForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient", "WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceComplete retrieves all of the results into a single object
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult, error) {
	return c.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx, id, ExtendedServerBlobAuditingPolicyOperationPredicate{})
}

// WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceManagedSqlServerBlobAuditingClient) WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate ExtendedServerBlobAuditingPolicyOperationPredicate) (resp WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult, err error) {
	items := make([]ExtendedServerBlobAuditingPolicy, 0)

	page, err := c.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx, id)
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

	out := WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceCompleteResult{
		Items: items,
	}
	return out, nil
}
