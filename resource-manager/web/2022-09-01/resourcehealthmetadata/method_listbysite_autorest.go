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

type ListBySiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ResourceHealthMetadata

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBySiteOperationResponse, error)
}

type ListBySiteCompleteResult struct {
	Items []ResourceHealthMetadata
}

func (r ListBySiteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBySiteOperationResponse) LoadMore(ctx context.Context) (resp ListBySiteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListBySite ...
func (c ResourceHealthMetadataClient) ListBySite(ctx context.Context, id SiteId) (resp ListBySiteOperationResponse, err error) {
	req, err := c.preparerForListBySite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySite", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySite", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBySite(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySite", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBySite prepares the ListBySite request.
func (c ResourceHealthMetadataClient) preparerForListBySite(ctx context.Context, id SiteId) (*http.Request, error) {
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

// preparerForListBySiteWithNextLink prepares the ListBySite request with the given nextLink token.
func (c ResourceHealthMetadataClient) preparerForListBySiteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBySite handles the response to the ListBySite request. The method always
// closes the http.Response Body.
func (c ResourceHealthMetadataClient) responderForListBySite(resp *http.Response) (result ListBySiteOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBySiteOperationResponse, err error) {
			req, err := c.preparerForListBySiteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySite", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySite", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBySite(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "ListBySite", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBySiteComplete retrieves all of the results into a single object
func (c ResourceHealthMetadataClient) ListBySiteComplete(ctx context.Context, id SiteId) (ListBySiteCompleteResult, error) {
	return c.ListBySiteCompleteMatchingPredicate(ctx, id, ResourceHealthMetadataOperationPredicate{})
}

// ListBySiteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ResourceHealthMetadataClient) ListBySiteCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate ResourceHealthMetadataOperationPredicate) (resp ListBySiteCompleteResult, err error) {
	items := make([]ResourceHealthMetadata, 0)

	page, err := c.ListBySite(ctx, id)
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

	out := ListBySiteCompleteResult{
		Items: items,
	}
	return out, nil
}
