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

type MachinesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Machine

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (MachinesListBySubscriptionOperationResponse, error)
}

type MachinesListBySubscriptionCompleteResult struct {
	Items []Machine
}

func (r MachinesListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r MachinesListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp MachinesListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// MachinesListBySubscription ...
func (c MachinesClient) MachinesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp MachinesListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForMachinesListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForMachinesListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForMachinesListBySubscription prepares the MachinesListBySubscription request.
func (c MachinesClient) preparerForMachinesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
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

// preparerForMachinesListBySubscriptionWithNextLink prepares the MachinesListBySubscription request with the given nextLink token.
func (c MachinesClient) preparerForMachinesListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForMachinesListBySubscription handles the response to the MachinesListBySubscription request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForMachinesListBySubscription(resp *http.Response) (result MachinesListBySubscriptionOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result MachinesListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForMachinesListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForMachinesListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// MachinesListBySubscriptionComplete retrieves all of the results into a single object
func (c MachinesClient) MachinesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (MachinesListBySubscriptionCompleteResult, error) {
	return c.MachinesListBySubscriptionCompleteMatchingPredicate(ctx, id, MachineOperationPredicate{})
}

// MachinesListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MachinesClient) MachinesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate MachineOperationPredicate) (resp MachinesListBySubscriptionCompleteResult, err error) {
	items := make([]Machine, 0)

	page, err := c.MachinesListBySubscription(ctx, id)
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

	out := MachinesListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
