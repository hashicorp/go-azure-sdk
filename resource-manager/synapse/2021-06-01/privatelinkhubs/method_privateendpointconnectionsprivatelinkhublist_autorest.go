package privatelinkhubs

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

type PrivateEndpointConnectionsPrivateLinkHubListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PrivateEndpointConnectionForPrivateLinkHub

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (PrivateEndpointConnectionsPrivateLinkHubListOperationResponse, error)
}

type PrivateEndpointConnectionsPrivateLinkHubListCompleteResult struct {
	Items []PrivateEndpointConnectionForPrivateLinkHub
}

func (r PrivateEndpointConnectionsPrivateLinkHubListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r PrivateEndpointConnectionsPrivateLinkHubListOperationResponse) LoadMore(ctx context.Context) (resp PrivateEndpointConnectionsPrivateLinkHubListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// PrivateEndpointConnectionsPrivateLinkHubList ...
func (c PrivateLinkHubsClient) PrivateEndpointConnectionsPrivateLinkHubList(ctx context.Context, id PrivateLinkHubId) (resp PrivateEndpointConnectionsPrivateLinkHubListOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsPrivateLinkHubList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkhubs.PrivateLinkHubsClient", "PrivateEndpointConnectionsPrivateLinkHubList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkhubs.PrivateLinkHubsClient", "PrivateEndpointConnectionsPrivateLinkHubList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForPrivateEndpointConnectionsPrivateLinkHubList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkhubs.PrivateLinkHubsClient", "PrivateEndpointConnectionsPrivateLinkHubList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForPrivateEndpointConnectionsPrivateLinkHubList prepares the PrivateEndpointConnectionsPrivateLinkHubList request.
func (c PrivateLinkHubsClient) preparerForPrivateEndpointConnectionsPrivateLinkHubList(ctx context.Context, id PrivateLinkHubId) (*http.Request, error) {
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

// preparerForPrivateEndpointConnectionsPrivateLinkHubListWithNextLink prepares the PrivateEndpointConnectionsPrivateLinkHubList request with the given nextLink token.
func (c PrivateLinkHubsClient) preparerForPrivateEndpointConnectionsPrivateLinkHubListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForPrivateEndpointConnectionsPrivateLinkHubList handles the response to the PrivateEndpointConnectionsPrivateLinkHubList request. The method always
// closes the http.Response Body.
func (c PrivateLinkHubsClient) responderForPrivateEndpointConnectionsPrivateLinkHubList(resp *http.Response) (result PrivateEndpointConnectionsPrivateLinkHubListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result PrivateEndpointConnectionsPrivateLinkHubListOperationResponse, err error) {
			req, err := c.preparerForPrivateEndpointConnectionsPrivateLinkHubListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelinkhubs.PrivateLinkHubsClient", "PrivateEndpointConnectionsPrivateLinkHubList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelinkhubs.PrivateLinkHubsClient", "PrivateEndpointConnectionsPrivateLinkHubList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForPrivateEndpointConnectionsPrivateLinkHubList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelinkhubs.PrivateLinkHubsClient", "PrivateEndpointConnectionsPrivateLinkHubList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// PrivateEndpointConnectionsPrivateLinkHubListComplete retrieves all of the results into a single object
func (c PrivateLinkHubsClient) PrivateEndpointConnectionsPrivateLinkHubListComplete(ctx context.Context, id PrivateLinkHubId) (PrivateEndpointConnectionsPrivateLinkHubListCompleteResult, error) {
	return c.PrivateEndpointConnectionsPrivateLinkHubListCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionForPrivateLinkHubOperationPredicate{})
}

// PrivateEndpointConnectionsPrivateLinkHubListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PrivateLinkHubsClient) PrivateEndpointConnectionsPrivateLinkHubListCompleteMatchingPredicate(ctx context.Context, id PrivateLinkHubId, predicate PrivateEndpointConnectionForPrivateLinkHubOperationPredicate) (resp PrivateEndpointConnectionsPrivateLinkHubListCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionForPrivateLinkHub, 0)

	page, err := c.PrivateEndpointConnectionsPrivateLinkHubList(ctx, id)
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

	out := PrivateEndpointConnectionsPrivateLinkHubListCompleteResult{
		Items: items,
	}
	return out, nil
}
