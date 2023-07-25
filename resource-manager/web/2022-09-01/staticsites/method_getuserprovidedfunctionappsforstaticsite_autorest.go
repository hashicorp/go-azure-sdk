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

type GetUserProvidedFunctionAppsForStaticSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteUserProvidedFunctionAppARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetUserProvidedFunctionAppsForStaticSiteOperationResponse, error)
}

type GetUserProvidedFunctionAppsForStaticSiteCompleteResult struct {
	Items []StaticSiteUserProvidedFunctionAppARMResource
}

func (r GetUserProvidedFunctionAppsForStaticSiteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetUserProvidedFunctionAppsForStaticSiteOperationResponse) LoadMore(ctx context.Context) (resp GetUserProvidedFunctionAppsForStaticSiteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetUserProvidedFunctionAppsForStaticSite ...
func (c StaticSitesClient) GetUserProvidedFunctionAppsForStaticSite(ctx context.Context, id StaticSiteId) (resp GetUserProvidedFunctionAppsForStaticSiteOperationResponse, err error) {
	req, err := c.preparerForGetUserProvidedFunctionAppsForStaticSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSite", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSite", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetUserProvidedFunctionAppsForStaticSite(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSite", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetUserProvidedFunctionAppsForStaticSite prepares the GetUserProvidedFunctionAppsForStaticSite request.
func (c StaticSitesClient) preparerForGetUserProvidedFunctionAppsForStaticSite(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/userProvidedFunctionApps", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetUserProvidedFunctionAppsForStaticSiteWithNextLink prepares the GetUserProvidedFunctionAppsForStaticSite request with the given nextLink token.
func (c StaticSitesClient) preparerForGetUserProvidedFunctionAppsForStaticSiteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetUserProvidedFunctionAppsForStaticSite handles the response to the GetUserProvidedFunctionAppsForStaticSite request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetUserProvidedFunctionAppsForStaticSite(resp *http.Response) (result GetUserProvidedFunctionAppsForStaticSiteOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteUserProvidedFunctionAppARMResource `json:"value"`
		NextLink *string                                        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetUserProvidedFunctionAppsForStaticSiteOperationResponse, err error) {
			req, err := c.preparerForGetUserProvidedFunctionAppsForStaticSiteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSite", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSite", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetUserProvidedFunctionAppsForStaticSite(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSite", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetUserProvidedFunctionAppsForStaticSiteComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetUserProvidedFunctionAppsForStaticSiteComplete(ctx context.Context, id StaticSiteId) (GetUserProvidedFunctionAppsForStaticSiteCompleteResult, error) {
	return c.GetUserProvidedFunctionAppsForStaticSiteCompleteMatchingPredicate(ctx, id, StaticSiteUserProvidedFunctionAppARMResourceOperationPredicate{})
}

// GetUserProvidedFunctionAppsForStaticSiteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetUserProvidedFunctionAppsForStaticSiteCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteUserProvidedFunctionAppARMResourceOperationPredicate) (resp GetUserProvidedFunctionAppsForStaticSiteCompleteResult, err error) {
	items := make([]StaticSiteUserProvidedFunctionAppARMResource, 0)

	page, err := c.GetUserProvidedFunctionAppsForStaticSite(ctx, id)
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

	out := GetUserProvidedFunctionAppsForStaticSiteCompleteResult{
		Items: items,
	}
	return out, nil
}
