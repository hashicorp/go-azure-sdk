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

type GetPrivateEndpointConnectionListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RemotePrivateEndpointConnectionARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetPrivateEndpointConnectionListOperationResponse, error)
}

type GetPrivateEndpointConnectionListCompleteResult struct {
	Items []RemotePrivateEndpointConnectionARMResource
}

func (r GetPrivateEndpointConnectionListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetPrivateEndpointConnectionListOperationResponse) LoadMore(ctx context.Context) (resp GetPrivateEndpointConnectionListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetPrivateEndpointConnectionList ...
func (c WebAppsClient) GetPrivateEndpointConnectionList(ctx context.Context, id SiteId) (resp GetPrivateEndpointConnectionListOperationResponse, err error) {
	req, err := c.preparerForGetPrivateEndpointConnectionList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetPrivateEndpointConnectionList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetPrivateEndpointConnectionList prepares the GetPrivateEndpointConnectionList request.
func (c WebAppsClient) preparerForGetPrivateEndpointConnectionList(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForGetPrivateEndpointConnectionListWithNextLink prepares the GetPrivateEndpointConnectionList request with the given nextLink token.
func (c WebAppsClient) preparerForGetPrivateEndpointConnectionListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetPrivateEndpointConnectionList handles the response to the GetPrivateEndpointConnectionList request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPrivateEndpointConnectionList(resp *http.Response) (result GetPrivateEndpointConnectionListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetPrivateEndpointConnectionListOperationResponse, err error) {
			req, err := c.preparerForGetPrivateEndpointConnectionListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetPrivateEndpointConnectionList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateEndpointConnectionList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetPrivateEndpointConnectionListComplete retrieves all of the results into a single object
func (c WebAppsClient) GetPrivateEndpointConnectionListComplete(ctx context.Context, id SiteId) (GetPrivateEndpointConnectionListCompleteResult, error) {
	return c.GetPrivateEndpointConnectionListCompleteMatchingPredicate(ctx, id, RemotePrivateEndpointConnectionARMResourceOperationPredicate{})
}

// GetPrivateEndpointConnectionListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) GetPrivateEndpointConnectionListCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate RemotePrivateEndpointConnectionARMResourceOperationPredicate) (resp GetPrivateEndpointConnectionListCompleteResult, err error) {
	items := make([]RemotePrivateEndpointConnectionARMResource, 0)

	page, err := c.GetPrivateEndpointConnectionList(ctx, id)
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

	out := GetPrivateEndpointConnectionListCompleteResult{
		Items: items,
	}
	return out, nil
}
