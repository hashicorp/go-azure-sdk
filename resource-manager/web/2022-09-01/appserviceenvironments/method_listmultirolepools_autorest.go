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

type ListMultiRolePoolsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]WorkerPoolResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListMultiRolePoolsOperationResponse, error)
}

type ListMultiRolePoolsCompleteResult struct {
	Items []WorkerPoolResource
}

func (r ListMultiRolePoolsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListMultiRolePoolsOperationResponse) LoadMore(ctx context.Context) (resp ListMultiRolePoolsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListMultiRolePools ...
func (c AppServiceEnvironmentsClient) ListMultiRolePools(ctx context.Context, id HostingEnvironmentId) (resp ListMultiRolePoolsOperationResponse, err error) {
	req, err := c.preparerForListMultiRolePools(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePools", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePools", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListMultiRolePools(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePools", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListMultiRolePools prepares the ListMultiRolePools request.
func (c AppServiceEnvironmentsClient) preparerForListMultiRolePools(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListMultiRolePoolsWithNextLink prepares the ListMultiRolePools request with the given nextLink token.
func (c AppServiceEnvironmentsClient) preparerForListMultiRolePoolsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListMultiRolePools handles the response to the ListMultiRolePools request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForListMultiRolePools(resp *http.Response) (result ListMultiRolePoolsOperationResponse, err error) {
	type page struct {
		Values   []WorkerPoolResource `json:"value"`
		NextLink *string              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListMultiRolePoolsOperationResponse, err error) {
			req, err := c.preparerForListMultiRolePoolsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePools", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePools", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListMultiRolePools(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "ListMultiRolePools", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListMultiRolePoolsComplete retrieves all of the results into a single object
func (c AppServiceEnvironmentsClient) ListMultiRolePoolsComplete(ctx context.Context, id HostingEnvironmentId) (ListMultiRolePoolsCompleteResult, error) {
	return c.ListMultiRolePoolsCompleteMatchingPredicate(ctx, id, WorkerPoolResourceOperationPredicate{})
}

// ListMultiRolePoolsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServiceEnvironmentsClient) ListMultiRolePoolsCompleteMatchingPredicate(ctx context.Context, id HostingEnvironmentId, predicate WorkerPoolResourceOperationPredicate) (resp ListMultiRolePoolsCompleteResult, err error) {
	items := make([]WorkerPoolResource, 0)

	page, err := c.ListMultiRolePools(ctx, id)
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

	out := ListMultiRolePoolsCompleteResult{
		Items: items,
	}
	return out, nil
}
