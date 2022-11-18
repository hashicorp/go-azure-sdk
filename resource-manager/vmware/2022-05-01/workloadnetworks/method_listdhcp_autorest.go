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

type ListDhcpOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkDhcp

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListDhcpOperationResponse, error)
}

type ListDhcpCompleteResult struct {
	Items []WorkloadNetworkDhcp
}

func (r ListDhcpOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListDhcpOperationResponse) LoadMore(ctx context.Context) (resp ListDhcpOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListDhcp ...
func (c WorkloadNetworksClient) ListDhcp(ctx context.Context, id PrivateCloudId) (resp ListDhcpOperationResponse, err error) {
	req, err := c.preparerForListDhcp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDhcp", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDhcp", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListDhcp(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDhcp", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListDhcp prepares the ListDhcp request.
func (c WorkloadNetworksClient) preparerForListDhcp(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/dhcpConfigurations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListDhcpWithNextLink prepares the ListDhcp request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListDhcpWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListDhcp handles the response to the ListDhcp request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListDhcp(resp *http.Response) (result ListDhcpOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkDhcp `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListDhcpOperationResponse, err error) {
			req, err := c.preparerForListDhcpWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDhcp", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDhcp", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListDhcp(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDhcp", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListDhcpComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListDhcpComplete(ctx context.Context, id PrivateCloudId) (ListDhcpCompleteResult, error) {
	return c.ListDhcpCompleteMatchingPredicate(ctx, id, WorkloadNetworkDhcpOperationPredicate{})
}

// ListDhcpCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListDhcpCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDhcpOperationPredicate) (resp ListDhcpCompleteResult, err error) {
	items := make([]WorkloadNetworkDhcp, 0)

	page, err := c.ListDhcp(ctx, id)
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

	out := ListDhcpCompleteResult{
		Items: items,
	}
	return out, nil
}
