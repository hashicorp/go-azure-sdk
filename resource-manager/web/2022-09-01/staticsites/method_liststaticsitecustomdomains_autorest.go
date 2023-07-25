package staticsites

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

type ListStaticSiteCustomDomainsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteCustomDomainOverviewARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListStaticSiteCustomDomainsOperationResponse, error)
}

type ListStaticSiteCustomDomainsCompleteResult struct {
	Items []StaticSiteCustomDomainOverviewARMResource
}

func (r ListStaticSiteCustomDomainsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListStaticSiteCustomDomainsOperationResponse) LoadMore(ctx context.Context) (resp ListStaticSiteCustomDomainsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListStaticSiteCustomDomains ...
func (c StaticSitesClient) ListStaticSiteCustomDomains(ctx context.Context, id StaticSiteId) (resp ListStaticSiteCustomDomainsOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteCustomDomains(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteCustomDomains", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteCustomDomains", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListStaticSiteCustomDomains(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteCustomDomains", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListStaticSiteCustomDomains prepares the ListStaticSiteCustomDomains request.
func (c StaticSitesClient) preparerForListStaticSiteCustomDomains(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/customDomains", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListStaticSiteCustomDomainsWithNextLink prepares the ListStaticSiteCustomDomains request with the given nextLink token.
func (c StaticSitesClient) preparerForListStaticSiteCustomDomainsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListStaticSiteCustomDomains handles the response to the ListStaticSiteCustomDomains request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteCustomDomains(resp *http.Response) (result ListStaticSiteCustomDomainsOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteCustomDomainOverviewARMResource `json:"value"`
		NextLink *string                                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListStaticSiteCustomDomainsOperationResponse, err error) {
			req, err := c.preparerForListStaticSiteCustomDomainsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteCustomDomains", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteCustomDomains", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListStaticSiteCustomDomains(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteCustomDomains", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListStaticSiteCustomDomainsComplete retrieves all of the results into a single object
func (c StaticSitesClient) ListStaticSiteCustomDomainsComplete(ctx context.Context, id StaticSiteId) (ListStaticSiteCustomDomainsCompleteResult, error) {
	return c.ListStaticSiteCustomDomainsCompleteMatchingPredicate(ctx, id, StaticSiteCustomDomainOverviewARMResourceOperationPredicate{})
}

// ListStaticSiteCustomDomainsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) ListStaticSiteCustomDomainsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteCustomDomainOverviewARMResourceOperationPredicate) (resp ListStaticSiteCustomDomainsCompleteResult, err error) {
	items := make([]StaticSiteCustomDomainOverviewARMResource, 0)

	page, err := c.ListStaticSiteCustomDomains(ctx, id)
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

	out := ListStaticSiteCustomDomainsCompleteResult{
		Items: items,
	}
	return out, nil
}
