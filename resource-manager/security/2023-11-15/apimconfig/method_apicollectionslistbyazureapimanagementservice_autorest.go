package apimconfig

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

type APICollectionsListByAzureApiManagementServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiCollection

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (APICollectionsListByAzureApiManagementServiceOperationResponse, error)
}

type APICollectionsListByAzureApiManagementServiceCompleteResult struct {
	Items []ApiCollection
}

func (r APICollectionsListByAzureApiManagementServiceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r APICollectionsListByAzureApiManagementServiceOperationResponse) LoadMore(ctx context.Context) (resp APICollectionsListByAzureApiManagementServiceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// APICollectionsListByAzureApiManagementService ...
func (c APIMConfigClient) APICollectionsListByAzureApiManagementService(ctx context.Context, id ServiceId) (resp APICollectionsListByAzureApiManagementServiceOperationResponse, err error) {
	req, err := c.preparerForAPICollectionsListByAzureApiManagementService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimconfig.APIMConfigClient", "APICollectionsListByAzureApiManagementService", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimconfig.APIMConfigClient", "APICollectionsListByAzureApiManagementService", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAPICollectionsListByAzureApiManagementService(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimconfig.APIMConfigClient", "APICollectionsListByAzureApiManagementService", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAPICollectionsListByAzureApiManagementService prepares the APICollectionsListByAzureApiManagementService request.
func (c APIMConfigClient) preparerForAPICollectionsListByAzureApiManagementService(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/apiCollections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAPICollectionsListByAzureApiManagementServiceWithNextLink prepares the APICollectionsListByAzureApiManagementService request with the given nextLink token.
func (c APIMConfigClient) preparerForAPICollectionsListByAzureApiManagementServiceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAPICollectionsListByAzureApiManagementService handles the response to the APICollectionsListByAzureApiManagementService request. The method always
// closes the http.Response Body.
func (c APIMConfigClient) responderForAPICollectionsListByAzureApiManagementService(resp *http.Response) (result APICollectionsListByAzureApiManagementServiceOperationResponse, err error) {
	type page struct {
		Values   []ApiCollection `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result APICollectionsListByAzureApiManagementServiceOperationResponse, err error) {
			req, err := c.preparerForAPICollectionsListByAzureApiManagementServiceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apimconfig.APIMConfigClient", "APICollectionsListByAzureApiManagementService", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "apimconfig.APIMConfigClient", "APICollectionsListByAzureApiManagementService", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAPICollectionsListByAzureApiManagementService(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "apimconfig.APIMConfigClient", "APICollectionsListByAzureApiManagementService", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// APICollectionsListByAzureApiManagementServiceComplete retrieves all of the results into a single object
func (c APIMConfigClient) APICollectionsListByAzureApiManagementServiceComplete(ctx context.Context, id ServiceId) (APICollectionsListByAzureApiManagementServiceCompleteResult, error) {
	return c.APICollectionsListByAzureApiManagementServiceCompleteMatchingPredicate(ctx, id, ApiCollectionOperationPredicate{})
}

// APICollectionsListByAzureApiManagementServiceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c APIMConfigClient) APICollectionsListByAzureApiManagementServiceCompleteMatchingPredicate(ctx context.Context, id ServiceId, predicate ApiCollectionOperationPredicate) (resp APICollectionsListByAzureApiManagementServiceCompleteResult, err error) {
	items := make([]ApiCollection, 0)

	page, err := c.APICollectionsListByAzureApiManagementService(ctx, id)
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

	out := APICollectionsListByAzureApiManagementServiceCompleteResult{
		Items: items,
	}
	return out, nil
}
