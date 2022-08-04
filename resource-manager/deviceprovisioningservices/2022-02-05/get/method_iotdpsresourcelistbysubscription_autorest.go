package get

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

type IotDpsResourceListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProvisioningServiceDescription

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IotDpsResourceListBySubscriptionOperationResponse, error)
}

type IotDpsResourceListBySubscriptionCompleteResult struct {
	Items []ProvisioningServiceDescription
}

func (r IotDpsResourceListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IotDpsResourceListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp IotDpsResourceListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// IotDpsResourceListBySubscription ...
func (c GETClient) IotDpsResourceListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp IotDpsResourceListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIotDpsResourceListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// IotDpsResourceListBySubscriptionComplete retrieves all of the results into a single object
func (c GETClient) IotDpsResourceListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (IotDpsResourceListBySubscriptionCompleteResult, error) {
	return c.IotDpsResourceListBySubscriptionCompleteMatchingPredicate(ctx, id, ProvisioningServiceDescriptionOperationPredicate{})
}

// IotDpsResourceListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c GETClient) IotDpsResourceListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ProvisioningServiceDescriptionOperationPredicate) (resp IotDpsResourceListBySubscriptionCompleteResult, err error) {
	items := make([]ProvisioningServiceDescription, 0)

	page, err := c.IotDpsResourceListBySubscription(ctx, id)
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

	out := IotDpsResourceListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForIotDpsResourceListBySubscription prepares the IotDpsResourceListBySubscription request.
func (c GETClient) preparerForIotDpsResourceListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Devices/provisioningServices", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForIotDpsResourceListBySubscriptionWithNextLink prepares the IotDpsResourceListBySubscription request with the given nextLink token.
func (c GETClient) preparerForIotDpsResourceListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIotDpsResourceListBySubscription handles the response to the IotDpsResourceListBySubscription request. The method always
// closes the http.Response Body.
func (c GETClient) responderForIotDpsResourceListBySubscription(resp *http.Response) (result IotDpsResourceListBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []ProvisioningServiceDescription `json:"value"`
		NextLink *string                          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IotDpsResourceListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForIotDpsResourceListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIotDpsResourceListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
