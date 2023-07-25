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

type ListBasicAuthOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteBasicAuthPropertiesARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBasicAuthOperationResponse, error)
}

type ListBasicAuthCompleteResult struct {
	Items []StaticSiteBasicAuthPropertiesARMResource
}

func (r ListBasicAuthOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBasicAuthOperationResponse) LoadMore(ctx context.Context) (resp ListBasicAuthOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBasicAuth ...
func (c StaticSitesClient) ListBasicAuth(ctx context.Context, id StaticSiteId) (resp ListBasicAuthOperationResponse, err error) {
	req, err := c.preparerForListBasicAuth(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListBasicAuth", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListBasicAuth", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBasicAuth(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListBasicAuth", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBasicAuth prepares the ListBasicAuth request.
func (c StaticSitesClient) preparerForListBasicAuth(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicAuth", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBasicAuthWithNextLink prepares the ListBasicAuth request with the given nextLink token.
func (c StaticSitesClient) preparerForListBasicAuthWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBasicAuth handles the response to the ListBasicAuth request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListBasicAuth(resp *http.Response) (result ListBasicAuthOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteBasicAuthPropertiesARMResource `json:"value"`
		NextLink *string                                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBasicAuthOperationResponse, err error) {
			req, err := c.preparerForListBasicAuthWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListBasicAuth", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListBasicAuth", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBasicAuth(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListBasicAuth", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBasicAuthComplete retrieves all of the results into a single object
func (c StaticSitesClient) ListBasicAuthComplete(ctx context.Context, id StaticSiteId) (ListBasicAuthCompleteResult, error) {
	return c.ListBasicAuthCompleteMatchingPredicate(ctx, id, StaticSiteBasicAuthPropertiesARMResourceOperationPredicate{})
}

// ListBasicAuthCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) ListBasicAuthCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteBasicAuthPropertiesARMResourceOperationPredicate) (resp ListBasicAuthCompleteResult, err error) {
	items := make([]StaticSiteBasicAuthPropertiesARMResource, 0)

	page, err := c.ListBasicAuth(ctx, id)
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

	out := ListBasicAuthCompleteResult{
		Items: items,
	}
	return out, nil
}
