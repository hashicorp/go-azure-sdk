package galleryimageversions

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

type ListByGalleryImageOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]GalleryImageVersion

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByGalleryImageOperationResponse, error)
}

type ListByGalleryImageCompleteResult struct {
	Items []GalleryImageVersion
}

func (r ListByGalleryImageOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByGalleryImageOperationResponse) LoadMore(ctx context.Context) (resp ListByGalleryImageOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByGalleryImage ...
func (c GalleryImageVersionsClient) ListByGalleryImage(ctx context.Context, id GalleryImageId) (resp ListByGalleryImageOperationResponse, err error) {
	req, err := c.preparerForListByGalleryImage(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "galleryimageversions.GalleryImageVersionsClient", "ListByGalleryImage", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "galleryimageversions.GalleryImageVersionsClient", "ListByGalleryImage", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByGalleryImage(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "galleryimageversions.GalleryImageVersionsClient", "ListByGalleryImage", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByGalleryImage prepares the ListByGalleryImage request.
func (c GalleryImageVersionsClient) preparerForListByGalleryImage(ctx context.Context, id GalleryImageId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByGalleryImageWithNextLink prepares the ListByGalleryImage request with the given nextLink token.
func (c GalleryImageVersionsClient) preparerForListByGalleryImageWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByGalleryImage handles the response to the ListByGalleryImage request. The method always
// closes the http.Response Body.
func (c GalleryImageVersionsClient) responderForListByGalleryImage(resp *http.Response) (result ListByGalleryImageOperationResponse, err error) {
	type page struct {
		Values   []GalleryImageVersion `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByGalleryImageOperationResponse, err error) {
			req, err := c.preparerForListByGalleryImageWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "galleryimageversions.GalleryImageVersionsClient", "ListByGalleryImage", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "galleryimageversions.GalleryImageVersionsClient", "ListByGalleryImage", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByGalleryImage(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "galleryimageversions.GalleryImageVersionsClient", "ListByGalleryImage", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByGalleryImageComplete retrieves all of the results into a single object
func (c GalleryImageVersionsClient) ListByGalleryImageComplete(ctx context.Context, id GalleryImageId) (ListByGalleryImageCompleteResult, error) {
	return c.ListByGalleryImageCompleteMatchingPredicate(ctx, id, GalleryImageVersionOperationPredicate{})
}

// ListByGalleryImageCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c GalleryImageVersionsClient) ListByGalleryImageCompleteMatchingPredicate(ctx context.Context, id GalleryImageId, predicate GalleryImageVersionOperationPredicate) (resp ListByGalleryImageCompleteResult, err error) {
	items := make([]GalleryImageVersion, 0)

	page, err := c.ListByGalleryImage(ctx, id)
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

	out := ListByGalleryImageCompleteResult{
		Items: items,
	}
	return out, nil
}
