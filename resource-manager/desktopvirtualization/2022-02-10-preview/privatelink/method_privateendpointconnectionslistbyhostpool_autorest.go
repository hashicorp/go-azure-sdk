package privatelink

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

type PrivateEndpointConnectionsListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PrivateEndpointConnectionWithSystemData

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (PrivateEndpointConnectionsListByHostPoolOperationResponse, error)
}

type PrivateEndpointConnectionsListByHostPoolCompleteResult struct {
	Items []PrivateEndpointConnectionWithSystemData
}

func (r PrivateEndpointConnectionsListByHostPoolOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r PrivateEndpointConnectionsListByHostPoolOperationResponse) LoadMore(ctx context.Context) (resp PrivateEndpointConnectionsListByHostPoolOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// PrivateEndpointConnectionsListByHostPool ...
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPool(ctx context.Context, id HostPoolId) (resp PrivateEndpointConnectionsListByHostPoolOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsListByHostPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsListByHostPool", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsListByHostPool", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForPrivateEndpointConnectionsListByHostPool(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsListByHostPool", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// PrivateEndpointConnectionsListByHostPoolComplete retrieves all of the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPoolComplete(ctx context.Context, id HostPoolId) (PrivateEndpointConnectionsListByHostPoolCompleteResult, error) {
	return c.PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionWithSystemDataOperationPredicate{})
}

// PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, predicate PrivateEndpointConnectionWithSystemDataOperationPredicate) (resp PrivateEndpointConnectionsListByHostPoolCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionWithSystemData, 0)

	page, err := c.PrivateEndpointConnectionsListByHostPool(ctx, id)
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

	out := PrivateEndpointConnectionsListByHostPoolCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForPrivateEndpointConnectionsListByHostPool prepares the PrivateEndpointConnectionsListByHostPool request.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsListByHostPool(ctx context.Context, id HostPoolId) (*http.Request, error) {
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

// preparerForPrivateEndpointConnectionsListByHostPoolWithNextLink prepares the PrivateEndpointConnectionsListByHostPool request with the given nextLink token.
func (c PrivateLinkClient) preparerForPrivateEndpointConnectionsListByHostPoolWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionsListByHostPool handles the response to the PrivateEndpointConnectionsListByHostPool request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForPrivateEndpointConnectionsListByHostPool(resp *http.Response) (result PrivateEndpointConnectionsListByHostPoolOperationResponse, err error) {
	type page struct {
		Values   []PrivateEndpointConnectionWithSystemData `json:"value"`
		NextLink *string                                   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result PrivateEndpointConnectionsListByHostPoolOperationResponse, err error) {
			req, err := c.preparerForPrivateEndpointConnectionsListByHostPoolWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsListByHostPool", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsListByHostPool", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForPrivateEndpointConnectionsListByHostPool(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "PrivateEndpointConnectionsListByHostPool", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
