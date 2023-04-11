package iotsecuritysolutionsanalytics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionsAnalyticsAggregatedAlertGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecurityAggregatedAlert
}

// IoTSecuritySolutionsAnalyticsAggregatedAlertGet ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsAggregatedAlertGet(ctx context.Context, id AggregatedAlertId) (result IoTSecuritySolutionsAnalyticsAggregatedAlertGetOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIoTSecuritySolutionsAnalyticsAggregatedAlertGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertGet prepares the IoTSecuritySolutionsAnalyticsAggregatedAlertGet request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertGet(ctx context.Context, id AggregatedAlertId) (*http.Request, error) {
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

// responderForIoTSecuritySolutionsAnalyticsAggregatedAlertGet handles the response to the IoTSecuritySolutionsAnalyticsAggregatedAlertGet request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsAggregatedAlertGet(resp *http.Response) (result IoTSecuritySolutionsAnalyticsAggregatedAlertGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
