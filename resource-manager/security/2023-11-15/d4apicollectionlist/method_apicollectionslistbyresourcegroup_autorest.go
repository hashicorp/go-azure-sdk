package d4apicollectionlist

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APICollectionsListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiCollection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (APICollectionsListByResourceGroupOperationResponse, error)
}

type APICollectionsListByResourceGroupCompleteResult struct {
	Items []ApiCollection
}

func (r APICollectionsListByResourceGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r APICollectionsListByResourceGroupOperationResponse) LoadMore(ctx context.Context) (resp APICollectionsListByResourceGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// APICollectionsListByResourceGroup ...
func (c D4APICollectionListClient) APICollectionsListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (resp APICollectionsListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForAPICollectionsListByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "d4apicollectionlist.D4APICollectionListClient", "APICollectionsListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "d4apicollectionlist.D4APICollectionListClient", "APICollectionsListByResourceGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAPICollectionsListByResourceGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "d4apicollectionlist.D4APICollectionListClient", "APICollectionsListByResourceGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAPICollectionsListByResourceGroup prepares the APICollectionsListByResourceGroup request.
func (c D4APICollectionListClient) preparerForAPICollectionsListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/apiCollections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAPICollectionsListByResourceGroupWithNextLink prepares the APICollectionsListByResourceGroup request with the given nextLink token.
func (c D4APICollectionListClient) preparerForAPICollectionsListByResourceGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAPICollectionsListByResourceGroup handles the response to the APICollectionsListByResourceGroup request. The method always
// closes the http.Response Body.
func (c D4APICollectionListClient) responderForAPICollectionsListByResourceGroup(resp *http.Response) (result APICollectionsListByResourceGroupOperationResponse, err error) {
	type page struct {
		Values   []ApiCollection `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result APICollectionsListByResourceGroupOperationResponse, err error) {
			req, err := c.preparerForAPICollectionsListByResourceGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "d4apicollectionlist.D4APICollectionListClient", "APICollectionsListByResourceGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "d4apicollectionlist.D4APICollectionListClient", "APICollectionsListByResourceGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAPICollectionsListByResourceGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "d4apicollectionlist.D4APICollectionListClient", "APICollectionsListByResourceGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// APICollectionsListByResourceGroupComplete retrieves all of the results into a single object
func (c D4APICollectionListClient) APICollectionsListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (APICollectionsListByResourceGroupCompleteResult, error) {
	return c.APICollectionsListByResourceGroupCompleteMatchingPredicate(ctx, id, ApiCollectionOperationPredicate{})
}

// APICollectionsListByResourceGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c D4APICollectionListClient) APICollectionsListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate ApiCollectionOperationPredicate) (resp APICollectionsListByResourceGroupCompleteResult, err error) {
	items := make([]ApiCollection, 0)

	page, err := c.APICollectionsListByResourceGroup(ctx, id)
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

	out := APICollectionsListByResourceGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
