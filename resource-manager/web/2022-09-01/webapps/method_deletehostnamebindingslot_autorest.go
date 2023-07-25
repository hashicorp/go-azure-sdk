package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteHostNameBindingSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteHostNameBindingSlot ...
func (c WebAppsClient) DeleteHostNameBindingSlot(ctx context.Context, id SlotHostNameBindingId) (result DeleteHostNameBindingSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteHostNameBindingSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostNameBindingSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostNameBindingSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteHostNameBindingSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostNameBindingSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteHostNameBindingSlot prepares the DeleteHostNameBindingSlot request.
func (c WebAppsClient) preparerForDeleteHostNameBindingSlot(ctx context.Context, id SlotHostNameBindingId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteHostNameBindingSlot handles the response to the DeleteHostNameBindingSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteHostNameBindingSlot(resp *http.Response) (result DeleteHostNameBindingSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
