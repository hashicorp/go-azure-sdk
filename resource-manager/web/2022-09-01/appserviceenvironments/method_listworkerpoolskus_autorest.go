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

type ListWorkerPoolSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SkuInfo

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListWorkerPoolSkusOperationResponse, error)
}

type ListWorkerPoolSkusCompleteResult struct {
	Items []SkuInfo
}

func (r ListWorkerPoolSkusOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListWorkerPoolSkusOperationResponse) LoadMore(ctx context.Context) (resp ListWorkerPoolSkusOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListWorkerPoolSkus ...
func (c AppServiceEnvironmentsClient) ListWorkerPoolSkus(ctx context.Context, id WorkerPoolId) (resp ListWorkerPoolSkusOperationResponse, err error) {
	req, err := c.preparerForListWorkerPoolSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWorkerPoolSkus", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWorkerPoolSkus", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListWorkerPoolSkus(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWorkerPoolSkus", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListWorkerPoolSkus prepares the ListWorkerPoolSkus request.
func (c AppServiceEnvironmentsClient) preparerForListWorkerPoolSkus(ctx context.Context, id WorkerPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWorkerPoolSkusWithNextLink prepares the ListWorkerPoolSkus request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListWorkerPoolSkusWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListWorkerPoolSkus handles the response to the ListWorkerPoolSkus request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListWorkerPoolSkus(resp *http.Response) (result ListWorkerPoolSkusOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListWorkerPoolSkusOperationResponse, err error) {
			req, err := c.preparerForListWorkerPoolSkusWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWorkerPoolSkus", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWorkerPoolSkus", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListWorkerPoolSkus(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWorkerPoolSkus", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListWorkerPoolSkusComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListWorkerPoolSkusComplete(ctx context.Context, id WorkerPoolId) (ListWorkerPoolSkusCompleteResult, error) {
	return c.ListWorkerPoolSkusCompleteMatchingPredicate(ctx, id, SkuInfoOperationPredicate{})
}

// ListWorkerPoolSkusCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListWorkerPoolSkusCompleteMatchingPredicate(ctx context.Context, id WorkerPoolId, predicate SkuInfoOperationPredicate) (resp ListWorkerPoolSkusCompleteResult, err error) {
	items := make([]SkuInfo, 0)

	page, err := c.ListWorkerPoolSkus(ctx, id)
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

	out := ListWorkerPoolSkusCompleteResult{
		Items: items,
	}
	return out, nil
}
