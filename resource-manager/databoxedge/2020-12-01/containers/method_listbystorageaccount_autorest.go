package containers

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

type ListByStorageAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Container

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByStorageAccountOperationResponse, error)
}

type ListByStorageAccountCompleteResult struct {
	Items []Container
}

func (r ListByStorageAccountOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByStorageAccountOperationResponse) LoadMore(ctx context.Context) (resp ListByStorageAccountOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByStorageAccount ...
func (c ContainersClient) ListByStorageAccount(ctx context.Context, id StorageAccountId) (resp ListByStorageAccountOperationResponse, err error) {
	req, err := c.preparerForListByStorageAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containers.ContainersClient", "ListByStorageAccount", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "containers.ContainersClient", "ListByStorageAccount", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByStorageAccount(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containers.ContainersClient", "ListByStorageAccount", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByStorageAccount prepares the ListByStorageAccount request.
func (c ContainersClient) preparerForListByStorageAccount(ctx context.Context, id StorageAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/containers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByStorageAccountWithNextLink prepares the ListByStorageAccount request with the given nextLink token.
func (c ContainersClient) preparerForListByStorageAccountWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByStorageAccount handles the response to the ListByStorageAccount request. The method always
// closes the http.Response Body.
func (c ContainersClient) responderForListByStorageAccount(resp *http.Response) (result ListByStorageAccountOperationResponse, err error) {
	type page struct {
		Values   []Container `json:"value"`
		NextLink *string     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByStorageAccountOperationResponse, err error) {
			req, err := c.preparerForListByStorageAccountWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "containers.ContainersClient", "ListByStorageAccount", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "containers.ContainersClient", "ListByStorageAccount", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByStorageAccount(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "containers.ContainersClient", "ListByStorageAccount", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByStorageAccountComplete retrieves all of the results into a single object
func (c ContainersClient) ListByStorageAccountComplete(ctx context.Context, id StorageAccountId) (ListByStorageAccountCompleteResult, error) {
	return c.ListByStorageAccountCompleteMatchingPredicate(ctx, id, ContainerOperationPredicate{})
}

// ListByStorageAccountCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ContainersClient) ListByStorageAccountCompleteMatchingPredicate(ctx context.Context, id StorageAccountId, predicate ContainerOperationPredicate) (resp ListByStorageAccountCompleteResult, err error) {
	items := make([]Container, 0)

	page, err := c.ListByStorageAccount(ctx, id)
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

	out := ListByStorageAccountCompleteResult{
		Items: items,
	}
	return out, nil
}
