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

type ListAtScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DeploymentOperation

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAtScopeOperationResponse, error)
}

type ListAtScopeCompleteResult struct {
	Items []DeploymentOperation
}

func (r ListAtScopeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAtScopeOperationResponse) LoadMore(ctx context.Context) (resp ListAtScopeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListAtScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtScopeOperationOptions() ListAtScopeOperationOptions {
	return ListAtScopeOperationOptions{}
}

func (o ListAtScopeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAtScopeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListAtScope ...
func (c DeploymentOperationsClient) ListAtScope(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions) (resp ListAtScopeOperationResponse, err error) {
	req, err := c.preparerForListAtScope(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtScope", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtScope", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAtScope(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtScope", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListAtScope prepares the ListAtScope request.
func (c DeploymentOperationsClient) preparerForListAtScope(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions) (*http.Request, error) {
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

// preparerForListAtScopeWithNextLink prepares the ListAtScope request with the given nextLink token.
func (c DeploymentOperationsClient) preparerForListAtScopeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAtScope handles the response to the ListAtScope request. The method always
// closes the http.Response Body.
func (c DeploymentOperationsClient) responderForListAtScope(resp *http.Response) (result ListAtScopeOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAtScopeOperationResponse, err error) {
			req, err := c.preparerForListAtScopeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtScope", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtScope", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAtScope(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deploymentoperations.DeploymentOperationsClient", "ListAtScope", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListAtScopeComplete retrieves all of the results into a single object
func (c DeploymentOperationsClient) ListAtScopeComplete(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions) (ListAtScopeCompleteResult, error) {
	return c.ListAtScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtScopeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DeploymentOperationsClient) ListAtScopeCompleteMatchingPredicate(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (resp ListAtScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	page, err := c.ListAtScope(ctx, id, options)
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

	out := ListAtScopeCompleteResult{
		Items: items,
	}
	return out, nil
}
