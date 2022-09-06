package volumes

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

type ListByVolumeGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Volume

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByVolumeGroupOperationResponse, error)
}

type ListByVolumeGroupCompleteResult struct {
	Items []Volume
}

func (r ListByVolumeGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByVolumeGroupOperationResponse) LoadMore(ctx context.Context) (resp ListByVolumeGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByVolumeGroup ...
func (c VolumesClient) ListByVolumeGroup(ctx context.Context, id VolumeGroupId) (resp ListByVolumeGroupOperationResponse, err error) {
	req, err := c.preparerForListByVolumeGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumes.VolumesClient", "ListByVolumeGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumes.VolumesClient", "ListByVolumeGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByVolumeGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumes.VolumesClient", "ListByVolumeGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByVolumeGroup prepares the ListByVolumeGroup request.
func (c VolumesClient) preparerForListByVolumeGroup(ctx context.Context, id VolumeGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/volumes", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByVolumeGroupWithNextLink prepares the ListByVolumeGroup request with the given nextLink token.
func (c VolumesClient) preparerForListByVolumeGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByVolumeGroup handles the response to the ListByVolumeGroup request. The method always
// closes the http.Response Body.
func (c VolumesClient) responderForListByVolumeGroup(resp *http.Response) (result ListByVolumeGroupOperationResponse, err error) {
	type page struct {
		Values   []Volume `json:"value"`
		NextLink *string  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByVolumeGroupOperationResponse, err error) {
			req, err := c.preparerForListByVolumeGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "volumes.VolumesClient", "ListByVolumeGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "volumes.VolumesClient", "ListByVolumeGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByVolumeGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "volumes.VolumesClient", "ListByVolumeGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByVolumeGroupComplete retrieves all of the results into a single object
func (c VolumesClient) ListByVolumeGroupComplete(ctx context.Context, id VolumeGroupId) (ListByVolumeGroupCompleteResult, error) {
	return c.ListByVolumeGroupCompleteMatchingPredicate(ctx, id, VolumeOperationPredicate{})
}

// ListByVolumeGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VolumesClient) ListByVolumeGroupCompleteMatchingPredicate(ctx context.Context, id VolumeGroupId, predicate VolumeOperationPredicate) (resp ListByVolumeGroupCompleteResult, err error) {
	items := make([]Volume, 0)

	page, err := c.ListByVolumeGroup(ctx, id)
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

	out := ListByVolumeGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
