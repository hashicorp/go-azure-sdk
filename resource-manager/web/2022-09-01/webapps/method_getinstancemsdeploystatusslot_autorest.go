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

type GetInstanceMsDeployStatusSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployStatus
}

// GetInstanceMsDeployStatusSlot ...
func (c WebAppsClient) GetInstanceMsDeployStatusSlot(ctx context.Context, id SlotInstanceId) (result GetInstanceMsDeployStatusSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceMsDeployStatusSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMsDeployStatusSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMsDeployStatusSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceMsDeployStatusSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceMsDeployStatusSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceMsDeployStatusSlot prepares the GetInstanceMsDeployStatusSlot request.
func (c WebAppsClient) preparerForGetInstanceMsDeployStatusSlot(ctx context.Context, id SlotInstanceId) (*http.Request, error) {
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

// responderForGetInstanceMsDeployStatusSlot handles the response to the GetInstanceMsDeployStatusSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceMsDeployStatusSlot(resp *http.Response) (result GetInstanceMsDeployStatusSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
