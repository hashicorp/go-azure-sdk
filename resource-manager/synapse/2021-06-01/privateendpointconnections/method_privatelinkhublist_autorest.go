package privateendpointconnections

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

type PrivateLinkHubListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PrivateEndpointConnectionForPrivateLinkHub

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (PrivateLinkHubListOperationResponse, error)
}

type PrivateLinkHubListCompleteResult struct {
	Items []PrivateEndpointConnectionForPrivateLinkHub
}

func (r PrivateLinkHubListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r PrivateLinkHubListOperationResponse) LoadMore(ctx context.Context) (resp PrivateLinkHubListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// PrivateLinkHubList ...
func (c PrivateEndpointConnectionsClient) PrivateLinkHubList(ctx context.Context, id PrivateLinkHubId) (resp PrivateLinkHubListOperationResponse, err error) {
	req, err := c.preparerForPrivateLinkHubList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForPrivateLinkHubList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForPrivateLinkHubList prepares the PrivateLinkHubList request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateLinkHubList(ctx context.Context, id PrivateLinkHubId) (*http.Request, error) {
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

// preparerForPrivateLinkHubListWithNextLink prepares the PrivateLinkHubList request with the given nextLink token.
func (c PrivateEndpointConnectionsClient) preparerForPrivateLinkHubListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForPrivateLinkHubList handles the response to the PrivateLinkHubList request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateLinkHubList(resp *http.Response) (result PrivateLinkHubListOperationResponse, err error) {
	type page struct {
		Values   []PrivateEndpointConnectionForPrivateLinkHub `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result PrivateLinkHubListOperationResponse, err error) {
			req, err := c.preparerForPrivateLinkHubListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForPrivateLinkHubList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateLinkHubList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// PrivateLinkHubListComplete retrieves all of the results into a single object
func (c PrivateEndpointConnectionsClient) PrivateLinkHubListComplete(ctx context.Context, id PrivateLinkHubId) (PrivateLinkHubListCompleteResult, error) {
	return c.PrivateLinkHubListCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionForPrivateLinkHubOperationPredicate{})
}

// PrivateLinkHubListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PrivateEndpointConnectionsClient) PrivateLinkHubListCompleteMatchingPredicate(ctx context.Context, id PrivateLinkHubId, predicate PrivateEndpointConnectionForPrivateLinkHubOperationPredicate) (resp PrivateLinkHubListCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionForPrivateLinkHub, 0)

	page, err := c.PrivateLinkHubList(ctx, id)
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

	out := PrivateLinkHubListCompleteResult{
		Items: items,
	}
	return out, nil
}
