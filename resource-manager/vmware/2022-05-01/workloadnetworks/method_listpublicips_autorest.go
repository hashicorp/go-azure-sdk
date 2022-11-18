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

type ListPublicIPsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkPublicIP

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListPublicIPsOperationResponse, error)
}

type ListPublicIPsCompleteResult struct {
	Items []WorkloadNetworkPublicIP
}

func (r ListPublicIPsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListPublicIPsOperationResponse) LoadMore(ctx context.Context) (resp ListPublicIPsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListPublicIPs ...
func (c WorkloadNetworksClient) ListPublicIPs(ctx context.Context, id PrivateCloudId) (resp ListPublicIPsOperationResponse, err error) {
	req, err := c.preparerForListPublicIPs(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPublicIPs", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPublicIPs", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListPublicIPs(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPublicIPs", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListPublicIPs prepares the ListPublicIPs request.
func (c WorkloadNetworksClient) preparerForListPublicIPs(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/publicIPs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListPublicIPsWithNextLink prepares the ListPublicIPs request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListPublicIPsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListPublicIPs handles the response to the ListPublicIPs request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListPublicIPs(resp *http.Response) (result ListPublicIPsOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkPublicIP `json:"value"`
		NextLink *string                   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListPublicIPsOperationResponse, err error) {
			req, err := c.preparerForListPublicIPsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPublicIPs", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPublicIPs", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListPublicIPs(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPublicIPs", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListPublicIPsComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListPublicIPsComplete(ctx context.Context, id PrivateCloudId) (ListPublicIPsCompleteResult, error) {
	return c.ListPublicIPsCompleteMatchingPredicate(ctx, id, WorkloadNetworkPublicIPOperationPredicate{})
}

// ListPublicIPsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListPublicIPsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkPublicIPOperationPredicate) (resp ListPublicIPsCompleteResult, err error) {
	items := make([]WorkloadNetworkPublicIP, 0)

	page, err := c.ListPublicIPs(ctx, id)
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

	out := ListPublicIPsCompleteResult{
		Items: items,
	}
	return out, nil
}
