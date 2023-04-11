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

type AssessmentsMetadataSubscriptionListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecurityAssessmentMetadata

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (AssessmentsMetadataSubscriptionListOperationResponse, error)
}

type AssessmentsMetadataSubscriptionListCompleteResult struct {
	Items []SecurityAssessmentMetadata
}

func (r AssessmentsMetadataSubscriptionListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r AssessmentsMetadataSubscriptionListOperationResponse) LoadMore(ctx context.Context) (resp AssessmentsMetadataSubscriptionListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// AssessmentsMetadataSubscriptionList ...
func (c AssessmentsMetadataClient) AssessmentsMetadataSubscriptionList(ctx context.Context, id commonids.SubscriptionId) (resp AssessmentsMetadataSubscriptionListOperationResponse, err error) {
	req, err := c.preparerForAssessmentsMetadataSubscriptionList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForAssessmentsMetadataSubscriptionList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForAssessmentsMetadataSubscriptionList prepares the AssessmentsMetadataSubscriptionList request.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataSubscriptionList(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
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

// preparerForAssessmentsMetadataSubscriptionListWithNextLink prepares the AssessmentsMetadataSubscriptionList request with the given nextLink token.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataSubscriptionListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForAssessmentsMetadataSubscriptionList handles the response to the AssessmentsMetadataSubscriptionList request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForAssessmentsMetadataSubscriptionList(resp *http.Response) (result AssessmentsMetadataSubscriptionListOperationResponse, err error) {
	type page struct {
		Values   []SecurityAssessmentMetadata `json:"value"`
		NextLink *string                      `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result AssessmentsMetadataSubscriptionListOperationResponse, err error) {
			req, err := c.preparerForAssessmentsMetadataSubscriptionListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForAssessmentsMetadataSubscriptionList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// AssessmentsMetadataSubscriptionListComplete retrieves all of the results into a single object
func (c AssessmentsMetadataClient) AssessmentsMetadataSubscriptionListComplete(ctx context.Context, id commonids.SubscriptionId) (AssessmentsMetadataSubscriptionListCompleteResult, error) {
	return c.AssessmentsMetadataSubscriptionListCompleteMatchingPredicate(ctx, id, SecurityAssessmentMetadataOperationPredicate{})
}

// AssessmentsMetadataSubscriptionListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AssessmentsMetadataClient) AssessmentsMetadataSubscriptionListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecurityAssessmentMetadataOperationPredicate) (resp AssessmentsMetadataSubscriptionListCompleteResult, err error) {
	items := make([]SecurityAssessmentMetadata, 0)

	page, err := c.AssessmentsMetadataSubscriptionList(ctx, id)
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

	out := AssessmentsMetadataSubscriptionListCompleteResult{
		Items: items,
	}
	return out, nil
}
