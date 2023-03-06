package privateendpointconnections

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

type ListByResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PrivateEndpointConnection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByResourceOperationResponse, error)
}

type ListByResourceCompleteResult struct {
	Items []PrivateEndpointConnection
}

func (r ListByResourceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByResourceOperationResponse) LoadMore(ctx context.Context) (resp ListByResourceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByResource ...
func (c PrivateEndpointConnectionsClient) ListByResource(ctx context.Context, id commonids.KeyVaultId) (resp ListByResourceOperationResponse, err error) {
	req, err := c.preparerForListByResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByResource", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByResource", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByResource(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByResource", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByResource prepares the ListByResource request.
func (c PrivateEndpointConnectionsClient) preparerForListByResource(ctx context.Context, id commonids.KeyVaultId) (*http.Request, error) {
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

// preparerForListByResourceWithNextLink prepares the ListByResource request with the given nextLink token.
func (c PrivateEndpointConnectionsClient) preparerForListByResourceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByResource handles the response to the ListByResource request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForListByResource(resp *http.Response) (result ListByResourceOperationResponse, err error) {
	type page struct {
		Values   []PrivateEndpointConnection `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByResourceOperationResponse, err error) {
			req, err := c.preparerForListByResourceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByResource", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByResource", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByResource(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByResource", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByResourceComplete retrieves all of the results into a single object
func (c PrivateEndpointConnectionsClient) ListByResourceComplete(ctx context.Context, id commonids.KeyVaultId) (ListByResourceCompleteResult, error) {
	return c.ListByResourceCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// ListByResourceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PrivateEndpointConnectionsClient) ListByResourceCompleteMatchingPredicate(ctx context.Context, id commonids.KeyVaultId, predicate PrivateEndpointConnectionOperationPredicate) (resp ListByResourceCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	page, err := c.ListByResource(ctx, id)
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

	out := ListByResourceCompleteResult{
		Items: items,
	}
	return out, nil
}
