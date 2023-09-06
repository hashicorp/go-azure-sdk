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

type AggregatedAlertDismissOperationResponse struct {
	HttpResponse *http.Response
}

// AggregatedAlertDismiss ...
func (c IoTSecuritySolutionsAnalyticsClient) AggregatedAlertDismiss(ctx context.Context, id AggregatedAlertId) (result AggregatedAlertDismissOperationResponse, err error) {
	req, err := c.preparerForAggregatedAlertDismiss(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "AggregatedAlertDismiss", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "AggregatedAlertDismiss", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAggregatedAlertDismiss(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "AggregatedAlertDismiss", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAggregatedAlertDismiss prepares the AggregatedAlertDismiss request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForAggregatedAlertDismiss(ctx context.Context, id AggregatedAlertId) (*http.Request, error) {
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

// responderForAggregatedAlertDismiss handles the response to the AggregatedAlertDismiss request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForAggregatedAlertDismiss(resp *http.Response) (result AggregatedAlertDismissOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
