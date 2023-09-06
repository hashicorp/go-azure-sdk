package iotsecuritysolutionsanalytics

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregatedAlertGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecurityAggregatedAlert
}

// AggregatedAlertGet ...
func (c IoTSecuritySolutionsAnalyticsClient) AggregatedAlertGet(ctx context.Context, id AggregatedAlertId) (result AggregatedAlertGetOperationResponse, err error) {
	req, err := c.preparerForAggregatedAlertGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "AggregatedAlertGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "AggregatedAlertGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAggregatedAlertGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "AggregatedAlertGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAggregatedAlertGet prepares the AggregatedAlertGet request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForAggregatedAlertGet(ctx context.Context, id AggregatedAlertId) (*http.Request, error) {
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

// responderForAggregatedAlertGet handles the response to the AggregatedAlertGet request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForAggregatedAlertGet(resp *http.Response) (result AggregatedAlertGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
