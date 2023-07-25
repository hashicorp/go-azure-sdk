package provider

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

type GetAvailableStacksOnPremOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ApplicationStackResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAvailableStacksOnPremOperationResponse, error)
}

type GetAvailableStacksOnPremCompleteResult struct {
	Items []ApplicationStackResource
}

func (r GetAvailableStacksOnPremOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAvailableStacksOnPremOperationResponse) LoadMore(ctx context.Context) (resp GetAvailableStacksOnPremOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type GetAvailableStacksOnPremOperationOptions struct {
	OsTypeSelected *ProviderOsTypeSelected
}

func DefaultGetAvailableStacksOnPremOperationOptions() GetAvailableStacksOnPremOperationOptions {
	return GetAvailableStacksOnPremOperationOptions{}
}

func (o GetAvailableStacksOnPremOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetAvailableStacksOnPremOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.OsTypeSelected != nil {
		out["osTypeSelected"] = *o.OsTypeSelected
	}

	return out
}

// GetAvailableStacksOnPrem ...
func (c ProviderClient) GetAvailableStacksOnPrem(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions) (resp GetAvailableStacksOnPremOperationResponse, err error) {
	req, err := c.preparerForGetAvailableStacksOnPrem(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacksOnPrem", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacksOnPrem", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAvailableStacksOnPrem(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacksOnPrem", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAvailableStacksOnPrem prepares the GetAvailableStacksOnPrem request.
func (c ProviderClient) preparerForGetAvailableStacksOnPrem(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/availableStacks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAvailableStacksOnPremWithNextLink prepares the GetAvailableStacksOnPrem request with the given nextLink token.
func (c ProviderClient) preparerForGetAvailableStacksOnPremWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetAvailableStacksOnPrem handles the response to the GetAvailableStacksOnPrem request. The method always
// closes the http.Response Body.
func (c ProviderClient) responderForGetAvailableStacksOnPrem(resp *http.Response) (result GetAvailableStacksOnPremOperationResponse, err error) {
	type page struct {
		Values   []ApplicationStackResource `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAvailableStacksOnPremOperationResponse, err error) {
			req, err := c.preparerForGetAvailableStacksOnPremWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacksOnPrem", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacksOnPrem", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAvailableStacksOnPrem(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "provider.ProviderClient", "GetAvailableStacksOnPrem", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAvailableStacksOnPremComplete retrieves all of the results into a single object
func (c ProviderClient) GetAvailableStacksOnPremComplete(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions) (GetAvailableStacksOnPremCompleteResult, error) {
	return c.GetAvailableStacksOnPremCompleteMatchingPredicate(ctx, id, options, ApplicationStackResourceOperationPredicate{})
}

// GetAvailableStacksOnPremCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProviderClient) GetAvailableStacksOnPremCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options GetAvailableStacksOnPremOperationOptions, predicate ApplicationStackResourceOperationPredicate) (resp GetAvailableStacksOnPremCompleteResult, err error) {
	items := make([]ApplicationStackResource, 0)

	page, err := c.GetAvailableStacksOnPrem(ctx, id, options)
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

	out := GetAvailableStacksOnPremCompleteResult{
		Items: items,
	}
	return out, nil
}
