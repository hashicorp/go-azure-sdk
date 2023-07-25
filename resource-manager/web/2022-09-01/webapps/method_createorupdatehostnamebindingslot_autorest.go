package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateHostNameBindingSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *HostNameBinding
}

// CreateOrUpdateHostNameBindingSlot ...
func (c WebAppsClient) CreateOrUpdateHostNameBindingSlot(ctx context.Context, id SlotHostNameBindingId, input HostNameBinding) (result CreateOrUpdateHostNameBindingSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateHostNameBindingSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHostNameBindingSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHostNameBindingSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateHostNameBindingSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHostNameBindingSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateHostNameBindingSlot prepares the CreateOrUpdateHostNameBindingSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateHostNameBindingSlot(ctx context.Context, id SlotHostNameBindingId, input HostNameBinding) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateHostNameBindingSlot handles the response to the CreateOrUpdateHostNameBindingSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateHostNameBindingSlot(resp *http.Response) (result CreateOrUpdateHostNameBindingSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
