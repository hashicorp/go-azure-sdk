package policydefinitions

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

type ListByManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PolicyDefinition

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByManagementGroupOperationResponse, error)
}

type ListByManagementGroupCompleteResult struct {
	Items []PolicyDefinition
}

func (r ListByManagementGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByManagementGroupOperationResponse) LoadMore(ctx context.Context) (resp ListByManagementGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByManagementGroupOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultListByManagementGroupOperationOptions() ListByManagementGroupOperationOptions {
	return ListByManagementGroupOperationOptions{}
}

func (o ListByManagementGroupOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByManagementGroupOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByManagementGroup ...
func (c PolicyDefinitionsClient) ListByManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options ListByManagementGroupOperationOptions) (resp ListByManagementGroupOperationResponse, err error) {
	req, err := c.preparerForListByManagementGroup(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListByManagementGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByManagementGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListByManagementGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByManagementGroup prepares the ListByManagementGroup request.
func (c PolicyDefinitionsClient) preparerForListByManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options ListByManagementGroupOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Authorization/policyDefinitions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByManagementGroupWithNextLink prepares the ListByManagementGroup request with the given nextLink token.
func (c PolicyDefinitionsClient) preparerForListByManagementGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByManagementGroup handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (c PolicyDefinitionsClient) responderForListByManagementGroup(resp *http.Response) (result ListByManagementGroupOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByManagementGroupOperationResponse, err error) {
			req, err := c.preparerForListByManagementGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListByManagementGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByManagementGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "ListByManagementGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByManagementGroupComplete retrieves all of the results into a single object
func (c PolicyDefinitionsClient) ListByManagementGroupComplete(ctx context.Context, id commonids.ManagementGroupId, options ListByManagementGroupOperationOptions) (ListByManagementGroupCompleteResult, error) {
	return c.ListByManagementGroupCompleteMatchingPredicate(ctx, id, options, PolicyDefinitionOperationPredicate{})
}

// ListByManagementGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PolicyDefinitionsClient) ListByManagementGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ManagementGroupId, options ListByManagementGroupOperationOptions, predicate PolicyDefinitionOperationPredicate) (resp ListByManagementGroupCompleteResult, err error) {
	items := make([]PolicyDefinition, 0)

	page, err := c.ListByManagementGroup(ctx, id, options)
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

	out := ListByManagementGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
