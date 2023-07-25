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

type GetInstanceMSDeployLogSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployLog
}

// GetInstanceMSDeployLogSlot ...
func (c WebAppsClient) GetInstanceMSDeployLogSlot(ctx context.Context, id SlotInstanceId) (result GetInstanceMSDeployLogSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceMSDeployLogSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMSDeployLogSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMSDeployLogSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceMSDeployLogSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMSDeployLogSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceMSDeployLogSlot prepares the GetInstanceMSDeployLogSlot request.
func (c WebAppsClient) preparerForGetInstanceMSDeployLogSlot(ctx context.Context, id SlotInstanceId) (*http.Request, error) {
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

// responderForGetInstanceMSDeployLogSlot handles the response to the GetInstanceMSDeployLogSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceMSDeployLogSlot(resp *http.Response) (result GetInstanceMSDeployLogSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
