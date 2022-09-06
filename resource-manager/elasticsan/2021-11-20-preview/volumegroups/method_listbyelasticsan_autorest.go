package volumegroups

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

type ListByElasticSanOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VolumeGroup

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByElasticSanOperationResponse, error)
}

type ListByElasticSanCompleteResult struct {
	Items []VolumeGroup
}

func (r ListByElasticSanOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByElasticSanOperationResponse) LoadMore(ctx context.Context) (resp ListByElasticSanOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByElasticSan ...
func (c VolumeGroupsClient) ListByElasticSan(ctx context.Context, id ElasticSanId) (resp ListByElasticSanOperationResponse, err error) {
	req, err := c.preparerForListByElasticSan(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumegroups.VolumeGroupsClient", "ListByElasticSan", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumegroups.VolumeGroupsClient", "ListByElasticSan", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByElasticSan(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "volumegroups.VolumeGroupsClient", "ListByElasticSan", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByElasticSan prepares the ListByElasticSan request.
func (c VolumeGroupsClient) preparerForListByElasticSan(ctx context.Context, id ElasticSanId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/volumeGroups", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByElasticSanWithNextLink prepares the ListByElasticSan request with the given nextLink token.
func (c VolumeGroupsClient) preparerForListByElasticSanWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByElasticSan handles the response to the ListByElasticSan request. The method always
// closes the http.Response Body.
func (c VolumeGroupsClient) responderForListByElasticSan(resp *http.Response) (result ListByElasticSanOperationResponse, err error) {
	type page struct {
		Values   []VolumeGroup `json:"value"`
		NextLink *string       `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByElasticSanOperationResponse, err error) {
			req, err := c.preparerForListByElasticSanWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "volumegroups.VolumeGroupsClient", "ListByElasticSan", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "volumegroups.VolumeGroupsClient", "ListByElasticSan", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByElasticSan(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "volumegroups.VolumeGroupsClient", "ListByElasticSan", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByElasticSanComplete retrieves all of the results into a single object
func (c VolumeGroupsClient) ListByElasticSanComplete(ctx context.Context, id ElasticSanId) (ListByElasticSanCompleteResult, error) {
	return c.ListByElasticSanCompleteMatchingPredicate(ctx, id, VolumeGroupOperationPredicate{})
}

// ListByElasticSanCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VolumeGroupsClient) ListByElasticSanCompleteMatchingPredicate(ctx context.Context, id ElasticSanId, predicate VolumeGroupOperationPredicate) (resp ListByElasticSanCompleteResult, err error) {
	items := make([]VolumeGroup, 0)

	page, err := c.ListByElasticSan(ctx, id)
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

	out := ListByElasticSanCompleteResult{
		Items: items,
	}
	return out, nil
}
