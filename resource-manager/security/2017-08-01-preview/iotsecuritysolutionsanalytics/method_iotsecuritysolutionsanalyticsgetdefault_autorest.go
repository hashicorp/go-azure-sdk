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

type IoTSecuritySolutionsAnalyticsGetDefaultOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecuritySolutionAnalyticsModel
}

// IoTSecuritySolutionsAnalyticsGetDefault ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsGetDefault(ctx context.Context, id IotSecuritySolutionId) (result IoTSecuritySolutionsAnalyticsGetDefaultOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsGetDefault(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsGetDefault", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsGetDefault", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIoTSecuritySolutionsAnalyticsGetDefault(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsGetDefault", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIoTSecuritySolutionsAnalyticsGetDefault prepares the IoTSecuritySolutionsAnalyticsGetDefault request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsGetDefault(ctx context.Context, id IotSecuritySolutionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/analyticsModels/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIoTSecuritySolutionsAnalyticsGetDefault handles the response to the IoTSecuritySolutionsAnalyticsGetDefault request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsGetDefault(resp *http.Response) (result IoTSecuritySolutionsAnalyticsGetDefaultOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
