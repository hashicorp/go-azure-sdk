package iotsecuritysolutionsanalytics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecuritySolutionsAnalyticsAggregatedAlertDismissOperationResponse struct {
	HttpResponse *http.Response
}

// IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss(ctx context.Context, id AggregatedAlertId) (result IoTSecuritySolutionsAnalyticsAggregatedAlertDismissOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertDismiss(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIoTSecuritySolutionsAnalyticsAggregatedAlertDismiss(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertDismiss prepares the IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsAggregatedAlertDismiss(ctx context.Context, id AggregatedAlertId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dismiss", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIoTSecuritySolutionsAnalyticsAggregatedAlertDismiss handles the response to the IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsAggregatedAlertDismiss(resp *http.Response) (result IoTSecuritySolutionsAnalyticsAggregatedAlertDismissOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
