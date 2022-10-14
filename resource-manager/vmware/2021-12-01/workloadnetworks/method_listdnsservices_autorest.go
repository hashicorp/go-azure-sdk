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

type ListDnsServicesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkDnsService

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListDnsServicesOperationResponse, error)
}

type ListDnsServicesCompleteResult struct {
	Items []WorkloadNetworkDnsService
}

func (r ListDnsServicesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListDnsServicesOperationResponse) LoadMore(ctx context.Context) (resp ListDnsServicesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListDnsServices ...
func (c WorkloadNetworksClient) ListDnsServices(ctx context.Context, id PrivateCloudId) (resp ListDnsServicesOperationResponse, err error) {
	req, err := c.preparerForListDnsServices(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsServices", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsServices", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListDnsServices(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsServices", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListDnsServices prepares the ListDnsServices request.
func (c WorkloadNetworksClient) preparerForListDnsServices(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/dnsServices", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListDnsServicesWithNextLink prepares the ListDnsServices request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListDnsServicesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListDnsServices handles the response to the ListDnsServices request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListDnsServices(resp *http.Response) (result ListDnsServicesOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkDnsService `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListDnsServicesOperationResponse, err error) {
			req, err := c.preparerForListDnsServicesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsServices", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsServices", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListDnsServices(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListDnsServices", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListDnsServicesComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListDnsServicesComplete(ctx context.Context, id PrivateCloudId) (ListDnsServicesCompleteResult, error) {
	return c.ListDnsServicesCompleteMatchingPredicate(ctx, id, WorkloadNetworkDnsServiceOperationPredicate{})
}

// ListDnsServicesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListDnsServicesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkDnsServiceOperationPredicate) (resp ListDnsServicesCompleteResult, err error) {
	items := make([]WorkloadNetworkDnsService, 0)

	page, err := c.ListDnsServices(ctx, id)
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

	out := ListDnsServicesCompleteResult{
		Items: items,
	}
	return out, nil
}
