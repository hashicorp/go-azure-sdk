package registrymanagement

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

type RegistriesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RegistryTrackedResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistriesListBySubscriptionOperationResponse, error)
}

type RegistriesListBySubscriptionCompleteResult struct {
	Items []RegistryTrackedResource
}

func (r RegistriesListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistriesListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp RegistriesListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// RegistriesListBySubscription ...
func (c RegistryManagementClient) RegistriesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp RegistriesListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForRegistriesListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistriesListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistriesListBySubscription prepares the RegistriesListBySubscription request.
func (c RegistryManagementClient) preparerForRegistriesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.MachineLearningServices/registries", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistriesListBySubscriptionWithNextLink prepares the RegistriesListBySubscription request with the given nextLink token.
func (c RegistryManagementClient) preparerForRegistriesListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistriesListBySubscription handles the response to the RegistriesListBySubscription request. The method always
// closes the http.Response Body.
func (c RegistryManagementClient) responderForRegistriesListBySubscription(resp *http.Response) (result RegistriesListBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []RegistryTrackedResource `json:"value"`
		NextLink *string                   `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistriesListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForRegistriesListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistriesListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistriesListBySubscriptionComplete retrieves all of the results into a single object
func (c RegistryManagementClient) RegistriesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (RegistriesListBySubscriptionCompleteResult, error) {
	return c.RegistriesListBySubscriptionCompleteMatchingPredicate(ctx, id, RegistryTrackedResourceOperationPredicate{})
}

// RegistriesListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RegistryManagementClient) RegistriesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate RegistryTrackedResourceOperationPredicate) (resp RegistriesListBySubscriptionCompleteResult, err error) {
	items := make([]RegistryTrackedResource, 0)

	page, err := c.RegistriesListBySubscription(ctx, id)
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

	out := RegistriesListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
