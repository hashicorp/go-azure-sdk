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

type ListWebWorkerUsagesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Usage

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListWebWorkerUsagesOperationResponse, error)
}

type ListWebWorkerUsagesCompleteResult struct {
	Items []Usage
}

func (r ListWebWorkerUsagesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListWebWorkerUsagesOperationResponse) LoadMore(ctx context.Context) (resp ListWebWorkerUsagesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListWebWorkerUsages ...
func (c AppServiceEnvironmentsClient) ListWebWorkerUsages(ctx context.Context, id WorkerPoolId) (resp ListWebWorkerUsagesOperationResponse, err error) {
	req, err := c.preparerForListWebWorkerUsages(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebWorkerUsages", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebWorkerUsages", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListWebWorkerUsages(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebWorkerUsages", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListWebWorkerUsages prepares the ListWebWorkerUsages request.
func (c AppServiceEnvironmentsClient) preparerForListWebWorkerUsages(ctx context.Context, id WorkerPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWebWorkerUsagesWithNextLink prepares the ListWebWorkerUsages request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListWebWorkerUsagesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListWebWorkerUsages handles the response to the ListWebWorkerUsages request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListWebWorkerUsages(resp *http.Response) (result ListWebWorkerUsagesOperationResponse, err error) {
	type page struct {
		Values   []Usage `json:"value"`
		NextLink *string `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListWebWorkerUsagesOperationResponse, err error) {
			req, err := c.preparerForListWebWorkerUsagesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebWorkerUsages", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebWorkerUsages", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListWebWorkerUsages(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListWebWorkerUsages", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListWebWorkerUsagesComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListWebWorkerUsagesComplete(ctx context.Context, id WorkerPoolId) (ListWebWorkerUsagesCompleteResult, error) {
	return c.ListWebWorkerUsagesCompleteMatchingPredicate(ctx, id, UsageOperationPredicate{})
}

// ListWebWorkerUsagesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListWebWorkerUsagesCompleteMatchingPredicate(ctx context.Context, id WorkerPoolId, predicate UsageOperationPredicate) (resp ListWebWorkerUsagesCompleteResult, err error) {
	items := make([]Usage, 0)

	page, err := c.ListWebWorkerUsages(ctx, id)
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

	out := ListWebWorkerUsagesCompleteResult{
		Items: items,
	}
	return out, nil
}
