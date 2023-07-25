package webapps

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

type ListBasicPublishingCredentialsPoliciesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CsmPublishingCredentialsPoliciesEntity

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBasicPublishingCredentialsPoliciesOperationResponse, error)
}

type ListBasicPublishingCredentialsPoliciesCompleteResult struct {
	Items []CsmPublishingCredentialsPoliciesEntity
}

func (r ListBasicPublishingCredentialsPoliciesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBasicPublishingCredentialsPoliciesOperationResponse) LoadMore(ctx context.Context) (resp ListBasicPublishingCredentialsPoliciesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBasicPublishingCredentialsPolicies ...
func (c WebAppsClient) ListBasicPublishingCredentialsPolicies(ctx context.Context, id SiteId) (resp ListBasicPublishingCredentialsPoliciesOperationResponse, err error) {
	req, err := c.preparerForListBasicPublishingCredentialsPolicies(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPolicies", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPolicies", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBasicPublishingCredentialsPolicies(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPolicies", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBasicPublishingCredentialsPolicies prepares the ListBasicPublishingCredentialsPolicies request.
func (c WebAppsClient) preparerForListBasicPublishingCredentialsPolicies(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBasicPublishingCredentialsPoliciesWithNextLink prepares the ListBasicPublishingCredentialsPolicies request with the given nextLink token.
func (c WebAppsClient) preparerForListBasicPublishingCredentialsPoliciesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBasicPublishingCredentialsPolicies handles the response to the ListBasicPublishingCredentialsPolicies request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListBasicPublishingCredentialsPolicies(resp *http.Response) (result ListBasicPublishingCredentialsPoliciesOperationResponse, err error) {
	type page struct {
		Values   []CsmPublishingCredentialsPoliciesEntity `json:"value"`
		NextLink *string                                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBasicPublishingCredentialsPoliciesOperationResponse, err error) {
			req, err := c.preparerForListBasicPublishingCredentialsPoliciesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPolicies", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPolicies", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBasicPublishingCredentialsPolicies(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPolicies", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBasicPublishingCredentialsPoliciesComplete retrieves all of the results into a single object
func (c WebAppsClient) ListBasicPublishingCredentialsPoliciesComplete(ctx context.Context, id SiteId) (ListBasicPublishingCredentialsPoliciesCompleteResult, error) {
	return c.ListBasicPublishingCredentialsPoliciesCompleteMatchingPredicate(ctx, id, CsmPublishingCredentialsPoliciesEntityOperationPredicate{})
}

// ListBasicPublishingCredentialsPoliciesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListBasicPublishingCredentialsPoliciesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate CsmPublishingCredentialsPoliciesEntityOperationPredicate) (resp ListBasicPublishingCredentialsPoliciesCompleteResult, err error) {
	items := make([]CsmPublishingCredentialsPoliciesEntity, 0)

	page, err := c.ListBasicPublishingCredentialsPolicies(ctx, id)
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

	out := ListBasicPublishingCredentialsPoliciesCompleteResult{
		Items: items,
	}
	return out, nil
}
