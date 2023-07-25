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

type GetPrivateEndpointConnectionListSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RemotePrivateEndpointConnectionARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetPrivateEndpointConnectionListSlotOperationResponse, error)
}

type GetPrivateEndpointConnectionListSlotCompleteResult struct {
	Items []RemotePrivateEndpointConnectionARMResource
}

func (r GetPrivateEndpointConnectionListSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetPrivateEndpointConnectionListSlotOperationResponse) LoadMore(ctx context.Context) (resp GetPrivateEndpointConnectionListSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetPrivateEndpointConnectionListSlot ...
func (c WebAppsClient) GetPrivateEndpointConnectionListSlot(ctx context.Context, id SlotId) (resp GetPrivateEndpointConnectionListSlotOperationResponse, err error) {
	req, err := c.preparerForGetPrivateEndpointConnectionListSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionListSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionListSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetPrivateEndpointConnectionListSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionListSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetPrivateEndpointConnectionListSlot prepares the GetPrivateEndpointConnectionListSlot request.
func (c WebAppsClient) preparerForGetPrivateEndpointConnectionListSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateEndpointConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetPrivateEndpointConnectionListSlotWithNextLink prepares the GetPrivateEndpointConnectionListSlot request with the given nextLink token.
func (c WebAppsClient) preparerForGetPrivateEndpointConnectionListSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetPrivateEndpointConnectionListSlot handles the response to the GetPrivateEndpointConnectionListSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPrivateEndpointConnectionListSlot(resp *http.Response) (result GetPrivateEndpointConnectionListSlotOperationResponse, err error) {
	type page struct {
		Values   []RemotePrivateEndpointConnectionARMResource `json:"value"`
		NextLink *string                                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetPrivateEndpointConnectionListSlotOperationResponse, err error) {
			req, err := c.preparerForGetPrivateEndpointConnectionListSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionListSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionListSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetPrivateEndpointConnectionListSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionListSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetPrivateEndpointConnectionListSlotComplete retrieves all of the results into a single object
func (c WebAppsClient) GetPrivateEndpointConnectionListSlotComplete(ctx context.Context, id SlotId) (GetPrivateEndpointConnectionListSlotCompleteResult, error) {
	return c.GetPrivateEndpointConnectionListSlotCompleteMatchingPredicate(ctx, id, RemotePrivateEndpointConnectionARMResourceOperationPredicate{})
}

// GetPrivateEndpointConnectionListSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) GetPrivateEndpointConnectionListSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate RemotePrivateEndpointConnectionARMResourceOperationPredicate) (resp GetPrivateEndpointConnectionListSlotCompleteResult, err error) {
	items := make([]RemotePrivateEndpointConnectionARMResource, 0)

	page, err := c.GetPrivateEndpointConnectionListSlot(ctx, id)
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

	out := GetPrivateEndpointConnectionListSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
