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

type IotDpsResourceListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProvisioningServiceDescription

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IotDpsResourceListByResourceGroupOperationResponse, error)
}

type IotDpsResourceListByResourceGroupCompleteResult struct {
	Items []ProvisioningServiceDescription
}

func (r IotDpsResourceListByResourceGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IotDpsResourceListByResourceGroupOperationResponse) LoadMore(ctx context.Context) (resp IotDpsResourceListByResourceGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// IotDpsResourceListByResourceGroup ...
func (c GETClient) IotDpsResourceListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (resp IotDpsResourceListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceListByResourceGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListByResourceGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIotDpsResourceListByResourceGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListByResourceGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// IotDpsResourceListByResourceGroupComplete retrieves all of the results into a single object
func (c GETClient) IotDpsResourceListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (IotDpsResourceListByResourceGroupCompleteResult, error) {
	return c.IotDpsResourceListByResourceGroupCompleteMatchingPredicate(ctx, id, ProvisioningServiceDescriptionOperationPredicate{})
}

// IotDpsResourceListByResourceGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c GETClient) IotDpsResourceListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate ProvisioningServiceDescriptionOperationPredicate) (resp IotDpsResourceListByResourceGroupCompleteResult, err error) {
	items := make([]ProvisioningServiceDescription, 0)

	page, err := c.IotDpsResourceListByResourceGroup(ctx, id)
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

	out := IotDpsResourceListByResourceGroupCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForIotDpsResourceListByResourceGroup prepares the IotDpsResourceListByResourceGroup request.
func (c GETClient) preparerForIotDpsResourceListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
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

// preparerForIotDpsResourceListByResourceGroupWithNextLink prepares the IotDpsResourceListByResourceGroup request with the given nextLink token.
func (c GETClient) preparerForIotDpsResourceListByResourceGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForIotDpsResourceListByResourceGroup handles the response to the IotDpsResourceListByResourceGroup request. The method always
// closes the http.Response Body.
func (c GETClient) responderForIotDpsResourceListByResourceGroup(resp *http.Response) (result IotDpsResourceListByResourceGroupOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IotDpsResourceListByResourceGroupOperationResponse, err error) {
			req, err := c.preparerForIotDpsResourceListByResourceGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListByResourceGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListByResourceGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIotDpsResourceListByResourceGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceListByResourceGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
