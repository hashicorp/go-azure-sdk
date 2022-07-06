package outboundnetworkdependenciesendpoints

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

type ClustersListOutboundNetworkDependenciesEndpointsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]OutboundNetworkDependenciesEndpoint

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ClustersListOutboundNetworkDependenciesEndpointsOperationResponse, error)
}

type ClustersListOutboundNetworkDependenciesEndpointsCompleteResult struct {
	Items []OutboundNetworkDependenciesEndpoint
}

func (r ClustersListOutboundNetworkDependenciesEndpointsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ClustersListOutboundNetworkDependenciesEndpointsOperationResponse) LoadMore(ctx context.Context) (resp ClustersListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ClustersListOutboundNetworkDependenciesEndpoints ...
func (c OutboundNetworkDependenciesEndpointsClient) ClustersListOutboundNetworkDependenciesEndpoints(ctx context.Context, id ClusterId) (resp ClustersListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	req, err := c.preparerForClustersListOutboundNetworkDependenciesEndpoints(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "ClustersListOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "ClustersListOutboundNetworkDependenciesEndpoints", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForClustersListOutboundNetworkDependenciesEndpoints(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "ClustersListOutboundNetworkDependenciesEndpoints", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ClustersListOutboundNetworkDependenciesEndpointsComplete retrieves all of the results into a single object
func (c OutboundNetworkDependenciesEndpointsClient) ClustersListOutboundNetworkDependenciesEndpointsComplete(ctx context.Context, id ClusterId) (ClustersListOutboundNetworkDependenciesEndpointsCompleteResult, error) {
	return c.ClustersListOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate(ctx, id, OutboundNetworkDependenciesEndpointOperationPredicate{})
}

// ClustersListOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c OutboundNetworkDependenciesEndpointsClient) ClustersListOutboundNetworkDependenciesEndpointsCompleteMatchingPredicate(ctx context.Context, id ClusterId, predicate OutboundNetworkDependenciesEndpointOperationPredicate) (resp ClustersListOutboundNetworkDependenciesEndpointsCompleteResult, err error) {
	items := make([]OutboundNetworkDependenciesEndpoint, 0)

	page, err := c.ClustersListOutboundNetworkDependenciesEndpoints(ctx, id)
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

	out := ClustersListOutboundNetworkDependenciesEndpointsCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForClustersListOutboundNetworkDependenciesEndpoints prepares the ClustersListOutboundNetworkDependenciesEndpoints request.
func (c OutboundNetworkDependenciesEndpointsClient) preparerForClustersListOutboundNetworkDependenciesEndpoints(ctx context.Context, id ClusterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/outboundNetworkDependenciesEndpoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForClustersListOutboundNetworkDependenciesEndpointsWithNextLink prepares the ClustersListOutboundNetworkDependenciesEndpoints request with the given nextLink token.
func (c OutboundNetworkDependenciesEndpointsClient) preparerForClustersListOutboundNetworkDependenciesEndpointsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForClustersListOutboundNetworkDependenciesEndpoints handles the response to the ClustersListOutboundNetworkDependenciesEndpoints request. The method always
// closes the http.Response Body.
func (c OutboundNetworkDependenciesEndpointsClient) responderForClustersListOutboundNetworkDependenciesEndpoints(resp *http.Response) (result ClustersListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
	type page struct {
		Values   []OutboundNetworkDependenciesEndpoint `json:"value"`
		NextLink *string                               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ClustersListOutboundNetworkDependenciesEndpointsOperationResponse, err error) {
			req, err := c.preparerForClustersListOutboundNetworkDependenciesEndpointsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "ClustersListOutboundNetworkDependenciesEndpoints", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "ClustersListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForClustersListOutboundNetworkDependenciesEndpoints(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient", "ClustersListOutboundNetworkDependenciesEndpoints", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
