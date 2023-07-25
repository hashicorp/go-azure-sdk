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

type GetInstanceMsDeployStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployStatus
}

// GetInstanceMsDeployStatus ...
func (c WebAppsClient) GetInstanceMsDeployStatus(ctx context.Context, id InstanceId) (result GetInstanceMsDeployStatusOperationResponse, err error) {
	req, err := c.preparerForGetInstanceMsDeployStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMsDeployStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMsDeployStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceMsDeployStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMsDeployStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceMsDeployStatus prepares the GetInstanceMsDeployStatus request.
func (c WebAppsClient) preparerForGetInstanceMsDeployStatus(ctx context.Context, id InstanceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/extensions/mSDeploy", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetInstanceMsDeployStatus handles the response to the GetInstanceMsDeployStatus request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceMsDeployStatus(resp *http.Response) (result GetInstanceMsDeployStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
