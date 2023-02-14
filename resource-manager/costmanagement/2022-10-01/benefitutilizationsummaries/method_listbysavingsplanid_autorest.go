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

type ListBySavingsPlanIdOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BenefitUtilizationSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListBySavingsPlanIdOperationResponse, error)
}

type ListBySavingsPlanIdCompleteResult struct {
	Items []BenefitUtilizationSummary
}

func (r ListBySavingsPlanIdOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListBySavingsPlanIdOperationResponse) LoadMore(ctx context.Context) (resp ListBySavingsPlanIdOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListBySavingsPlanIdOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListBySavingsPlanIdOperationOptions() ListBySavingsPlanIdOperationOptions {
	return ListBySavingsPlanIdOperationOptions{}
}

func (o ListBySavingsPlanIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListBySavingsPlanIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.GrainParameter != nil {
		out["grainParameter"] = *o.GrainParameter
	}

	return out
}

// ListBySavingsPlanId ...
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanId(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions) (resp ListBySavingsPlanIdOperationResponse, err error) {
	req, err := c.preparerForListBySavingsPlanId(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanId", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanId", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListBySavingsPlanId(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanId", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListBySavingsPlanId prepares the ListBySavingsPlanId request.
func (c BenefitUtilizationSummariesClient) preparerForListBySavingsPlanId(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions) (*http.Request, error) {
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

// preparerForListBySavingsPlanIdWithNextLink prepares the ListBySavingsPlanId request with the given nextLink token.
func (c BenefitUtilizationSummariesClient) preparerForListBySavingsPlanIdWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListBySavingsPlanId handles the response to the ListBySavingsPlanId request. The method always
// closes the http.Response Body.
func (c BenefitUtilizationSummariesClient) responderForListBySavingsPlanId(resp *http.Response) (result ListBySavingsPlanIdOperationResponse, err error) {
	type page struct {
		Values   []BenefitUtilizationSummary `json:"value"`
		NextLink *string                     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListBySavingsPlanIdOperationResponse, err error) {
			req, err := c.preparerForListBySavingsPlanIdWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanId", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanId", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListBySavingsPlanId(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListBySavingsPlanId", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListBySavingsPlanIdComplete retrieves all of the results into a single object
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanIdComplete(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions) (ListBySavingsPlanIdCompleteResult, error) {
	return c.ListBySavingsPlanIdCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListBySavingsPlanIdCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BenefitUtilizationSummariesClient) ListBySavingsPlanIdCompleteMatchingPredicate(ctx context.Context, id SavingsPlanId, options ListBySavingsPlanIdOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (resp ListBySavingsPlanIdCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	page, err := c.ListBySavingsPlanId(ctx, id, options)
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

	out := ListBySavingsPlanIdCompleteResult{
		Items: items,
	}
	return out, nil
}
