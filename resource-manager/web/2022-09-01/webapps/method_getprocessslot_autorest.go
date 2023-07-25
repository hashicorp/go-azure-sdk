package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetProcessSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessInfo
}

// GetProcessSlot ...
func (c WebAppsClient) GetProcessSlot(ctx context.Context, id SlotProcessId) (result GetProcessSlotOperationResponse, err error) {
	req, err := c.preparerForGetProcessSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetProcessSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetProcessSlot prepares the GetProcessSlot request.
func (c WebAppsClient) preparerForGetProcessSlot(ctx context.Context, id SlotProcessId) (*http.Request, error) {
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

// responderForGetProcessSlot handles the response to the GetProcessSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetProcessSlot(resp *http.Response) (result GetProcessSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
