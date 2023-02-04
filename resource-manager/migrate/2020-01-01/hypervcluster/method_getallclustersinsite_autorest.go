package hypervcluster

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

type GetAllClustersInSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]HyperVCluster

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAllClustersInSiteOperationResponse, error)
}

type GetAllClustersInSiteCompleteResult struct {
	Items []HyperVCluster
}

func (r GetAllClustersInSiteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAllClustersInSiteOperationResponse) LoadMore(ctx context.Context) (resp GetAllClustersInSiteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetAllClustersInSiteOperationOptions struct {
	Filter *string
}

func DefaultGetAllClustersInSiteOperationOptions() GetAllClustersInSiteOperationOptions {
	return GetAllClustersInSiteOperationOptions{}
}

func (o GetAllClustersInSiteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetAllClustersInSiteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// GetAllClustersInSite ...
func (c HyperVClusterClient) GetAllClustersInSite(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions) (resp GetAllClustersInSiteOperationResponse, err error) {
	req, err := c.preparerForGetAllClustersInSite(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetAllClustersInSite", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetAllClustersInSite", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAllClustersInSite(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetAllClustersInSite", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAllClustersInSite prepares the GetAllClustersInSite request.
func (c HyperVClusterClient) preparerForGetAllClustersInSite(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/clusters", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAllClustersInSiteWithNextLink prepares the GetAllClustersInSite request with the given nextLink token.
func (c HyperVClusterClient) preparerForGetAllClustersInSiteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAllClustersInSite handles the response to the GetAllClustersInSite request. The method always
// closes the http.Response Body.
func (c HyperVClusterClient) responderForGetAllClustersInSite(resp *http.Response) (result GetAllClustersInSiteOperationResponse, err error) {
	type page struct {
		Values   []HyperVCluster `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAllClustersInSiteOperationResponse, err error) {
			req, err := c.preparerForGetAllClustersInSiteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetAllClustersInSite", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetAllClustersInSite", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAllClustersInSite(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "hypervcluster.HyperVClusterClient", "GetAllClustersInSite", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAllClustersInSiteComplete retrieves all of the results into a single object
func (c HyperVClusterClient) GetAllClustersInSiteComplete(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions) (GetAllClustersInSiteCompleteResult, error) {
	return c.GetAllClustersInSiteCompleteMatchingPredicate(ctx, id, options, HyperVClusterOperationPredicate{})
}

// GetAllClustersInSiteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c HyperVClusterClient) GetAllClustersInSiteCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, options GetAllClustersInSiteOperationOptions, predicate HyperVClusterOperationPredicate) (resp GetAllClustersInSiteCompleteResult, err error) {
	items := make([]HyperVCluster, 0)

	page, err := c.GetAllClustersInSite(ctx, id, options)
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

	out := GetAllClustersInSiteCompleteResult{
		Items: items,
	}
	return out, nil
}
