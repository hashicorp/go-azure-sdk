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

type IoTSecuritySolutionsAnalyticsGetAllOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecuritySolutionAnalyticsModelList
}

// IoTSecuritySolutionsAnalyticsGetAll ...
func (c IoTSecuritySolutionsAnalyticsClient) IoTSecuritySolutionsAnalyticsGetAll(ctx context.Context, id IotSecuritySolutionId) (result IoTSecuritySolutionsAnalyticsGetAllOperationResponse, err error) {
	req, err := c.preparerForIoTSecuritySolutionsAnalyticsGetAll(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsGetAll", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsGetAll", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIoTSecuritySolutionsAnalyticsGetAll(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient", "IoTSecuritySolutionsAnalyticsGetAll", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIoTSecuritySolutionsAnalyticsGetAll prepares the IoTSecuritySolutionsAnalyticsGetAll request.
func (c IoTSecuritySolutionsAnalyticsClient) preparerForIoTSecuritySolutionsAnalyticsGetAll(ctx context.Context, id IotSecuritySolutionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/analyticsModels", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIoTSecuritySolutionsAnalyticsGetAll handles the response to the IoTSecuritySolutionsAnalyticsGetAll request. The method always
// closes the http.Response Body.
func (c IoTSecuritySolutionsAnalyticsClient) responderForIoTSecuritySolutionsAnalyticsGetAll(resp *http.Response) (result IoTSecuritySolutionsAnalyticsGetAllOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
