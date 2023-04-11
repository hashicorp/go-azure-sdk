package jitnetworkaccesspolicies

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

type ListByResourceGroupAndRegionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]JitNetworkAccessPolicy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByResourceGroupAndRegionOperationResponse, error)
}

type ListByResourceGroupAndRegionCompleteResult struct {
	Items []JitNetworkAccessPolicy
}

func (r ListByResourceGroupAndRegionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByResourceGroupAndRegionOperationResponse) LoadMore(ctx context.Context) (resp ListByResourceGroupAndRegionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByResourceGroupAndRegion ...
func (c JitNetworkAccessPoliciesClient) ListByResourceGroupAndRegion(ctx context.Context, id ProviderLocationId) (resp ListByResourceGroupAndRegionOperationResponse, err error) {
	req, err := c.preparerForListByResourceGroupAndRegion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByResourceGroupAndRegion", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByResourceGroupAndRegion", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByResourceGroupAndRegion(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByResourceGroupAndRegion", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByResourceGroupAndRegion prepares the ListByResourceGroupAndRegion request.
func (c JitNetworkAccessPoliciesClient) preparerForListByResourceGroupAndRegion(ctx context.Context, id ProviderLocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/jitNetworkAccessPolicies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByResourceGroupAndRegionWithNextLink prepares the ListByResourceGroupAndRegion request with the given nextLink token.
func (c JitNetworkAccessPoliciesClient) preparerForListByResourceGroupAndRegionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByResourceGroupAndRegion handles the response to the ListByResourceGroupAndRegion request. The method always
// closes the http.Response Body.
func (c JitNetworkAccessPoliciesClient) responderForListByResourceGroupAndRegion(resp *http.Response) (result ListByResourceGroupAndRegionOperationResponse, err error) {
	type page struct {
		Values   []JitNetworkAccessPolicy `json:"value"`
		NextLink *string                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByResourceGroupAndRegionOperationResponse, err error) {
			req, err := c.preparerForListByResourceGroupAndRegionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByResourceGroupAndRegion", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByResourceGroupAndRegion", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByResourceGroupAndRegion(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByResourceGroupAndRegion", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByResourceGroupAndRegionComplete retrieves all of the results into a single object
func (c JitNetworkAccessPoliciesClient) ListByResourceGroupAndRegionComplete(ctx context.Context, id ProviderLocationId) (ListByResourceGroupAndRegionCompleteResult, error) {
	return c.ListByResourceGroupAndRegionCompleteMatchingPredicate(ctx, id, JitNetworkAccessPolicyOperationPredicate{})
}

// ListByResourceGroupAndRegionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c JitNetworkAccessPoliciesClient) ListByResourceGroupAndRegionCompleteMatchingPredicate(ctx context.Context, id ProviderLocationId, predicate JitNetworkAccessPolicyOperationPredicate) (resp ListByResourceGroupAndRegionCompleteResult, err error) {
	items := make([]JitNetworkAccessPolicy, 0)

	page, err := c.ListByResourceGroupAndRegion(ctx, id)
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

	out := ListByResourceGroupAndRegionCompleteResult{
		Items: items,
	}
	return out, nil
}
