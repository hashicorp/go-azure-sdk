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

type ListConfigurationsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SiteConfigResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListConfigurationsSlotOperationResponse, error)
}

type ListConfigurationsSlotCompleteResult struct {
	Items []SiteConfigResource
}

func (r ListConfigurationsSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListConfigurationsSlotOperationResponse) LoadMore(ctx context.Context) (resp ListConfigurationsSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListConfigurationsSlot ...
func (c WebAppsClient) ListConfigurationsSlot(ctx context.Context, id SlotId) (resp ListConfigurationsSlotOperationResponse, err error) {
	req, err := c.preparerForListConfigurationsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationsSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationsSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListConfigurationsSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationsSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListConfigurationsSlot prepares the ListConfigurationsSlot request.
func (c WebAppsClient) preparerForListConfigurationsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// preparerForListConfigurationsSlotWithNextLink prepares the ListConfigurationsSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListConfigurationsSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListConfigurationsSlot handles the response to the ListConfigurationsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListConfigurationsSlot(resp *http.Response) (result ListConfigurationsSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListConfigurationsSlotOperationResponse, err error) {
			req, err := c.preparerForListConfigurationsSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationsSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationsSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListConfigurationsSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListConfigurationsSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListConfigurationsSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListConfigurationsSlotComplete(ctx context.Context, id SlotId) (ListConfigurationsSlotCompleteResult, error) {
	return c.ListConfigurationsSlotCompleteMatchingPredicate(ctx, id, SiteConfigResourceOperationPredicate{})
}

// ListConfigurationsSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListConfigurationsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate SiteConfigResourceOperationPredicate) (resp ListConfigurationsSlotCompleteResult, err error) {
	items := make([]SiteConfigResource, 0)

	page, err := c.ListConfigurationsSlot(ctx, id)
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

	out := ListConfigurationsSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
