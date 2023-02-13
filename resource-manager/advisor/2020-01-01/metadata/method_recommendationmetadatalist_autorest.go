package metadata

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

type RecommendationMetadataListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]MetadataEntity

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RecommendationMetadataListOperationResponse, error)
}

type RecommendationMetadataListCompleteResult struct {
	Items []MetadataEntity
}

func (r RecommendationMetadataListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RecommendationMetadataListOperationResponse) LoadMore(ctx context.Context) (resp RecommendationMetadataListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// RecommendationMetadataList ...
func (c MetadataClient) RecommendationMetadataList(ctx context.Context) (resp RecommendationMetadataListOperationResponse, err error) {
	req, err := c.preparerForRecommendationMetadataList(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRecommendationMetadataList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRecommendationMetadataList prepares the RecommendationMetadataList request.
func (c MetadataClient) preparerForRecommendationMetadataList(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Advisor/metadata"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRecommendationMetadataListWithNextLink prepares the RecommendationMetadataList request with the given nextLink token.
func (c MetadataClient) preparerForRecommendationMetadataListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRecommendationMetadataList handles the response to the RecommendationMetadataList request. The method always
// closes the http.Response Body.
func (c MetadataClient) responderForRecommendationMetadataList(resp *http.Response) (result RecommendationMetadataListOperationResponse, err error) {
	type page struct {
		Values   []MetadataEntity `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RecommendationMetadataListOperationResponse, err error) {
			req, err := c.preparerForRecommendationMetadataListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRecommendationMetadataList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "metadata.MetadataClient", "RecommendationMetadataList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RecommendationMetadataListComplete retrieves all of the results into a single object
func (c MetadataClient) RecommendationMetadataListComplete(ctx context.Context) (RecommendationMetadataListCompleteResult, error) {
	return c.RecommendationMetadataListCompleteMatchingPredicate(ctx, MetadataEntityOperationPredicate{})
}

// RecommendationMetadataListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MetadataClient) RecommendationMetadataListCompleteMatchingPredicate(ctx context.Context, predicate MetadataEntityOperationPredicate) (resp RecommendationMetadataListCompleteResult, err error) {
	items := make([]MetadataEntity, 0)

	page, err := c.RecommendationMetadataList(ctx)
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

	out := RecommendationMetadataListCompleteResult{
		Items: items,
	}
	return out, nil
}
