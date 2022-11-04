package marketplaceregistrationdefinitions

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

type WithoutScopeListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]MarketplaceRegistrationDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (WithoutScopeListOperationResponse, error)
}

type WithoutScopeListCompleteResult struct {
	Items []MarketplaceRegistrationDefinition
}

func (r WithoutScopeListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r WithoutScopeListOperationResponse) LoadMore(ctx context.Context) (resp WithoutScopeListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type WithoutScopeListOperationOptions struct {
	Filter *string
}

func DefaultWithoutScopeListOperationOptions() WithoutScopeListOperationOptions {
	return WithoutScopeListOperationOptions{}
}

func (o WithoutScopeListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o WithoutScopeListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// WithoutScopeList ...
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeList(ctx context.Context, options WithoutScopeListOperationOptions) (resp WithoutScopeListOperationResponse, err error) {
	req, err := c.preparerForWithoutScopeList(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForWithoutScopeList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForWithoutScopeList prepares the WithoutScopeList request.
func (c MarketplaceRegistrationDefinitionsClient) preparerForWithoutScopeList(ctx context.Context, options WithoutScopeListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath("/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForWithoutScopeListWithNextLink prepares the WithoutScopeList request with the given nextLink token.
func (c MarketplaceRegistrationDefinitionsClient) preparerForWithoutScopeListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForWithoutScopeList handles the response to the WithoutScopeList request. The method always
// closes the http.Response Body.
func (c MarketplaceRegistrationDefinitionsClient) responderForWithoutScopeList(resp *http.Response) (result WithoutScopeListOperationResponse, err error) {
	type page struct {
		Values   []MarketplaceRegistrationDefinition `json:"value"`
		NextLink *string                             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result WithoutScopeListOperationResponse, err error) {
			req, err := c.preparerForWithoutScopeListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForWithoutScopeList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// WithoutScopeListComplete retrieves all of the results into a single object
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeListComplete(ctx context.Context, options WithoutScopeListOperationOptions) (WithoutScopeListCompleteResult, error) {
	return c.WithoutScopeListCompleteMatchingPredicate(ctx, options, MarketplaceRegistrationDefinitionOperationPredicate{})
}

// WithoutScopeListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeListCompleteMatchingPredicate(ctx context.Context, options WithoutScopeListOperationOptions, predicate MarketplaceRegistrationDefinitionOperationPredicate) (resp WithoutScopeListCompleteResult, err error) {
	items := make([]MarketplaceRegistrationDefinition, 0)

	page, err := c.WithoutScopeList(ctx, options)
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

	out := WithoutScopeListCompleteResult{
		Items: items,
	}
	return out, nil
}
