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

type GetProcessDumpSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetProcessDumpSlot ...
func (c WebAppsClient) GetProcessDumpSlot(ctx context.Context, id SlotProcessId) (result GetProcessDumpSlotOperationResponse, err error) {
	req, err := c.preparerForGetProcessDumpSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessDumpSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessDumpSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetProcessDumpSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessDumpSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetProcessDumpSlot prepares the GetProcessDumpSlot request.
func (c WebAppsClient) preparerForGetProcessDumpSlot(ctx context.Context, id SlotProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dump", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetProcessDumpSlot handles the response to the GetProcessDumpSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetProcessDumpSlot(resp *http.Response) (result GetProcessDumpSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
