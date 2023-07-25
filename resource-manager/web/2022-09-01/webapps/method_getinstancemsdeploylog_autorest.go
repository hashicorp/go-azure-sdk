package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceMSDeployLogOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployLog
}

// GetInstanceMSDeployLog ...
func (c WebAppsClient) GetInstanceMSDeployLog(ctx context.Context, id InstanceId) (result GetInstanceMSDeployLogOperationResponse, err error) {
	req, err := c.preparerForGetInstanceMSDeployLog(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMSDeployLog", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMSDeployLog", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceMSDeployLog(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMSDeployLog", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceMSDeployLog prepares the GetInstanceMSDeployLog request.
func (c WebAppsClient) preparerForGetInstanceMSDeployLog(ctx context.Context, id InstanceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensions/mSDeploy/log", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetInstanceMSDeployLog handles the response to the GetInstanceMSDeployLog request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceMSDeployLog(resp *http.Response) (result GetInstanceMSDeployLogOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
