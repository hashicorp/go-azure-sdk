package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisableRecommendationForSubscriptionOperationResponse struct {
	HttpResponse *http.Response
}

// DisableRecommendationForSubscription ...
func (c RecommendationsClient) DisableRecommendationForSubscription(ctx context.Context, id RecommendationId) (result DisableRecommendationForSubscriptionOperationResponse, err error) {
	req, err := c.preparerForDisableRecommendationForSubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForSubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForSubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDisableRecommendationForSubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableRecommendationForSubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDisableRecommendationForSubscription prepares the DisableRecommendationForSubscription request.
func (c RecommendationsClient) preparerForDisableRecommendationForSubscription(ctx context.Context, id RecommendationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/disable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDisableRecommendationForSubscription handles the response to the DisableRecommendationForSubscription request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForDisableRecommendationForSubscription(resp *http.Response) (result DisableRecommendationForSubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
