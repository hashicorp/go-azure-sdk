package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceProcessSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessInfo
}

// GetInstanceProcessSlot ...
func (c WebAppsClient) GetInstanceProcessSlot(ctx context.Context, id SlotInstanceProcessId) (result GetInstanceProcessSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceProcessSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceProcessSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceProcessSlot prepares the GetInstanceProcessSlot request.
func (c WebAppsClient) preparerForGetInstanceProcessSlot(ctx context.Context, id SlotInstanceProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetInstanceProcessSlot handles the response to the GetInstanceProcessSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceProcessSlot(resp *http.Response) (result GetInstanceProcessSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
