package subvolumes

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

type ListByVolumeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SubvolumeInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByVolumeOperationResponse, error)
}

type ListByVolumeCompleteResult struct {
	Items []SubvolumeInfo
}

func (r ListByVolumeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByVolumeOperationResponse) LoadMore(ctx context.Context) (resp ListByVolumeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByVolume ...
func (c SubVolumesClient) ListByVolume(ctx context.Context, id VolumeId) (resp ListByVolumeOperationResponse, err error) {
	req, err := c.preparerForListByVolume(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subvolumes.SubVolumesClient", "ListByVolume", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subvolumes.SubVolumesClient", "ListByVolume", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByVolume(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subvolumes.SubVolumesClient", "ListByVolume", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListByVolumeComplete retrieves all of the results into a single object
func (c SubVolumesClient) ListByVolumeComplete(ctx context.Context, id VolumeId) (ListByVolumeCompleteResult, error) {
	return c.ListByVolumeCompleteMatchingPredicate(ctx, id, SubvolumeInfoOperationPredicate{})
}

// ListByVolumeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SubVolumesClient) ListByVolumeCompleteMatchingPredicate(ctx context.Context, id VolumeId, predicate SubvolumeInfoOperationPredicate) (resp ListByVolumeCompleteResult, err error) {
	items := make([]SubvolumeInfo, 0)

	page, err := c.ListByVolume(ctx, id)
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

	out := ListByVolumeCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListByVolume prepares the ListByVolume request.
func (c SubVolumesClient) preparerForListByVolume(ctx context.Context, id VolumeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/subVolumes", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByVolumeWithNextLink prepares the ListByVolume request with the given nextLink token.
func (c SubVolumesClient) preparerForListByVolumeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByVolume handles the response to the ListByVolume request. The method always
// closes the http.Response Body.
func (c SubVolumesClient) responderForListByVolume(resp *http.Response) (result ListByVolumeOperationResponse, err error) {
	type page struct {
		Values   []SubvolumeInfo `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByVolumeOperationResponse, err error) {
			req, err := c.preparerForListByVolumeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subvolumes.SubVolumesClient", "ListByVolume", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "subvolumes.SubVolumesClient", "ListByVolume", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByVolume(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subvolumes.SubVolumesClient", "ListByVolume", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
