package deployments

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

type ListAtTenantScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DeploymentExtended

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAtTenantScopeOperationResponse, error)
}

type ListAtTenantScopeCompleteResult struct {
	Items []DeploymentExtended
}

func (r ListAtTenantScopeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAtTenantScopeOperationResponse) LoadMore(ctx context.Context) (resp ListAtTenantScopeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListAtTenantScopeOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultListAtTenantScopeOperationOptions() ListAtTenantScopeOperationOptions {
	return ListAtTenantScopeOperationOptions{}
}

func (o ListAtTenantScopeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAtTenantScopeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListAtTenantScope ...
func (c DeploymentsClient) ListAtTenantScope(ctx context.Context, options ListAtTenantScopeOperationOptions) (resp ListAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForListAtTenantScope(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ListAtTenantScope", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ListAtTenantScope", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAtTenantScope(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ListAtTenantScope", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListAtTenantScope prepares the ListAtTenantScope request.
func (c DeploymentsClient) preparerForListAtTenantScope(ctx context.Context, options ListAtTenantScopeOperationOptions) (*http.Request, error) {
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
		autorest.WithPath("/providers/Microsoft.Resources/deployments"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAtTenantScopeWithNextLink prepares the ListAtTenantScope request with the given nextLink token.
func (c DeploymentsClient) preparerForListAtTenantScopeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAtTenantScope handles the response to the ListAtTenantScope request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForListAtTenantScope(resp *http.Response) (result ListAtTenantScopeOperationResponse, err error) {
	type page struct {
		Values   []DeploymentExtended `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAtTenantScopeOperationResponse, err error) {
			req, err := c.preparerForListAtTenantScopeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ListAtTenantScope", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ListAtTenantScope", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAtTenantScope(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ListAtTenantScope", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListAtTenantScopeComplete retrieves all of the results into a single object
func (c DeploymentsClient) ListAtTenantScopeComplete(ctx context.Context, options ListAtTenantScopeOperationOptions) (ListAtTenantScopeCompleteResult, error) {
	return c.ListAtTenantScopeCompleteMatchingPredicate(ctx, options, DeploymentExtendedOperationPredicate{})
}

// ListAtTenantScopeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DeploymentsClient) ListAtTenantScopeCompleteMatchingPredicate(ctx context.Context, options ListAtTenantScopeOperationOptions, predicate DeploymentExtendedOperationPredicate) (resp ListAtTenantScopeCompleteResult, err error) {
	items := make([]DeploymentExtended, 0)

	page, err := c.ListAtTenantScope(ctx, options)
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

	out := ListAtTenantScopeCompleteResult{
		Items: items,
	}
	return out, nil
}
