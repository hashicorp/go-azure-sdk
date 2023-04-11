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

type ListByRegionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]JitNetworkAccessPolicy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByRegionOperationResponse, error)
}

type ListByRegionCompleteResult struct {
	Items []JitNetworkAccessPolicy
}

func (r ListByRegionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByRegionOperationResponse) LoadMore(ctx context.Context) (resp ListByRegionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByRegion ...
func (c JitNetworkAccessPoliciesClient) ListByRegion(ctx context.Context, id LocationId) (resp ListByRegionOperationResponse, err error) {
	req, err := c.preparerForListByRegion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByRegion", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByRegion", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByRegion(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByRegion", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByRegion prepares the ListByRegion request.
func (c JitNetworkAccessPoliciesClient) preparerForListByRegion(ctx context.Context, id LocationId) (*http.Request, error) {
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

// preparerForListByRegionWithNextLink prepares the ListByRegion request with the given nextLink token.
func (c JitNetworkAccessPoliciesClient) preparerForListByRegionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByRegion handles the response to the ListByRegion request. The method always
// closes the http.Response Body.
func (c JitNetworkAccessPoliciesClient) responderForListByRegion(resp *http.Response) (result ListByRegionOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByRegionOperationResponse, err error) {
			req, err := c.preparerForListByRegionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByRegion", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByRegion", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByRegion(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "ListByRegion", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByRegionComplete retrieves all of the results into a single object
func (c JitNetworkAccessPoliciesClient) ListByRegionComplete(ctx context.Context, id LocationId) (ListByRegionCompleteResult, error) {
	return c.ListByRegionCompleteMatchingPredicate(ctx, id, JitNetworkAccessPolicyOperationPredicate{})
}

// ListByRegionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c JitNetworkAccessPoliciesClient) ListByRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate JitNetworkAccessPolicyOperationPredicate) (resp ListByRegionCompleteResult, err error) {
	items := make([]JitNetworkAccessPolicy, 0)

	page, err := c.ListByRegion(ctx, id)
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

	out := ListByRegionCompleteResult{
		Items: items,
	}
	return out, nil
}
