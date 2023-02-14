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

type ListByBillingAccountIdOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]BenefitUtilizationSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByBillingAccountIdOperationResponse, error)
}

type ListByBillingAccountIdCompleteResult struct {
	Items []BenefitUtilizationSummary
}

func (r ListByBillingAccountIdOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByBillingAccountIdOperationResponse) LoadMore(ctx context.Context) (resp ListByBillingAccountIdOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByBillingAccountIdOperationOptions struct {
	Filter         *string
	GrainParameter *GrainParameter
}

func DefaultListByBillingAccountIdOperationOptions() ListByBillingAccountIdOperationOptions {
	return ListByBillingAccountIdOperationOptions{}
}

func (o ListByBillingAccountIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByBillingAccountIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["filter"] = *o.Filter
	}

	if o.GrainParameter != nil {
		out["grainParameter"] = *o.GrainParameter
	}

	return out
}

// ListByBillingAccountId ...
func (c BenefitUtilizationSummariesClient) ListByBillingAccountId(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions) (resp ListByBillingAccountIdOperationResponse, err error) {
	req, err := c.preparerForListByBillingAccountId(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListByBillingAccountId", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListByBillingAccountId", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByBillingAccountId(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListByBillingAccountId", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByBillingAccountId prepares the ListByBillingAccountId request.
func (c BenefitUtilizationSummariesClient) preparerForListByBillingAccountId(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions) (*http.Request, error) {
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

// preparerForListByBillingAccountIdWithNextLink prepares the ListByBillingAccountId request with the given nextLink token.
func (c BenefitUtilizationSummariesClient) preparerForListByBillingAccountIdWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByBillingAccountId handles the response to the ListByBillingAccountId request. The method always
// closes the http.Response Body.
func (c BenefitUtilizationSummariesClient) responderForListByBillingAccountId(resp *http.Response) (result ListByBillingAccountIdOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByBillingAccountIdOperationResponse, err error) {
			req, err := c.preparerForListByBillingAccountIdWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListByBillingAccountId", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListByBillingAccountId", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByBillingAccountId(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "benefitutilizationsummaries.BenefitUtilizationSummariesClient", "ListByBillingAccountId", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByBillingAccountIdComplete retrieves all of the results into a single object
func (c BenefitUtilizationSummariesClient) ListByBillingAccountIdComplete(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions) (ListByBillingAccountIdCompleteResult, error) {
	return c.ListByBillingAccountIdCompleteMatchingPredicate(ctx, id, options, BenefitUtilizationSummaryOperationPredicate{})
}

// ListByBillingAccountIdCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c BenefitUtilizationSummariesClient) ListByBillingAccountIdCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ListByBillingAccountIdOperationOptions, predicate BenefitUtilizationSummaryOperationPredicate) (resp ListByBillingAccountIdCompleteResult, err error) {
	items := make([]BenefitUtilizationSummary, 0)

	page, err := c.ListByBillingAccountId(ctx, id, options)
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

	out := ListByBillingAccountIdCompleteResult{
		Items: items,
	}
	return out, nil
}
