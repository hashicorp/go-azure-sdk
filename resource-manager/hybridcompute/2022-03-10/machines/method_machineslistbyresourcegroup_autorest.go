package machines

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

type MachinesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Machine

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (MachinesListByResourceGroupOperationResponse, error)
}

type MachinesListByResourceGroupCompleteResult struct {
	Items []Machine
}

func (r MachinesListByResourceGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r MachinesListByResourceGroupOperationResponse) LoadMore(ctx context.Context) (resp MachinesListByResourceGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// MachinesListByResourceGroup ...
func (c MachinesClient) MachinesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (resp MachinesListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForMachinesListByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListByResourceGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForMachinesListByResourceGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListByResourceGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForMachinesListByResourceGroup prepares the MachinesListByResourceGroup request.
func (c MachinesClient) preparerForMachinesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.HybridCompute/machines", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForMachinesListByResourceGroupWithNextLink prepares the MachinesListByResourceGroup request with the given nextLink token.
func (c MachinesClient) preparerForMachinesListByResourceGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForMachinesListByResourceGroup handles the response to the MachinesListByResourceGroup request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForMachinesListByResourceGroup(resp *http.Response) (result MachinesListByResourceGroupOperationResponse, err error) {
	type page struct {
		Values   []Machine `json:"value"`
		NextLink *string   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result MachinesListByResourceGroupOperationResponse, err error) {
			req, err := c.preparerForMachinesListByResourceGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListByResourceGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListByResourceGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForMachinesListByResourceGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListByResourceGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// MachinesListByResourceGroupComplete retrieves all of the results into a single object
func (c MachinesClient) MachinesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (MachinesListByResourceGroupCompleteResult, error) {
	return c.MachinesListByResourceGroupCompleteMatchingPredicate(ctx, id, MachineOperationPredicate{})
}

// MachinesListByResourceGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MachinesClient) MachinesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate MachineOperationPredicate) (resp MachinesListByResourceGroupCompleteResult, err error) {
	items := make([]Machine, 0)

	page, err := c.MachinesListByResourceGroup(ctx, id)
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

	out := MachinesListByResourceGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
