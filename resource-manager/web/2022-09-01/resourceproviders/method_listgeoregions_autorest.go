package resourceproviders

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

type ListGeoRegionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GeoRegion

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListGeoRegionsOperationResponse, error)
}

type ListGeoRegionsCompleteResult struct {
	Items []GeoRegion
}

func (r ListGeoRegionsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListGeoRegionsOperationResponse) LoadMore(ctx context.Context) (resp ListGeoRegionsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListGeoRegionsOperationOptions struct {
	LinuxDynamicWorkersEnabled *bool
	LinuxWorkersEnabled        *bool
	Sku                        *SkuName
	XenonWorkersEnabled        *bool
}

func DefaultListGeoRegionsOperationOptions() ListGeoRegionsOperationOptions {
	return ListGeoRegionsOperationOptions{}
}

func (o ListGeoRegionsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListGeoRegionsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.LinuxDynamicWorkersEnabled != nil {
		out["linuxDynamicWorkersEnabled"] = *o.LinuxDynamicWorkersEnabled
	}

	if o.LinuxWorkersEnabled != nil {
		out["linuxWorkersEnabled"] = *o.LinuxWorkersEnabled
	}

	if o.Sku != nil {
		out["sku"] = *o.Sku
	}

	if o.XenonWorkersEnabled != nil {
		out["xenonWorkersEnabled"] = *o.XenonWorkersEnabled
	}

	return out
}

// ListGeoRegions ...
func (c ResourceProvidersClient) ListGeoRegions(ctx context.Context, id commonids.SubscriptionId, options ListGeoRegionsOperationOptions) (resp ListGeoRegionsOperationResponse, err error) {
	req, err := c.preparerForListGeoRegions(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListGeoRegions", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListGeoRegions", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListGeoRegions(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListGeoRegions", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListGeoRegions prepares the ListGeoRegions request.
func (c ResourceProvidersClient) preparerForListGeoRegions(ctx context.Context, id commonids.SubscriptionId, options ListGeoRegionsOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/geoRegions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListGeoRegionsWithNextLink prepares the ListGeoRegions request with the given nextLink token.
func (c ResourceProvidersClient) preparerForListGeoRegionsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListGeoRegions handles the response to the ListGeoRegions request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForListGeoRegions(resp *http.Response) (result ListGeoRegionsOperationResponse, err error) {
	type page struct {
		Values   []GeoRegion `json:"value"`
		NextLink *string     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListGeoRegionsOperationResponse, err error) {
			req, err := c.preparerForListGeoRegionsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListGeoRegions", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListGeoRegions", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListGeoRegions(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ListGeoRegions", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListGeoRegionsComplete retrieves all of the results into a single object
func (c ResourceProvidersClient) ListGeoRegionsComplete(ctx context.Context, id commonids.SubscriptionId, options ListGeoRegionsOperationOptions) (ListGeoRegionsCompleteResult, error) {
	return c.ListGeoRegionsCompleteMatchingPredicate(ctx, id, options, GeoRegionOperationPredicate{})
}

// ListGeoRegionsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceProvidersClient) ListGeoRegionsCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options ListGeoRegionsOperationOptions, predicate GeoRegionOperationPredicate) (resp ListGeoRegionsCompleteResult, err error) {
	items := make([]GeoRegion, 0)

	page, err := c.ListGeoRegions(ctx, id, options)
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

	out := ListGeoRegionsCompleteResult{
		Items: items,
	}
	return out, nil
}
