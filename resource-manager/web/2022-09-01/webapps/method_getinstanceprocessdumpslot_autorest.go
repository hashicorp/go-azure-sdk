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

type GetInstanceProcessDumpSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetInstanceProcessDumpSlot ...
func (c WebAppsClient) GetInstanceProcessDumpSlot(ctx context.Context, id SlotInstanceProcessId) (result GetInstanceProcessDumpSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceProcessDumpSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessDumpSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessDumpSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceProcessDumpSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceProcessDumpSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceProcessDumpSlot prepares the GetInstanceProcessDumpSlot request.
func (c WebAppsClient) preparerForGetInstanceProcessDumpSlot(ctx context.Context, id SlotInstanceProcessId) (*http.Request, error) {
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

// responderForGetInstanceProcessDumpSlot handles the response to the GetInstanceProcessDumpSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceProcessDumpSlot(resp *http.Response) (result GetInstanceProcessDumpSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
