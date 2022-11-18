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

type ListSegmentsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkSegment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSegmentsOperationResponse, error)
}

type ListSegmentsCompleteResult struct {
	Items []WorkloadNetworkSegment
}

func (r ListSegmentsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSegmentsOperationResponse) LoadMore(ctx context.Context) (resp ListSegmentsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSegments ...
func (c WorkloadNetworksClient) ListSegments(ctx context.Context, id PrivateCloudId) (resp ListSegmentsOperationResponse, err error) {
	req, err := c.preparerForListSegments(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListSegments", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListSegments", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSegments(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListSegments", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSegments prepares the ListSegments request.
func (c WorkloadNetworksClient) preparerForListSegments(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/segments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSegmentsWithNextLink prepares the ListSegments request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListSegmentsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListSegments handles the response to the ListSegments request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListSegments(resp *http.Response) (result ListSegmentsOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkSegment `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSegmentsOperationResponse, err error) {
			req, err := c.preparerForListSegmentsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListSegments", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListSegments", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSegments(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListSegments", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSegmentsComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListSegmentsComplete(ctx context.Context, id PrivateCloudId) (ListSegmentsCompleteResult, error) {
	return c.ListSegmentsCompleteMatchingPredicate(ctx, id, WorkloadNetworkSegmentOperationPredicate{})
}

// ListSegmentsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListSegmentsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkSegmentOperationPredicate) (resp ListSegmentsCompleteResult, err error) {
	items := make([]WorkloadNetworkSegment, 0)

	page, err := c.ListSegments(ctx, id)
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

	out := ListSegmentsCompleteResult{
		Items: items,
	}
	return out, nil
}
