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

type GatewaysListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GatewayResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GatewaysListOperationResponse, error)
}

type GatewaysListCompleteResult struct {
	Items []GatewayResource
}

func (r GatewaysListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GatewaysListOperationResponse) LoadMore(ctx context.Context) (resp GatewaysListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GatewaysList ...
func (c AppPlatformClient) GatewaysList(ctx context.Context, id SpringId) (resp GatewaysListOperationResponse, err error) {
	req, err := c.preparerForGatewaysList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGatewaysList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGatewaysList prepares the GatewaysList request.
func (c AppPlatformClient) preparerForGatewaysList(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/gateways", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGatewaysListWithNextLink prepares the GatewaysList request with the given nextLink token.
func (c AppPlatformClient) preparerForGatewaysListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGatewaysList handles the response to the GatewaysList request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewaysList(resp *http.Response) (result GatewaysListOperationResponse, err error) {
	type page struct {
		Values   []GatewayResource `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GatewaysListOperationResponse, err error) {
			req, err := c.preparerForGatewaysListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGatewaysList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GatewaysListComplete retrieves all of the results into a single object
func (c AppPlatformClient) GatewaysListComplete(ctx context.Context, id SpringId) (GatewaysListCompleteResult, error) {
	return c.GatewaysListCompleteMatchingPredicate(ctx, id, GatewayResourceOperationPredicate{})
}

// GatewaysListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppPlatformClient) GatewaysListCompleteMatchingPredicate(ctx context.Context, id SpringId, predicate GatewayResourceOperationPredicate) (resp GatewaysListCompleteResult, err error) {
	items := make([]GatewayResource, 0)

	page, err := c.GatewaysList(ctx, id)
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

	out := GatewaysListCompleteResult{
		Items: items,
	}
	return out, nil
}
