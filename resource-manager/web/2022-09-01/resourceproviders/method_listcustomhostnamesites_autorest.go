package resourceproviders

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCustomHostNameSitesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CustomHostnameSites

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListCustomHostNameSitesOperationResponse, error)
}

type ListCustomHostNameSitesCompleteResult struct {
	Items []CustomHostnameSites
}

func (r ListCustomHostNameSitesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListCustomHostNameSitesOperationResponse) LoadMore(ctx context.Context) (resp ListCustomHostNameSitesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListCustomHostNameSitesOperationOptions struct {
	Hostname *string
}

func DefaultListCustomHostNameSitesOperationOptions() ListCustomHostNameSitesOperationOptions {
	return ListCustomHostNameSitesOperationOptions{}
}

func (o ListCustomHostNameSitesOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListCustomHostNameSitesOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Hostname != nil {
		out["hostname"] = *o.Hostname
	}

	return out
}

// ListCustomHostNameSites ...
func (c ResourceProvidersClient) ListCustomHostNameSites(ctx context.Context, id commonids.SubscriptionId, options ListCustomHostNameSitesOperationOptions) (resp ListCustomHostNameSitesOperationResponse, err error) {
	req, err := c.preparerForListCustomHostNameSites(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListCustomHostNameSites", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListCustomHostNameSites", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListCustomHostNameSites(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListCustomHostNameSites", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListCustomHostNameSites prepares the ListCustomHostNameSites request.
func (c ResourceProvidersClient) preparerForListCustomHostNameSites(ctx context.Context, id commonids.SubscriptionId, options ListCustomHostNameSitesOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/customhostnameSites", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListCustomHostNameSitesWithNextLink prepares the ListCustomHostNameSites request with the given nextLink token.
func (c ResourceProvidersClient) preparerForListCustomHostNameSitesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListCustomHostNameSites handles the response to the ListCustomHostNameSites request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForListCustomHostNameSites(resp *http.Response) (result ListCustomHostNameSitesOperationResponse, err error) {
	type page struct {
		Values   []CustomHostnameSites `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListCustomHostNameSitesOperationResponse, err error) {
			req, err := c.preparerForListCustomHostNameSitesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListCustomHostNameSites", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListCustomHostNameSites", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListCustomHostNameSites(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListCustomHostNameSites", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListCustomHostNameSitesComplete retrieves all of the results into a single object
func (c ResourceProvidersClient) ListCustomHostNameSitesComplete(ctx context.Context, id commonids.SubscriptionId, options ListCustomHostNameSitesOperationOptions) (ListCustomHostNameSitesCompleteResult, error) {
	return c.ListCustomHostNameSitesCompleteMatchingPredicate(ctx, id, options, CustomHostnameSitesOperationPredicate{})
}

// ListCustomHostNameSitesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceProvidersClient) ListCustomHostNameSitesCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ListCustomHostNameSitesOperationOptions, predicate CustomHostnameSitesOperationPredicate) (resp ListCustomHostNameSitesCompleteResult, err error) {
	items := make([]CustomHostnameSites, 0)

	page, err := c.ListCustomHostNameSites(ctx, id, options)
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

	out := ListCustomHostNameSitesCompleteResult{
		Items: items,
	}
	return out, nil
}
