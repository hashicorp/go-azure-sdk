package skus

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiManagementSkusListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApiManagementSku

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ApiManagementSkusListOperationResponse, error)
}

type ApiManagementSkusListCompleteResult struct {
	Items []ApiManagementSku
}

func (r ApiManagementSkusListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ApiManagementSkusListOperationResponse) LoadMore(ctx context.Context) (resp ApiManagementSkusListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ApiManagementSkusList ...
func (c SkusClient) ApiManagementSkusList(ctx context.Context, id commonids.SubscriptionId) (resp ApiManagementSkusListOperationResponse, err error) {
	req, err := c.preparerForApiManagementSkusList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "skus.SkusClient", "ApiManagementSkusList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "skus.SkusClient", "ApiManagementSkusList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForApiManagementSkusList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "skus.SkusClient", "ApiManagementSkusList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForApiManagementSkusList prepares the ApiManagementSkusList request.
func (c SkusClient) preparerForApiManagementSkusList(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.ApiManagement/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForApiManagementSkusListWithNextLink prepares the ApiManagementSkusList request with the given nextLink token.
func (c SkusClient) preparerForApiManagementSkusListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForApiManagementSkusList handles the response to the ApiManagementSkusList request. The method always
// closes the http.Response Body.
func (c SkusClient) responderForApiManagementSkusList(resp *http.Response) (result ApiManagementSkusListOperationResponse, err error) {
	type page struct {
		Values   []ApiManagementSku `json:"value"`
		NextLink *string            `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ApiManagementSkusListOperationResponse, err error) {
			req, err := c.preparerForApiManagementSkusListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "skus.SkusClient", "ApiManagementSkusList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "skus.SkusClient", "ApiManagementSkusList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForApiManagementSkusList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "skus.SkusClient", "ApiManagementSkusList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ApiManagementSkusListComplete retrieves all of the results into a single object
func (c SkusClient) ApiManagementSkusListComplete(ctx context.Context, id commonids.SubscriptionId) (ApiManagementSkusListCompleteResult, error) {
	return c.ApiManagementSkusListCompleteMatchingPredicate(ctx, id, ApiManagementSkuOperationPredicate{})
}

// ApiManagementSkusListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SkusClient) ApiManagementSkusListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ApiManagementSkuOperationPredicate) (resp ApiManagementSkusListCompleteResult, err error) {
	items := make([]ApiManagementSku, 0)

	page, err := c.ApiManagementSkusList(ctx, id)
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

	out := ApiManagementSkusListCompleteResult{
		Items: items,
	}
	return out, nil
}
