package resourcehealthmetadata

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

type ListBySiteSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ResourceHealthMetadata

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBySiteSlotOperationResponse, error)
}

type ListBySiteSlotCompleteResult struct {
	Items []ResourceHealthMetadata
}

func (r ListBySiteSlotOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBySiteSlotOperationResponse) LoadMore(ctx context.Context) (resp ListBySiteSlotOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBySiteSlot ...
func (c ResourceHealthMetadataClient) ListBySiteSlot(ctx context.Context, id SlotId) (resp ListBySiteSlotOperationResponse, err error) {
	req, err := c.preparerForListBySiteSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySiteSlot", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySiteSlot", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBySiteSlot(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySiteSlot", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBySiteSlot prepares the ListBySiteSlot request.
func (c ResourceHealthMetadataClient) preparerForListBySiteSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resourceHealthMetadata", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBySiteSlotWithNextLink prepares the ListBySiteSlot request with the given nextLink token.
func (c ResourceHealthMetadataClient) preparerForListBySiteSlotWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBySiteSlot handles the response to the ListBySiteSlot request. The method always
// closes the http.Response Body.
func (c ResourceHealthMetadataClient) responderForListBySiteSlot(resp *http.Response) (result ListBySiteSlotOperationResponse, err error) {
	type page struct {
		Values   []ResourceHealthMetadata `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBySiteSlotOperationResponse, err error) {
			req, err := c.preparerForListBySiteSlotWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySiteSlot", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySiteSlot", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBySiteSlot(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySiteSlot", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBySiteSlotComplete retrieves all of the results into a single object
func (c ResourceHealthMetadataClient) ListBySiteSlotComplete(ctx context.Context, id SlotId) (ListBySiteSlotCompleteResult, error) {
	return c.ListBySiteSlotCompleteMatchingPredicate(ctx, id, ResourceHealthMetadataOperationPredicate{})
}

// ListBySiteSlotCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceHealthMetadataClient) ListBySiteSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ResourceHealthMetadataOperationPredicate) (resp ListBySiteSlotCompleteResult, err error) {
	items := make([]ResourceHealthMetadata, 0)

	page, err := c.ListBySiteSlot(ctx, id)
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

	out := ListBySiteSlotCompleteResult{
		Items: items,
	}
	return out, nil
}
