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

type GetMSDeployStatusSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *MSDeployStatus
}

// GetMSDeployStatusSlot ...
func (c WebAppsClient) GetMSDeployStatusSlot(ctx context.Context, id SlotId) (result GetMSDeployStatusSlotOperationResponse, err error) {
	req, err := c.preparerForGetMSDeployStatusSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployStatusSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployStatusSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMSDeployStatusSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMSDeployStatusSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMSDeployStatusSlot prepares the GetMSDeployStatusSlot request.
func (c WebAppsClient) preparerForGetMSDeployStatusSlot(ctx context.Context, id SlotId) (*http.Request, error) {
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

// responderForGetMSDeployStatusSlot handles the response to the GetMSDeployStatusSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetMSDeployStatusSlot(resp *http.Response) (result GetMSDeployStatusSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
