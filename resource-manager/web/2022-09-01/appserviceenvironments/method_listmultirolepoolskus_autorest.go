package appserviceenvironments

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

type ListMultiRolePoolSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SkuInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListMultiRolePoolSkusOperationResponse, error)
}

type ListMultiRolePoolSkusCompleteResult struct {
	Items []SkuInfo
}

func (r ListMultiRolePoolSkusOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListMultiRolePoolSkusOperationResponse) LoadMore(ctx context.Context) (resp ListMultiRolePoolSkusOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListMultiRolePoolSkus ...
func (c AppServiceEnvironmentsClient) ListMultiRolePoolSkus(ctx context.Context, id HostingEnvironmentId) (resp ListMultiRolePoolSkusOperationResponse, err error) {
	req, err := c.preparerForListMultiRolePoolSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolSkus", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolSkus", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListMultiRolePoolSkus(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolSkus", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListMultiRolePoolSkus prepares the ListMultiRolePoolSkus request.
func (c AppServiceEnvironmentsClient) preparerForListMultiRolePoolSkus(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools/default/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListMultiRolePoolSkusWithNextLink prepares the ListMultiRolePoolSkus request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListMultiRolePoolSkusWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListMultiRolePoolSkus handles the response to the ListMultiRolePoolSkus request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListMultiRolePoolSkus(resp *http.Response) (result ListMultiRolePoolSkusOperationResponse, err error) {
	type page struct {
		Values   []SkuInfo `json:"value"`
		NextLink *string   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListMultiRolePoolSkusOperationResponse, err error) {
			req, err := c.preparerForListMultiRolePoolSkusWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolSkus", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolSkus", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListMultiRolePoolSkus(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePoolSkus", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListMultiRolePoolSkusComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRolePoolSkusComplete(ctx context.Context, id HostingEnvironmentId) (ListMultiRolePoolSkusCompleteResult, error) {
	return c.ListMultiRolePoolSkusCompleteMatchingPredicate(ctx, id, SkuInfoOperationPredicate{})
}

// ListMultiRolePoolSkusCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListMultiRolePoolSkusCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate SkuInfoOperationPredicate) (resp ListMultiRolePoolSkusCompleteResult, err error) {
	items := make([]SkuInfo, 0)

	page, err := c.ListMultiRolePoolSkus(ctx, id)
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

	out := ListMultiRolePoolSkusCompleteResult{
		Items: items,
	}
	return out, nil
}
