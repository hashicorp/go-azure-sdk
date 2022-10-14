package workloadnetworks

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

type ListGatewaysOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkGateway

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListGatewaysOperationResponse, error)
}

type ListGatewaysCompleteResult struct {
	Items []WorkloadNetworkGateway
}

func (r ListGatewaysOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListGatewaysOperationResponse) LoadMore(ctx context.Context) (resp ListGatewaysOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListGateways ...
func (c WorkloadNetworksClient) ListGateways(ctx context.Context, id PrivateCloudId) (resp ListGatewaysOperationResponse, err error) {
	req, err := c.preparerForListGateways(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListGateways", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListGateways", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListGateways(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListGateways", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListGateways prepares the ListGateways request.
func (c WorkloadNetworksClient) preparerForListGateways(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/gateways", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListGatewaysWithNextLink prepares the ListGateways request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListGatewaysWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListGateways handles the response to the ListGateways request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListGateways(resp *http.Response) (result ListGatewaysOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkGateway `json:"value"`
		NextLink *string                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListGatewaysOperationResponse, err error) {
			req, err := c.preparerForListGatewaysWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListGateways", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListGateways", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListGateways(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListGateways", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListGatewaysComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListGatewaysComplete(ctx context.Context, id PrivateCloudId) (ListGatewaysCompleteResult, error) {
	return c.ListGatewaysCompleteMatchingPredicate(ctx, id, WorkloadNetworkGatewayOperationPredicate{})
}

// ListGatewaysCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListGatewaysCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkGatewayOperationPredicate) (resp ListGatewaysCompleteResult, err error) {
	items := make([]WorkloadNetworkGateway, 0)

	page, err := c.ListGateways(ctx, id)
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

	out := ListGatewaysCompleteResult{
		Items: items,
	}
	return out, nil
}
