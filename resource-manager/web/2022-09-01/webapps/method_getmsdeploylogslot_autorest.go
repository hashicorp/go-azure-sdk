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

type GetMSDeployLogSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployLog
}

// GetMSDeployLogSlot ...
func (c WebAppsClient) GetMSDeployLogSlot(ctx context.Context, id SlotId) (result GetMSDeployLogSlotOperationResponse, err error) {
	req, err := c.preparerForGetMSDeployLogSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployLogSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployLogSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMSDeployLogSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployLogSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMSDeployLogSlot prepares the GetMSDeployLogSlot request.
func (c WebAppsClient) preparerForGetMSDeployLogSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// responderForGetMSDeployLogSlot handles the response to the GetMSDeployLogSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetMSDeployLogSlot(resp *http.Response) (result GetMSDeployLogSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
