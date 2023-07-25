package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHostNameBindingSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *HostNameBinding
}

// GetHostNameBindingSlot ...
func (c WebAppsClient) GetHostNameBindingSlot(ctx context.Context, id SlotHostNameBindingId) (result GetHostNameBindingSlotOperationResponse, err error) {
	req, err := c.preparerForGetHostNameBindingSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHostNameBindingSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHostNameBindingSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHostNameBindingSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetHostNameBindingSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHostNameBindingSlot prepares the GetHostNameBindingSlot request.
func (c WebAppsClient) preparerForGetHostNameBindingSlot(ctx context.Context, id SlotHostNameBindingId) (*http.Request, error) {
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

// responderForGetHostNameBindingSlot handles the response to the GetHostNameBindingSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetHostNameBindingSlot(resp *http.Response) (result GetHostNameBindingSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
