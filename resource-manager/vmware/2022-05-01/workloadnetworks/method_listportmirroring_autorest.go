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

type ListPortMirroringOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkloadNetworkPortMirroring

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListPortMirroringOperationResponse, error)
}

type ListPortMirroringCompleteResult struct {
	Items []WorkloadNetworkPortMirroring
}

func (r ListPortMirroringOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListPortMirroringOperationResponse) LoadMore(ctx context.Context) (resp ListPortMirroringOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListPortMirroring ...
func (c WorkloadNetworksClient) ListPortMirroring(ctx context.Context, id PrivateCloudId) (resp ListPortMirroringOperationResponse, err error) {
	req, err := c.preparerForListPortMirroring(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPortMirroring", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPortMirroring", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListPortMirroring(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPortMirroring", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListPortMirroring prepares the ListPortMirroring request.
func (c WorkloadNetworksClient) preparerForListPortMirroring(ctx context.Context, id PrivateCloudId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/workloadNetworks/default/portMirroringProfiles", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListPortMirroringWithNextLink prepares the ListPortMirroring request with the given nextLink token.
func (c WorkloadNetworksClient) preparerForListPortMirroringWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListPortMirroring handles the response to the ListPortMirroring request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForListPortMirroring(resp *http.Response) (result ListPortMirroringOperationResponse, err error) {
	type page struct {
		Values   []WorkloadNetworkPortMirroring `json:"value"`
		NextLink *string                        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListPortMirroringOperationResponse, err error) {
			req, err := c.preparerForListPortMirroringWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPortMirroring", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPortMirroring", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListPortMirroring(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "ListPortMirroring", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListPortMirroringComplete retrieves all of the results into a single object
func (c WorkloadNetworksClient) ListPortMirroringComplete(ctx context.Context, id PrivateCloudId) (ListPortMirroringCompleteResult, error) {
	return c.ListPortMirroringCompleteMatchingPredicate(ctx, id, WorkloadNetworkPortMirroringOperationPredicate{})
}

// ListPortMirroringCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WorkloadNetworksClient) ListPortMirroringCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkPortMirroringOperationPredicate) (resp ListPortMirroringCompleteResult, err error) {
	items := make([]WorkloadNetworkPortMirroring, 0)

	page, err := c.ListPortMirroring(ctx, id)
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

	out := ListPortMirroringCompleteResult{
		Items: items,
	}
	return out, nil
}
