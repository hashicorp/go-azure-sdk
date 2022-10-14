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

type ListVirtualMachinesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkVirtualMachine

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListVirtualMachinesOperationResponse, error)
}

type ListVirtualMachinesCompleteResult struct {
	Items []WorkloadNetworkVirtualMachine
}

func (r ListVirtualMachinesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListVirtualMachinesOperationResponse) LoadMore(ctx context.Context) (resp ListVirtualMachinesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListVirtualMachines ...
func (c WorkloadNetworksClient) ListVirtualMachines(ctx context.Context, id PrivateCloudId) (resp ListVirtualMachinesOperationResponse, err error) {
	req, err := c.preparerForListVirtualMachines(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListVirtualMachines", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListVirtualMachines", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListVirtualMachines(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListVirtualMachines", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListVirtualMachines prepares the ListVirtualMachines request.
func (c WorkloadNetworksClient) preparerForListVirtualMachines(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/virtualMachines", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListVirtualMachinesWithNextLink prepares the ListVirtualMachines request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListVirtualMachinesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListVirtualMachines handles the response to the ListVirtualMachines request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListVirtualMachines(resp *http.Response) (result ListVirtualMachinesOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkVirtualMachine `json:"value"`
		NextLink *string                         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListVirtualMachinesOperationResponse, err error) {
			req, err := c.preparerForListVirtualMachinesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListVirtualMachines", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListVirtualMachines", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListVirtualMachines(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListVirtualMachines", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListVirtualMachinesComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListVirtualMachinesComplete(ctx context.Context, id PrivateCloudId) (ListVirtualMachinesCompleteResult, error) {
	return c.ListVirtualMachinesCompleteMatchingPredicate(ctx, id, WorkloadNetworkVirtualMachineOperationPredicate{})
}

// ListVirtualMachinesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListVirtualMachinesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkVirtualMachineOperationPredicate) (resp ListVirtualMachinesCompleteResult, err error) {
	items := make([]WorkloadNetworkVirtualMachine, 0)

	page, err := c.ListVirtualMachines(ctx, id)
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

	out := ListVirtualMachinesCompleteResult{
		Items: items,
	}
	return out, nil
}
