package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetProcessModuleSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProcessModuleInfo
}

// GetProcessModuleSlot ...
func (c WebAppsClient) GetProcessModuleSlot(ctx context.Context, id ProcessModuleId) (result GetProcessModuleSlotOperationResponse, err error) {
	req, err := c.preparerForGetProcessModuleSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessModuleSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessModuleSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetProcessModuleSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessModuleSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetProcessModuleSlot prepares the GetProcessModuleSlot request.
func (c WebAppsClient) preparerForGetProcessModuleSlot(ctx context.Context, id ProcessModuleId) (*http.Request, error) {
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

// responderForGetProcessModuleSlot handles the response to the GetProcessModuleSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetProcessModuleSlot(resp *http.Response) (result GetProcessModuleSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
