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

type ListDnsZonesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkDnsZone

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListDnsZonesOperationResponse, error)
}

type ListDnsZonesCompleteResult struct {
	Items []WorkloadNetworkDnsZone
}

func (r ListDnsZonesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListDnsZonesOperationResponse) LoadMore(ctx context.Context) (resp ListDnsZonesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListDnsZones ...
func (c WorkloadNetworksClient) ListDnsZones(ctx context.Context, id PrivateCloudId) (resp ListDnsZonesOperationResponse, err error) {
	req, err := c.preparerForListDnsZones(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsZones", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsZones", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListDnsZones(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsZones", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListDnsZones prepares the ListDnsZones request.
func (c WorkloadNetworksClient) preparerForListDnsZones(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/dnsZones", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListDnsZonesWithNextLink prepares the ListDnsZones request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListDnsZonesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListDnsZones handles the response to the ListDnsZones request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListDnsZones(resp *http.Response) (result ListDnsZonesOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkDnsZone `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListDnsZonesOperationResponse, err error) {
			req, err := c.preparerForListDnsZonesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsZones", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsZones", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListDnsZones(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsZones", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListDnsZonesComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListDnsZonesComplete(ctx context.Context, id PrivateCloudId) (ListDnsZonesCompleteResult, error) {
	return c.ListDnsZonesCompleteMatchingPredicate(ctx, id, WorkloadNetworkDnsZoneOperationPredicate{})
}

// ListDnsZonesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListDnsZonesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDnsZoneOperationPredicate) (resp ListDnsZonesCompleteResult, err error) {
	items := make([]WorkloadNetworkDnsZone, 0)

	page, err := c.ListDnsZones(ctx, id)
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

	out := ListDnsZonesCompleteResult{
		Items: items,
	}
	return out, nil
}
