package benefitutilizationsummaries

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

type ListBySavingsPlanOrderOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BenefitUtilizationSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBySavingsPlanOrderOperationResponse, error)
}

type ListBySavingsPlanOrderCompleteResult struct {
	Items []BenefitUtilizationSummary
}

func (r ListBySavingsPlanOrderOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBySavingsPlanOrderOperationResponse) LoadMore(ctx context.Context) (resp ListBySavingsPlanOrderOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListBySavingsPlanOrderOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListBySavingsPlanOrderOperationOptions() ListBySavingsPlanOrderOperationOptions {
	return ListBySavingsPlanOrderOperationOptions{}
}

func (o ListBySavingsPlanOrderOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListBySavingsPlanOrderOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.GrainParameter != nil {
		out["grainParameter"] = *o.GrainParameter
	}

	return out
}

// ListBySavingsPlanOrder ...
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanOrder(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions) (resp ListBySavingsPlanOrderOperationResponse, err error) {
	req, err := c.preparerForListBySavingsPlanOrder(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanOrder", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanOrder", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBySavingsPlanOrder(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanOrder", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBySavingsPlanOrder prepares the ListBySavingsPlanOrder request.
func (c BenefitUtilizationSummariesClient) preparerForListBySavingsPlanOrder(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CostManagement/benefitUtilizationSummaries", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListBySavingsPlanOrderWithNextLink prepares the ListBySavingsPlanOrder request with the given nextLink token.
func (c BenefitUtilizationSummariesClient) preparerForListBySavingsPlanOrderWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBySavingsPlanOrder handles the response to the ListBySavingsPlanOrder request. The method always
// closes the http.Response Body.
func (c BenefitUtilizationSummariesClient) responderForListBySavingsPlanOrder(resp *http.Response) (result ListBySavingsPlanOrderOperationResponse, err error) {
	type page struct {
		Values   []BenefitUtilizationSummary `json:"value"`
		NextLink *string                     `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBySavingsPlanOrderOperationResponse, err error) {
			req, err := c.preparerForListBySavingsPlanOrderWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanOrder", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanOrder", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBySavingsPlanOrder(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanOrder", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBySavingsPlanOrderComplete retrieves all of the results into a single object
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanOrderComplete(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions) (ListBySavingsPlanOrderCompleteResult, error) {
	return c.ListBySavingsPlanOrderCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListBySavingsPlanOrderCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanOrderCompleteMatchingPredicate(ctx context.Context, id SavingsPlanOrderId, options ListBySavingsPlanOrderOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (resp ListBySavingsPlanOrderCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	page, err := c.ListBySavingsPlanOrder(ctx, id, options)
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

	out := ListBySavingsPlanOrderCompleteResult{
		Items: items,
	}
	return out, nil
}
