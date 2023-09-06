package iotsecuritysolutionsanalytics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecurityAggregatedRecommendation
}

// RecommendationGet ...
func (c IoTSecuritySolutionsAnalyticsClient) RecommendationGet(ctx context.Context, id AggregatedRecommendationId) (result RecommendationGetOperationResponse, err error) {
	req, err := c.preparerForRecommendationGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "RecommendationGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "RecommendationGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecommendationGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "RecommendationGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecommendationGet prepares the RecommendationGet request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForRecommendationGet(ctx context.Context, id AggregatedRecommendationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRecommendationGet handles the response to the RecommendationGet request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForRecommendationGet(resp *http.Response) (result RecommendationGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
