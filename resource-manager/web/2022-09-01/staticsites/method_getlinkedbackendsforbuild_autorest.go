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

type GetLinkedBackendsForBuildOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteLinkedBackendARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetLinkedBackendsForBuildOperationResponse, error)
}

type GetLinkedBackendsForBuildCompleteResult struct {
	Items []StaticSiteLinkedBackendARMResource
}

func (r GetLinkedBackendsForBuildOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetLinkedBackendsForBuildOperationResponse) LoadMore(ctx context.Context) (resp GetLinkedBackendsForBuildOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetLinkedBackendsForBuild ...
func (c StaticSitesClient) GetLinkedBackendsForBuild(ctx context.Context, id BuildId) (resp GetLinkedBackendsForBuildOperationResponse, err error) {
	req, err := c.preparerForGetLinkedBackendsForBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendsForBuild", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendsForBuild", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetLinkedBackendsForBuild(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendsForBuild", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetLinkedBackendsForBuild prepares the GetLinkedBackendsForBuild request.
func (c StaticSitesClient) preparerForGetLinkedBackendsForBuild(ctx context.Context, id BuildId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/linkedBackends", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetLinkedBackendsForBuildWithNextLink prepares the GetLinkedBackendsForBuild request with the given nextLink token.
func (c StaticSitesClient) preparerForGetLinkedBackendsForBuildWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetLinkedBackendsForBuild handles the response to the GetLinkedBackendsForBuild request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetLinkedBackendsForBuild(resp *http.Response) (result GetLinkedBackendsForBuildOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteLinkedBackendARMResource `json:"value"`
		NextLink *string                              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetLinkedBackendsForBuildOperationResponse, err error) {
			req, err := c.preparerForGetLinkedBackendsForBuildWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendsForBuild", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendsForBuild", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetLinkedBackendsForBuild(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetLinkedBackendsForBuild", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetLinkedBackendsForBuildComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetLinkedBackendsForBuildComplete(ctx context.Context, id BuildId) (GetLinkedBackendsForBuildCompleteResult, error) {
	return c.GetLinkedBackendsForBuildCompleteMatchingPredicate(ctx, id, StaticSiteLinkedBackendARMResourceOperationPredicate{})
}

// GetLinkedBackendsForBuildCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetLinkedBackendsForBuildCompleteMatchingPredicate(ctx context.Context, id BuildId, predicate StaticSiteLinkedBackendARMResourceOperationPredicate) (resp GetLinkedBackendsForBuildCompleteResult, err error) {
	items := make([]StaticSiteLinkedBackendARMResource, 0)

	page, err := c.GetLinkedBackendsForBuild(ctx, id)
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

	out := GetLinkedBackendsForBuildCompleteResult{
		Items: items,
	}
	return out, nil
}
