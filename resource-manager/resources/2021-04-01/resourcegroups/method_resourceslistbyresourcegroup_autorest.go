package resourcegroups

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

type ResourcesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GenericResourceExpanded

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ResourcesListByResourceGroupOperationResponse, error)
}

type ResourcesListByResourceGroupCompleteResult struct {
	Items []GenericResourceExpanded
}

func (r ResourcesListByResourceGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ResourcesListByResourceGroupOperationResponse) LoadMore(ctx context.Context) (resp ResourcesListByResourceGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ResourcesListByResourceGroupOperationOptions struct {
	Expand *string
	Filter *string
	Top    *int64
}

func DefaultResourcesListByResourceGroupOperationOptions() ResourcesListByResourceGroupOperationOptions {
	return ResourcesListByResourceGroupOperationOptions{}
}

func (o ResourcesListByResourceGroupOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ResourcesListByResourceGroupOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ResourcesListByResourceGroup ...
func (c ResourceGroupsClient) ResourcesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options ResourcesListByResourceGroupOperationOptions) (resp ResourcesListByResourceGroupOperationResponse, err error) {
	req, err := c.preparerForResourcesListByResourceGroup(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ResourcesListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ResourcesListByResourceGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForResourcesListByResourceGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ResourcesListByResourceGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForResourcesListByResourceGroup prepares the ResourcesListByResourceGroup request.
func (c ResourceGroupsClient) preparerForResourcesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options ResourcesListByResourceGroupOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/resources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForResourcesListByResourceGroupWithNextLink prepares the ResourcesListByResourceGroup request with the given nextLink token.
func (c ResourceGroupsClient) preparerForResourcesListByResourceGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForResourcesListByResourceGroup handles the response to the ResourcesListByResourceGroup request. The method always
// closes the http.Response Body.
func (c ResourceGroupsClient) responderForResourcesListByResourceGroup(resp *http.Response) (result ResourcesListByResourceGroupOperationResponse, err error) {
	type page struct {
		Values   []GenericResourceExpanded `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ResourcesListByResourceGroupOperationResponse, err error) {
			req, err := c.preparerForResourcesListByResourceGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ResourcesListByResourceGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ResourcesListByResourceGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForResourcesListByResourceGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "ResourcesListByResourceGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ResourcesListByResourceGroupComplete retrieves all of the results into a single object
func (c ResourceGroupsClient) ResourcesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options ResourcesListByResourceGroupOperationOptions) (ResourcesListByResourceGroupCompleteResult, error) {
	return c.ResourcesListByResourceGroupCompleteMatchingPredicate(ctx, id, options, GenericResourceExpandedOperationPredicate{})
}

// ResourcesListByResourceGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceGroupsClient) ResourcesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options ResourcesListByResourceGroupOperationOptions, predicate GenericResourceExpandedOperationPredicate) (resp ResourcesListByResourceGroupCompleteResult, err error) {
	items := make([]GenericResourceExpanded, 0)

	page, err := c.ResourcesListByResourceGroup(ctx, id, options)
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

	out := ResourcesListByResourceGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
