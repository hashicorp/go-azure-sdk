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

type GetStaticSiteBuildsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteBuildARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetStaticSiteBuildsOperationResponse, error)
}

type GetStaticSiteBuildsCompleteResult struct {
	Items []StaticSiteBuildARMResource
}

func (r GetStaticSiteBuildsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetStaticSiteBuildsOperationResponse) LoadMore(ctx context.Context) (resp GetStaticSiteBuildsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetStaticSiteBuilds ...
func (c StaticSitesClient) GetStaticSiteBuilds(ctx context.Context, id StaticSiteId) (resp GetStaticSiteBuildsOperationResponse, err error) {
	req, err := c.preparerForGetStaticSiteBuilds(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuilds", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuilds", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetStaticSiteBuilds(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuilds", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetStaticSiteBuilds prepares the GetStaticSiteBuilds request.
func (c StaticSitesClient) preparerForGetStaticSiteBuilds(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/builds", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetStaticSiteBuildsWithNextLink prepares the GetStaticSiteBuilds request with the given nextLink token.
func (c StaticSitesClient) preparerForGetStaticSiteBuildsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetStaticSiteBuilds handles the response to the GetStaticSiteBuilds request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetStaticSiteBuilds(resp *http.Response) (result GetStaticSiteBuildsOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteBuildARMResource `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetStaticSiteBuildsOperationResponse, err error) {
			req, err := c.preparerForGetStaticSiteBuildsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuilds", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuilds", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetStaticSiteBuilds(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetStaticSiteBuilds", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetStaticSiteBuildsComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetStaticSiteBuildsComplete(ctx context.Context, id StaticSiteId) (GetStaticSiteBuildsCompleteResult, error) {
	return c.GetStaticSiteBuildsCompleteMatchingPredicate(ctx, id, StaticSiteBuildARMResourceOperationPredicate{})
}

// GetStaticSiteBuildsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetStaticSiteBuildsCompleteMatchingPredicate(ctx context.Context, id StaticSiteId, predicate StaticSiteBuildARMResourceOperationPredicate) (resp GetStaticSiteBuildsCompleteResult, err error) {
	items := make([]StaticSiteBuildARMResource, 0)

	page, err := c.GetStaticSiteBuilds(ctx, id)
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

	out := GetStaticSiteBuildsCompleteResult{
		Items: items,
	}
	return out, nil
}
