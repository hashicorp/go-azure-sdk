package appplatform

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

type GatewayRouteConfigsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GatewayRouteConfigResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GatewayRouteConfigsListOperationResponse, error)
}

type GatewayRouteConfigsListCompleteResult struct {
	Items []GatewayRouteConfigResource
}

func (r GatewayRouteConfigsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GatewayRouteConfigsListOperationResponse) LoadMore(ctx context.Context) (resp GatewayRouteConfigsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GatewayRouteConfigsList ...
func (c AppPlatformClient) GatewayRouteConfigsList(ctx context.Context, id GatewayId) (resp GatewayRouteConfigsListOperationResponse, err error) {
	req, err := c.preparerForGatewayRouteConfigsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGatewayRouteConfigsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGatewayRouteConfigsList prepares the GatewayRouteConfigsList request.
func (c AppPlatformClient) preparerForGatewayRouteConfigsList(ctx context.Context, id GatewayId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/routeConfigs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGatewayRouteConfigsListWithNextLink prepares the GatewayRouteConfigsList request with the given nextLink token.
func (c AppPlatformClient) preparerForGatewayRouteConfigsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGatewayRouteConfigsList handles the response to the GatewayRouteConfigsList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewayRouteConfigsList(resp *http.Response) (result GatewayRouteConfigsListOperationResponse, err error) {
	type page struct {
		Values   []GatewayRouteConfigResource `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GatewayRouteConfigsListOperationResponse, err error) {
			req, err := c.preparerForGatewayRouteConfigsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGatewayRouteConfigsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayRouteConfigsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GatewayRouteConfigsListComplete retrieves all of the results into a single object
func (c AppPlatformClient) GatewayRouteConfigsListComplete(ctx context.Context, id GatewayId) (GatewayRouteConfigsListCompleteResult, error) {
	return c.GatewayRouteConfigsListCompleteMatchingPredicate(ctx, id, GatewayRouteConfigResourceOperationPredicate{})
}

// GatewayRouteConfigsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) GatewayRouteConfigsListCompleteMatchingPredicate(ctx context.Context, id GatewayId, predicate GatewayRouteConfigResourceOperationPredicate) (resp GatewayRouteConfigsListCompleteResult, err error) {
	items := make([]GatewayRouteConfigResource, 0)

	page, err := c.GatewayRouteConfigsList(ctx, id)
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

	out := GatewayRouteConfigsListCompleteResult{
		Items: items,
	}
	return out, nil
}
