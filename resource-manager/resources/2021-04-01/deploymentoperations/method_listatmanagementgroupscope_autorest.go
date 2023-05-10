package deploymentoperations

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

type ListAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DeploymentOperation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAtManagementGroupScopeOperationResponse, error)
}

type ListAtManagementGroupScopeCompleteResult struct {
	Items []DeploymentOperation
}

func (r ListAtManagementGroupScopeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAtManagementGroupScopeOperationResponse) LoadMore(ctx context.Context) (resp ListAtManagementGroupScopeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListAtManagementGroupScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtManagementGroupScopeOperationOptions() ListAtManagementGroupScopeOperationOptions {
	return ListAtManagementGroupScopeOperationOptions{}
}

func (o ListAtManagementGroupScopeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAtManagementGroupScopeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListAtManagementGroupScope ...
func (c DeploymentOperationsClient) ListAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions) (resp ListAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForListAtManagementGroupScope(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtManagementGroupScope", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAtManagementGroupScope(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtManagementGroupScope", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListAtManagementGroupScope prepares the ListAtManagementGroupScope request.
func (c DeploymentOperationsClient) preparerForListAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/operations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAtManagementGroupScopeWithNextLink prepares the ListAtManagementGroupScope request with the given nextLink token.
func (c DeploymentOperationsClient) preparerForListAtManagementGroupScopeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAtManagementGroupScope handles the response to the ListAtManagementGroupScope request. The method always
// closes the http.Response Body.
func (c DeploymentOperationsClient) responderForListAtManagementGroupScope(resp *http.Response) (result ListAtManagementGroupScopeOperationResponse, err error) {
	type page struct {
		Values   []DeploymentOperation `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAtManagementGroupScopeOperationResponse, err error) {
			req, err := c.preparerForListAtManagementGroupScopeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtManagementGroupScope", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtManagementGroupScope", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAtManagementGroupScope(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtManagementGroupScope", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListAtManagementGroupScopeComplete retrieves all of the results into a single object
func (c DeploymentOperationsClient) ListAtManagementGroupScopeComplete(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions) (ListAtManagementGroupScopeCompleteResult, error) {
	return c.ListAtManagementGroupScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtManagementGroupScopeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DeploymentOperationsClient) ListAtManagementGroupScopeCompleteMatchingPredicate(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (resp ListAtManagementGroupScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	page, err := c.ListAtManagementGroupScope(ctx, id, options)
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

	out := ListAtManagementGroupScopeCompleteResult{
		Items: items,
	}
	return out, nil
}
