package integrationruntime

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringDataListOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeMonitoringData
}

// MonitoringDataList ...
func (c IntegrationRuntimeClient) MonitoringDataList(ctx context.Context, id IntegrationRuntimeId) (result MonitoringDataListOperationResponse, err error) {
	req, err := c.preparerForMonitoringDataList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "MonitoringDataList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "MonitoringDataList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMonitoringDataList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "MonitoringDataList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMonitoringDataList prepares the MonitoringDataList request.
func (c IntegrationRuntimeClient) preparerForMonitoringDataList(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
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

// responderForMonitoringDataList handles the response to the MonitoringDataList request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForMonitoringDataList(resp *http.Response) (result MonitoringDataListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
