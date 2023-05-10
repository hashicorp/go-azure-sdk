package policydefinitions

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

type ListBuiltInOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PolicyDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBuiltInOperationResponse, error)
}

type ListBuiltInCompleteResult struct {
	Items []PolicyDefinition
}

func (r ListBuiltInOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBuiltInOperationResponse) LoadMore(ctx context.Context) (resp ListBuiltInOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListBuiltInOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultListBuiltInOperationOptions() ListBuiltInOperationOptions {
	return ListBuiltInOperationOptions{}
}

func (o ListBuiltInOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListBuiltInOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListBuiltIn ...
func (c PolicyDefinitionsClient) ListBuiltIn(ctx context.Context, options ListBuiltInOperationOptions) (resp ListBuiltInOperationResponse, err error) {
	req, err := c.preparerForListBuiltIn(ctx, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListBuiltIn", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBuiltIn(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListBuiltIn", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBuiltIn prepares the ListBuiltIn request.
func (c PolicyDefinitionsClient) preparerForListBuiltIn(ctx context.Context, options ListBuiltInOperationOptions) (*http.Request, error) {
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
		autorest.WithPath("/providers/Microsoft.Authorization/policyDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBuiltInWithNextLink prepares the ListBuiltIn request with the given nextLink token.
func (c PolicyDefinitionsClient) preparerForListBuiltInWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBuiltIn handles the response to the ListBuiltIn request. The method always
// closes the http.Response Body.
func (c PolicyDefinitionsClient) responderForListBuiltIn(resp *http.Response) (result ListBuiltInOperationResponse, err error) {
	type page struct {
		Values   []PolicyDefinition `json:"value"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBuiltInOperationResponse, err error) {
			req, err := c.preparerForListBuiltInWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListBuiltIn", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBuiltIn(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListBuiltIn", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBuiltInComplete retrieves all of the results into a single object
func (c PolicyDefinitionsClient) ListBuiltInComplete(ctx context.Context, options ListBuiltInOperationOptions) (ListBuiltInCompleteResult, error) {
	return c.ListBuiltInCompleteMatchingPredicate(ctx, options, PolicyDefinitionOperationPredicate{})
}

// ListBuiltInCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PolicyDefinitionsClient) ListBuiltInCompleteMatchingPredicate(ctx context.Context, options ListBuiltInOperationOptions, predicate PolicyDefinitionOperationPredicate) (resp ListBuiltInCompleteResult, err error) {
	items := make([]PolicyDefinition, 0)

	page, err := c.ListBuiltIn(ctx, options)
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

	out := ListBuiltInCompleteResult{
		Items: items,
	}
	return out, nil
}
