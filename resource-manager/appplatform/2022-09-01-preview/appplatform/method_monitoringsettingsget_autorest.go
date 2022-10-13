package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringSettingsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *MonitoringSettingResource
}

// MonitoringSettingsGet ...
func (c AppPlatformClient) MonitoringSettingsGet(ctx context.Context, id SpringId) (result MonitoringSettingsGetOperationResponse, err error) {
	req, err := c.preparerForMonitoringSettingsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMonitoringSettingsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "MonitoringSettingsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMonitoringSettingsGet prepares the MonitoringSettingsGet request.
func (c AppPlatformClient) preparerForMonitoringSettingsGet(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/monitoringSettings/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForMonitoringSettingsGet handles the response to the MonitoringSettingsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForMonitoringSettingsGet(resp *http.Response) (result MonitoringSettingsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
