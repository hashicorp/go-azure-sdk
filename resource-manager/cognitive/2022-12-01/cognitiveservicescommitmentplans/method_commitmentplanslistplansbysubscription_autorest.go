package cognitiveservicescommitmentplans

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

type CommitmentPlansListPlansBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CommitmentPlan

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (CommitmentPlansListPlansBySubscriptionOperationResponse, error)
}

type CommitmentPlansListPlansBySubscriptionCompleteResult struct {
	Items []CommitmentPlan
}

func (r CommitmentPlansListPlansBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r CommitmentPlansListPlansBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp CommitmentPlansListPlansBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// CommitmentPlansListPlansBySubscription ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp CommitmentPlansListPlansBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForCommitmentPlansListPlansBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansListPlansBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansListPlansBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForCommitmentPlansListPlansBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansListPlansBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForCommitmentPlansListPlansBySubscription prepares the CommitmentPlansListPlansBySubscription request.
func (c CognitiveServicesCommitmentPlansClient) preparerForCommitmentPlansListPlansBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CognitiveServices/commitmentPlans", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForCommitmentPlansListPlansBySubscriptionWithNextLink prepares the CommitmentPlansListPlansBySubscription request with the given nextLink token.
func (c CognitiveServicesCommitmentPlansClient) preparerForCommitmentPlansListPlansBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForCommitmentPlansListPlansBySubscription handles the response to the CommitmentPlansListPlansBySubscription request. The method always
// closes the http.Response Body.
func (c CognitiveServicesCommitmentPlansClient) responderForCommitmentPlansListPlansBySubscription(resp *http.Response) (result CommitmentPlansListPlansBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []CommitmentPlan `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result CommitmentPlansListPlansBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForCommitmentPlansListPlansBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansListPlansBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansListPlansBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForCommitmentPlansListPlansBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansListPlansBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// CommitmentPlansListPlansBySubscriptionComplete retrieves all of the results into a single object
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (CommitmentPlansListPlansBySubscriptionCompleteResult, error) {
	return c.CommitmentPlansListPlansBySubscriptionCompleteMatchingPredicate(ctx, id, CommitmentPlanOperationPredicate{})
}

// CommitmentPlansListPlansBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate CommitmentPlanOperationPredicate) (resp CommitmentPlansListPlansBySubscriptionCompleteResult, err error) {
	items := make([]CommitmentPlan, 0)

	page, err := c.CommitmentPlansListPlansBySubscription(ctx, id)
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

	out := CommitmentPlansListPlansBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
