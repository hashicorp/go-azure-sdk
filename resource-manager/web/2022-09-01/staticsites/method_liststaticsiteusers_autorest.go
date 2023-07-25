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

type ListStaticSiteUsersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteUserARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListStaticSiteUsersOperationResponse, error)
}

type ListStaticSiteUsersCompleteResult struct {
	Items []StaticSiteUserARMResource
}

func (r ListStaticSiteUsersOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListStaticSiteUsersOperationResponse) LoadMore(ctx context.Context) (resp ListStaticSiteUsersOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListStaticSiteUsers ...
func (c StaticSitesClient) ListStaticSiteUsers(ctx context.Context, id AuthProviderId) (resp ListStaticSiteUsersOperationResponse, err error) {
	req, err := c.preparerForListStaticSiteUsers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteUsers", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteUsers", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListStaticSiteUsers(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteUsers", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListStaticSiteUsers prepares the ListStaticSiteUsers request.
func (c StaticSitesClient) preparerForListStaticSiteUsers(ctx context.Context, id AuthProviderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listUsers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListStaticSiteUsersWithNextLink prepares the ListStaticSiteUsers request with the given nextLink token.
func (c StaticSitesClient) preparerForListStaticSiteUsersWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListStaticSiteUsers handles the response to the ListStaticSiteUsers request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForListStaticSiteUsers(resp *http.Response) (result ListStaticSiteUsersOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteUserARMResource `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListStaticSiteUsersOperationResponse, err error) {
			req, err := c.preparerForListStaticSiteUsersWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteUsers", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteUsers", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListStaticSiteUsers(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ListStaticSiteUsers", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListStaticSiteUsersComplete retrieves all of the results into a single object
func (c StaticSitesClient) ListStaticSiteUsersComplete(ctx context.Context, id AuthProviderId) (ListStaticSiteUsersCompleteResult, error) {
	return c.ListStaticSiteUsersCompleteMatchingPredicate(ctx, id, StaticSiteUserARMResourceOperationPredicate{})
}

// ListStaticSiteUsersCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) ListStaticSiteUsersCompleteMatchingPredicate(ctx context.Context, id AuthProviderId, predicate StaticSiteUserARMResourceOperationPredicate) (resp ListStaticSiteUsersCompleteResult, err error) {
	items := make([]StaticSiteUserARMResource, 0)

	page, err := c.ListStaticSiteUsers(ctx, id)
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

	out := ListStaticSiteUsersCompleteResult{
		Items: items,
	}
	return out, nil
}
