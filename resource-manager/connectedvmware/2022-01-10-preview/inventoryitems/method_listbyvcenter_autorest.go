package inventoryitems

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

type ListByVCenterOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]InventoryItem

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByVCenterOperationResponse, error)
}

type ListByVCenterCompleteResult struct {
	Items []InventoryItem
}

func (r ListByVCenterOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByVCenterOperationResponse) LoadMore(ctx context.Context) (resp ListByVCenterOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByVCenter ...
func (c InventoryItemsClient) ListByVCenter(ctx context.Context, id VCenterId) (resp ListByVCenterOperationResponse, err error) {
	req, err := c.preparerForListByVCenter(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "inventoryitems.InventoryItemsClient", "ListByVCenter", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "inventoryitems.InventoryItemsClient", "ListByVCenter", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByVCenter(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "inventoryitems.InventoryItemsClient", "ListByVCenter", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByVCenter prepares the ListByVCenter request.
func (c InventoryItemsClient) preparerForListByVCenter(ctx context.Context, id VCenterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/inventoryItems", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByVCenterWithNextLink prepares the ListByVCenter request with the given nextLink token.
func (c InventoryItemsClient) preparerForListByVCenterWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByVCenter handles the response to the ListByVCenter request. The method always
// closes the http.Response Body.
func (c InventoryItemsClient) responderForListByVCenter(resp *http.Response) (result ListByVCenterOperationResponse, err error) {
	type page struct {
		Values   []InventoryItem `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByVCenterOperationResponse, err error) {
			req, err := c.preparerForListByVCenterWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "inventoryitems.InventoryItemsClient", "ListByVCenter", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "inventoryitems.InventoryItemsClient", "ListByVCenter", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByVCenter(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "inventoryitems.InventoryItemsClient", "ListByVCenter", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByVCenterComplete retrieves all of the results into a single object
func (c InventoryItemsClient) ListByVCenterComplete(ctx context.Context, id VCenterId) (ListByVCenterCompleteResult, error) {
	return c.ListByVCenterCompleteMatchingPredicate(ctx, id, InventoryItemOperationPredicate{})
}

// ListByVCenterCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c InventoryItemsClient) ListByVCenterCompleteMatchingPredicate(ctx context.Context, id VCenterId, predicate InventoryItemOperationPredicate) (resp ListByVCenterCompleteResult, err error) {
	items := make([]InventoryItem, 0)

	page, err := c.ListByVCenter(ctx, id)
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

	out := ListByVCenterCompleteResult{
		Items: items,
	}
	return out, nil
}
