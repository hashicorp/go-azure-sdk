package assessmentsmetadata

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

type AssessmentsMetadataListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecurityAssessmentMetadataResponse

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AssessmentsMetadataListBySubscriptionOperationResponse, error)
}

type AssessmentsMetadataListBySubscriptionCompleteResult struct {
	Items []SecurityAssessmentMetadataResponse
}

func (r AssessmentsMetadataListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AssessmentsMetadataListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp AssessmentsMetadataListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// AssessmentsMetadataListBySubscription ...
func (c AssessmentsMetadataClient) AssessmentsMetadataListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp AssessmentsMetadataListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForAssessmentsMetadataListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAssessmentsMetadataListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAssessmentsMetadataListBySubscription prepares the AssessmentsMetadataListBySubscription request.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/assessmentMetadata", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForAssessmentsMetadataListBySubscriptionWithNextLink prepares the AssessmentsMetadataListBySubscription request with the given nextLink token.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAssessmentsMetadataListBySubscription handles the response to the AssessmentsMetadataListBySubscription request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForAssessmentsMetadataListBySubscription(resp *http.Response) (result AssessmentsMetadataListBySubscriptionOperationResponse, err error) {
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AssessmentsMetadataListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForAssessmentsMetadataListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAssessmentsMetadataListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AssessmentsMetadataListBySubscriptionComplete retrieves all of the results into a single object
func (c AssessmentsMetadataClient) AssessmentsMetadataListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (AssessmentsMetadataListBySubscriptionCompleteResult, error) {
	return c.AssessmentsMetadataListBySubscriptionCompleteMatchingPredicate(ctx, id, SecurityAssessmentMetadataResponseOperationPredicate{})
}

// AssessmentsMetadataListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AssessmentsMetadataClient) AssessmentsMetadataListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecurityAssessmentMetadataResponseOperationPredicate) (resp AssessmentsMetadataListBySubscriptionCompleteResult, err error) {
	items := make([]SecurityAssessmentMetadataResponse, 0)

	page, err := c.AssessmentsMetadataListBySubscription(ctx, id)
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

	out := AssessmentsMetadataListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
