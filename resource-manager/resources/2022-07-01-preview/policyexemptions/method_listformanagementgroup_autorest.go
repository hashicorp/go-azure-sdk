package policyexemptions

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

type ListForManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PolicyExemption

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListForManagementGroupOperationResponse, error)
}

type ListForManagementGroupCompleteResult struct {
	Items []PolicyExemption
}

func (r ListForManagementGroupOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListForManagementGroupOperationResponse) LoadMore(ctx context.Context) (resp ListForManagementGroupOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListForManagementGroupOperationOptions struct {
	Filter *string
}

func DefaultListForManagementGroupOperationOptions() ListForManagementGroupOperationOptions {
	return ListForManagementGroupOperationOptions{}
}

func (o ListForManagementGroupOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListForManagementGroupOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListForManagementGroup ...
func (c PolicyExemptionsClient) ListForManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options ListForManagementGroupOperationOptions) (resp ListForManagementGroupOperationResponse, err error) {
	req, err := c.preparerForListForManagementGroup(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForManagementGroup", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForManagementGroup", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListForManagementGroup(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForManagementGroup", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListForManagementGroup prepares the ListForManagementGroup request.
func (c PolicyExemptionsClient) preparerForListForManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options ListForManagementGroupOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Authorization/policyExemptions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListForManagementGroupWithNextLink prepares the ListForManagementGroup request with the given nextLink token.
func (c PolicyExemptionsClient) preparerForListForManagementGroupWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListForManagementGroup handles the response to the ListForManagementGroup request. The method always
// closes the http.Response Body.
func (c PolicyExemptionsClient) responderForListForManagementGroup(resp *http.Response) (result ListForManagementGroupOperationResponse, err error) {
	type page struct {
		Values   []PolicyExemption `json:"value"`
		NextLink *string           `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListForManagementGroupOperationResponse, err error) {
			req, err := c.preparerForListForManagementGroupWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForManagementGroup", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForManagementGroup", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListForManagementGroup(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "policyexemptions.PolicyExemptionsClient", "ListForManagementGroup", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListForManagementGroupComplete retrieves all of the results into a single object
func (c PolicyExemptionsClient) ListForManagementGroupComplete(ctx context.Context, id commonids.ManagementGroupId, options ListForManagementGroupOperationOptions) (ListForManagementGroupCompleteResult, error) {
	return c.ListForManagementGroupCompleteMatchingPredicate(ctx, id, options, PolicyExemptionOperationPredicate{})
}

// ListForManagementGroupCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PolicyExemptionsClient) ListForManagementGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ManagementGroupId, options ListForManagementGroupOperationOptions, predicate PolicyExemptionOperationPredicate) (resp ListForManagementGroupCompleteResult, err error) {
	items := make([]PolicyExemption, 0)

	page, err := c.ListForManagementGroup(ctx, id, options)
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

	out := ListForManagementGroupCompleteResult{
		Items: items,
	}
	return out, nil
}
