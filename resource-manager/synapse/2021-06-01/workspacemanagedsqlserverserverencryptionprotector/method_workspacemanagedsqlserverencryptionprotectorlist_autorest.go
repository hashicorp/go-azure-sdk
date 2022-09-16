package workspacemanagedsqlserverserverencryptionprotector

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

type WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]EncryptionProtector

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse, error)
}

type WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult struct {
	Items []EncryptionProtector
}

func (r WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse) LoadMore(ctx context.Context) (resp WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// WorkspaceManagedSqlServerEncryptionProtectorList ...
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorList(ctx context.Context, id WorkspaceId) (resp WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse, err error) {
	req, err := c.preparerForWorkspaceManagedSqlServerEncryptionProtectorList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWorkspaceManagedSqlServerEncryptionProtectorList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWorkspaceManagedSqlServerEncryptionProtectorList prepares the WorkspaceManagedSqlServerEncryptionProtectorList request.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) preparerForWorkspaceManagedSqlServerEncryptionProtectorList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/encryptionProtector", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWorkspaceManagedSqlServerEncryptionProtectorListWithNextLink prepares the WorkspaceManagedSqlServerEncryptionProtectorList request with the given nextLink token.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) preparerForWorkspaceManagedSqlServerEncryptionProtectorListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWorkspaceManagedSqlServerEncryptionProtectorList handles the response to the WorkspaceManagedSqlServerEncryptionProtectorList request. The method always
// closes the http.Response Body.
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) responderForWorkspaceManagedSqlServerEncryptionProtectorList(resp *http.Response) (result WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse, err error) {
	type page struct {
		Values   []EncryptionProtector `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WorkspaceManagedSqlServerEncryptionProtectorListOperationResponse, err error) {
			req, err := c.preparerForWorkspaceManagedSqlServerEncryptionProtectorListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWorkspaceManagedSqlServerEncryptionProtectorList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient", "WorkspaceManagedSqlServerEncryptionProtectorList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WorkspaceManagedSqlServerEncryptionProtectorListComplete retrieves all of the results into a single object
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorListComplete(ctx context.Context, id WorkspaceId) (WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult, error) {
	return c.WorkspaceManagedSqlServerEncryptionProtectorListCompleteMatchingPredicate(ctx, id, EncryptionProtectorOperationPredicate{})
}

// WorkspaceManagedSqlServerEncryptionProtectorListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkspaceManagedSqlServerServerEncryptionProtectorClient) WorkspaceManagedSqlServerEncryptionProtectorListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate EncryptionProtectorOperationPredicate) (resp WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult, err error) {
	items := make([]EncryptionProtector, 0)

	page, err := c.WorkspaceManagedSqlServerEncryptionProtectorList(ctx, id)
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

	out := WorkspaceManagedSqlServerEncryptionProtectorListCompleteResult{
		Items: items,
	}
	return out, nil
}
