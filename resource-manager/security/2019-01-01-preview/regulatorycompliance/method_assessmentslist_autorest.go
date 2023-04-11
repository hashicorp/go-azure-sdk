package regulatorycompliance

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

type AssessmentsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RegulatoryComplianceAssessment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AssessmentsListOperationResponse, error)
}

type AssessmentsListCompleteResult struct {
	Items []RegulatoryComplianceAssessment
}

func (r AssessmentsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AssessmentsListOperationResponse) LoadMore(ctx context.Context) (resp AssessmentsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type AssessmentsListOperationOptions struct {
	Filter *string
}

func DefaultAssessmentsListOperationOptions() AssessmentsListOperationOptions {
	return AssessmentsListOperationOptions{}
}

func (o AssessmentsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AssessmentsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// AssessmentsList ...
func (c RegulatoryComplianceClient) AssessmentsList(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions) (resp AssessmentsListOperationResponse, err error) {
	req, err := c.preparerForAssessmentsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAssessmentsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAssessmentsList prepares the AssessmentsList request.
func (c RegulatoryComplianceClient) preparerForAssessmentsList(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/regulatoryComplianceAssessments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAssessmentsListWithNextLink prepares the AssessmentsList request with the given nextLink token.
func (c RegulatoryComplianceClient) preparerForAssessmentsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAssessmentsList handles the response to the AssessmentsList request. The method always
// closes the http.Response Body.
func (c RegulatoryComplianceClient) responderForAssessmentsList(resp *http.Response) (result AssessmentsListOperationResponse, err error) {
	type page struct {
		Values   []RegulatoryComplianceAssessment `json:"value"`
		NextLink *string                          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AssessmentsListOperationResponse, err error) {
			req, err := c.preparerForAssessmentsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAssessmentsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "regulatorycompliance.RegulatoryComplianceClient", "AssessmentsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AssessmentsListComplete retrieves all of the results into a single object
func (c RegulatoryComplianceClient) AssessmentsListComplete(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions) (AssessmentsListCompleteResult, error) {
	return c.AssessmentsListCompleteMatchingPredicate(ctx, id, options, RegulatoryComplianceAssessmentOperationPredicate{})
}

// AssessmentsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RegulatoryComplianceClient) AssessmentsListCompleteMatchingPredicate(ctx context.Context, id RegulatoryComplianceControlId, options AssessmentsListOperationOptions, predicate RegulatoryComplianceAssessmentOperationPredicate) (resp AssessmentsListCompleteResult, err error) {
	items := make([]RegulatoryComplianceAssessment, 0)

	page, err := c.AssessmentsList(ctx, id, options)
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

	out := AssessmentsListCompleteResult{
		Items: items,
	}
	return out, nil
}
