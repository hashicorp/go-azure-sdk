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

type ListBasicPublishingCredentialsPoliciesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CsmPublishingCredentialsPoliciesEntity

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBasicPublishingCredentialsPoliciesSlotOperationResponse, error)
}

type ListBasicPublishingCredentialsPoliciesSlotCompleteResult struct {
	Items []CsmPublishingCredentialsPoliciesEntity
}

func (r ListBasicPublishingCredentialsPoliciesSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBasicPublishingCredentialsPoliciesSlotOperationResponse) LoadMore(ctx context.Context) (resp ListBasicPublishingCredentialsPoliciesSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBasicPublishingCredentialsPoliciesSlot ...
func (c WebAppsClient) ListBasicPublishingCredentialsPoliciesSlot(ctx context.Context, id SlotId) (resp ListBasicPublishingCredentialsPoliciesSlotOperationResponse, err error) {
	req, err := c.preparerForListBasicPublishingCredentialsPoliciesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPoliciesSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPoliciesSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBasicPublishingCredentialsPoliciesSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPoliciesSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBasicPublishingCredentialsPoliciesSlot prepares the ListBasicPublishingCredentialsPoliciesSlot request.
func (c WebAppsClient) preparerForListBasicPublishingCredentialsPoliciesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// preparerForListBasicPublishingCredentialsPoliciesSlotWithNextLink prepares the ListBasicPublishingCredentialsPoliciesSlot request with the given nextLink token.
func (c WebAppsClient) preparerForListBasicPublishingCredentialsPoliciesSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBasicPublishingCredentialsPoliciesSlot handles the response to the ListBasicPublishingCredentialsPoliciesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListBasicPublishingCredentialsPoliciesSlot(resp *http.Response) (result ListBasicPublishingCredentialsPoliciesSlotOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBasicPublishingCredentialsPoliciesSlotOperationResponse, err error) {
			req, err := c.preparerForListBasicPublishingCredentialsPoliciesSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPoliciesSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPoliciesSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBasicPublishingCredentialsPoliciesSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListBasicPublishingCredentialsPoliciesSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBasicPublishingCredentialsPoliciesSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) ListBasicPublishingCredentialsPoliciesSlotComplete(ctx context.Context, id SlotId) (ListBasicPublishingCredentialsPoliciesSlotCompleteResult, error) {
	return c.ListBasicPublishingCredentialsPoliciesSlotCompleteMatchingPredicate(ctx, id, CsmPublishingCredentialsPoliciesEntityOperationPredicate{})
}

// ListBasicPublishingCredentialsPoliciesSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListBasicPublishingCredentialsPoliciesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate CsmPublishingCredentialsPoliciesEntityOperationPredicate) (resp ListBasicPublishingCredentialsPoliciesSlotCompleteResult, err error) {
	items := make([]CsmPublishingCredentialsPoliciesEntity, 0)

	page, err := c.ListBasicPublishingCredentialsPoliciesSlot(ctx, id)
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

	out := ListBasicPublishingCredentialsPoliciesSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
