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

type GetMSDeployLogOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployLog
}

// GetMSDeployLog ...
func (c WebAppsClient) GetMSDeployLog(ctx context.Context, id SiteId) (result GetMSDeployLogOperationResponse, err error) {
	req, err := c.preparerForGetMSDeployLog(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployLog", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployLog", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMSDeployLog(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployLog", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMSDeployLog prepares the GetMSDeployLog request.
func (c WebAppsClient) preparerForGetMSDeployLog(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForGetMSDeployLog handles the response to the GetMSDeployLog request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetMSDeployLog(resp *http.Response) (result GetMSDeployLogOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
