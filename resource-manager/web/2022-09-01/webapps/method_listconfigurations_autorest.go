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

type ListConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SiteConfigResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListConfigurationsOperationResponse, error)
}

type ListConfigurationsCompleteResult struct {
	Items []SiteConfigResource
}

func (r ListConfigurationsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListConfigurationsOperationResponse) LoadMore(ctx context.Context) (resp ListConfigurationsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListConfigurations ...
func (c WebAppsClient) ListConfigurations(ctx context.Context, id SiteId) (resp ListConfigurationsOperationResponse, err error) {
	req, err := c.preparerForListConfigurations(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurations", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurations", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListConfigurations(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurations", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListConfigurations prepares the ListConfigurations request.
func (c WebAppsClient) preparerForListConfigurations(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListConfigurationsWithNextLink prepares the ListConfigurations request with the given nextLink token.
func (c WebAppsClient) preparerForListConfigurationsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListConfigurations handles the response to the ListConfigurations request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListConfigurations(resp *http.Response) (result ListConfigurationsOperationResponse, err error) {
	type page struct {
		Values   []SiteConfigResource `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListConfigurationsOperationResponse, err error) {
			req, err := c.preparerForListConfigurationsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurations", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurations", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListConfigurations(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurations", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListConfigurationsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListConfigurationsComplete(ctx context.Context, id SiteId) (ListConfigurationsCompleteResult, error) {
	return c.ListConfigurationsCompleteMatchingPredicate(ctx, id, SiteConfigResourceOperationPredicate{})
}

// ListConfigurationsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListConfigurationsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate SiteConfigResourceOperationPredicate) (resp ListConfigurationsCompleteResult, err error) {
	items := make([]SiteConfigResource, 0)

	page, err := c.ListConfigurations(ctx, id)
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

	out := ListConfigurationsCompleteResult{
		Items: items,
	}
	return out, nil
}
