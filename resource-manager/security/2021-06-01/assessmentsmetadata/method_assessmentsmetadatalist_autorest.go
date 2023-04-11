package assessmentsmetadata

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

type AssessmentsMetadataListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecurityAssessmentMetadataResponse

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AssessmentsMetadataListOperationResponse, error)
}

type AssessmentsMetadataListCompleteResult struct {
	Items []SecurityAssessmentMetadataResponse
}

func (r AssessmentsMetadataListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AssessmentsMetadataListOperationResponse) LoadMore(ctx context.Context) (resp AssessmentsMetadataListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// AssessmentsMetadataList ...
func (c AssessmentsMetadataClient) AssessmentsMetadataList(ctx context.Context) (resp AssessmentsMetadataListOperationResponse, err error) {
	req, err := c.preparerForAssessmentsMetadataList(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAssessmentsMetadataList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAssessmentsMetadataList prepares the AssessmentsMetadataList request.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataList(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Security/assessmentMetadata"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAssessmentsMetadataListWithNextLink prepares the AssessmentsMetadataList request with the given nextLink token.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAssessmentsMetadataList handles the response to the AssessmentsMetadataList request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForAssessmentsMetadataList(resp *http.Response) (result AssessmentsMetadataListOperationResponse, err error) {
	type page struct {
		Values   []SecurityAssessmentMetadataResponse `json:"value"`
		NextLink *string                              `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AssessmentsMetadataListOperationResponse, err error) {
			req, err := c.preparerForAssessmentsMetadataListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAssessmentsMetadataList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AssessmentsMetadataListComplete retrieves all of the results into a single object
func (c AssessmentsMetadataClient) AssessmentsMetadataListComplete(ctx context.Context) (AssessmentsMetadataListCompleteResult, error) {
	return c.AssessmentsMetadataListCompleteMatchingPredicate(ctx, SecurityAssessmentMetadataResponseOperationPredicate{})
}

// AssessmentsMetadataListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AssessmentsMetadataClient) AssessmentsMetadataListCompleteMatchingPredicate(ctx context.Context, predicate SecurityAssessmentMetadataResponseOperationPredicate) (resp AssessmentsMetadataListCompleteResult, err error) {
	items := make([]SecurityAssessmentMetadataResponse, 0)

	page, err := c.AssessmentsMetadataList(ctx)
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

	out := AssessmentsMetadataListCompleteResult{
		Items: items,
	}
	return out, nil
}
