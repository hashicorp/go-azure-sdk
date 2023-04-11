package iotsecuritysolutionsanalytics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionsAnalyticsRecommendationGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecurityAggregatedRecommendation
}

// IoTSecuritySolutionsAnalyticsRecommendationGet ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsRecommendationGet(ctx context.Context, id AggregatedRecommendationId) (result IoTSecuritySolutionsAnalyticsRecommendationGetOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsRecommendationGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIoTSecuritySolutionsAnalyticsRecommendationGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsRecommendationGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIoTSecuritySolutionsAnalyticsRecommendationGet prepares the IoTSecuritySolutionsAnalyticsRecommendationGet request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsRecommendationGet(ctx context.Context, id AggregatedRecommendationId) (*http.Request, error) {
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

// responderForIoTSecuritySolutionsAnalyticsRecommendationGet handles the response to the IoTSecuritySolutionsAnalyticsRecommendationGet request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsRecommendationGet(resp *http.Response) (result IoTSecuritySolutionsAnalyticsRecommendationGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
