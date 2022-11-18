package managedvmsizes

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedUnsupportedVMSizesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedVMSize
}

// ManagedUnsupportedVMSizesGet ...
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesGet(ctx context.Context, id ManagedUnsupportedVMSizeId) (result ManagedUnsupportedVMSizesGetOperationResponse, err error) {
	req, err := c.preparerForManagedUnsupportedVMSizesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForManagedUnsupportedVMSizesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedvmsizes.ManagedVMSizesClient", "ManagedUnsupportedVMSizesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForManagedUnsupportedVMSizesGet prepares the ManagedUnsupportedVMSizesGet request.
func (c ManagedVMSizesClient) preparerForManagedUnsupportedVMSizesGet(ctx context.Context, id ManagedUnsupportedVMSizeId) (*http.Request, error) {
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

// responderForManagedUnsupportedVMSizesGet handles the response to the ManagedUnsupportedVMSizesGet request. The method always
// closes the http.Response Body.
func (c ManagedVMSizesClient) responderForManagedUnsupportedVMSizesGet(resp *http.Response) (result ManagedUnsupportedVMSizesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
