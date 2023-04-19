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

type RegistriesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RegistryTrackedResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistriesListOperationResponse, error)
}

type RegistriesListCompleteResult struct {
	Items []RegistryTrackedResource
}

func (r RegistriesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistriesListOperationResponse) LoadMore(ctx context.Context) (resp RegistriesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// RegistriesList ...
func (c RegistryManagementClient) RegistriesList(ctx context.Context, id commonids.ResourceGroupId) (resp RegistriesListOperationResponse, err error) {
	req, err := c.preparerForRegistriesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistriesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistriesList prepares the RegistriesList request.
func (c RegistryManagementClient) preparerForRegistriesList(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
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

// preparerForRegistriesListWithNextLink prepares the RegistriesList request with the given nextLink token.
func (c RegistryManagementClient) preparerForRegistriesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRegistriesList handles the response to the RegistriesList request. The method always
// closes the http.Response Body.
func (c RegistryManagementClient) responderForRegistriesList(resp *http.Response) (result RegistriesListOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistriesListOperationResponse, err error) {
			req, err := c.preparerForRegistriesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistriesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistriesListComplete retrieves all of the results into a single object
func (c RegistryManagementClient) RegistriesListComplete(ctx context.Context, id commonids.ResourceGroupId) (RegistriesListCompleteResult, error) {
	return c.RegistriesListCompleteMatchingPredicate(ctx, id, RegistryTrackedResourceOperationPredicate{})
}

// RegistriesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RegistryManagementClient) RegistriesListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate RegistryTrackedResourceOperationPredicate) (resp RegistriesListCompleteResult, err error) {
	items := make([]RegistryTrackedResource, 0)

	page, err := c.RegistriesList(ctx, id)
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

	out := RegistriesListCompleteResult{
		Items: items,
	}
	return out, nil
}
