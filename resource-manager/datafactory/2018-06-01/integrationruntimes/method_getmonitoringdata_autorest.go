package integrationruntimes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMonitoringDataOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeMonitoringData
}

// GetMonitoringData ...
func (c IntegrationRuntimesClient) GetMonitoringData(ctx context.Context, id IntegrationRuntimeId) (result GetMonitoringDataOperationResponse, err error) {
	req, err := c.preparerForGetMonitoringData(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "GetMonitoringData", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "GetMonitoringData", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMonitoringData(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "GetMonitoringData", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMonitoringData prepares the GetMonitoringData request.
func (c IntegrationRuntimesClient) preparerForGetMonitoringData(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/monitoringData", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetMonitoringData handles the response to the GetMonitoringData request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForGetMonitoringData(resp *http.Response) (result GetMonitoringDataOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
