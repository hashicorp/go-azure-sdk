package authorizationaccesspolicy

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

type ListByAuthorizationOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]AuthorizationAccessPolicyContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByAuthorizationOperationResponse, error)
}

type ListByAuthorizationCompleteResult struct {
	Items []AuthorizationAccessPolicyContract
}

func (r ListByAuthorizationOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByAuthorizationOperationResponse) LoadMore(ctx context.Context) (resp ListByAuthorizationOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByAuthorizationOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByAuthorizationOperationOptions() ListByAuthorizationOperationOptions {
	return ListByAuthorizationOperationOptions{}
}

func (o ListByAuthorizationOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByAuthorizationOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByAuthorization ...
func (c AuthorizationAccessPolicyClient) ListByAuthorization(ctx context.Context, id AuthorizationId, options ListByAuthorizationOperationOptions) (resp ListByAuthorizationOperationResponse, err error) {
	req, err := c.preparerForListByAuthorization(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationaccesspolicy.AuthorizationAccessPolicyClient", "ListByAuthorization", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationaccesspolicy.AuthorizationAccessPolicyClient", "ListByAuthorization", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByAuthorization(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorizationaccesspolicy.AuthorizationAccessPolicyClient", "ListByAuthorization", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByAuthorization prepares the ListByAuthorization request.
func (c AuthorizationAccessPolicyClient) preparerForListByAuthorization(ctx context.Context, id AuthorizationId, options ListByAuthorizationOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/accessPolicies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByAuthorizationWithNextLink prepares the ListByAuthorization request with the given nextLink token.
func (c AuthorizationAccessPolicyClient) preparerForListByAuthorizationWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByAuthorization handles the response to the ListByAuthorization request. The method always
// closes the http.Response Body.
func (c AuthorizationAccessPolicyClient) responderForListByAuthorization(resp *http.Response) (result ListByAuthorizationOperationResponse, err error) {
	type page struct {
		Values   []AuthorizationAccessPolicyContract `json:"value"`
		NextLink *string                             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByAuthorizationOperationResponse, err error) {
			req, err := c.preparerForListByAuthorizationWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "authorizationaccesspolicy.AuthorizationAccessPolicyClient", "ListByAuthorization", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "authorizationaccesspolicy.AuthorizationAccessPolicyClient", "ListByAuthorization", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByAuthorization(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "authorizationaccesspolicy.AuthorizationAccessPolicyClient", "ListByAuthorization", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByAuthorizationComplete retrieves all of the results into a single object
func (c AuthorizationAccessPolicyClient) ListByAuthorizationComplete(ctx context.Context, id AuthorizationId, options ListByAuthorizationOperationOptions) (ListByAuthorizationCompleteResult, error) {
	return c.ListByAuthorizationCompleteMatchingPredicate(ctx, id, options, AuthorizationAccessPolicyContractOperationPredicate{})
}

// ListByAuthorizationCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AuthorizationAccessPolicyClient) ListByAuthorizationCompleteMatchingPredicate(ctx context.Context, id AuthorizationId, options ListByAuthorizationOperationOptions, predicate AuthorizationAccessPolicyContractOperationPredicate) (resp ListByAuthorizationCompleteResult, err error) {
	items := make([]AuthorizationAccessPolicyContract, 0)

	page, err := c.ListByAuthorization(ctx, id, options)
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

	out := ListByAuthorizationCompleteResult{
		Items: items,
	}
	return out, nil
}
