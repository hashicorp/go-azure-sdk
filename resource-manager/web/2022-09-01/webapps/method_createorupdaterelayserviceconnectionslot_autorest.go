package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateRelayServiceConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *RelayServiceConnectionEntity
}

// CreateOrUpdateRelayServiceConnectionSlot ...
func (c WebAppsClient) CreateOrUpdateRelayServiceConnectionSlot(ctx context.Context, id SlotHybridConnectionId, input RelayServiceConnectionEntity) (result CreateOrUpdateRelayServiceConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateRelayServiceConnectionSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateRelayServiceConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateRelayServiceConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateRelayServiceConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateRelayServiceConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateRelayServiceConnectionSlot prepares the CreateOrUpdateRelayServiceConnectionSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateRelayServiceConnectionSlot(ctx context.Context, id SlotHybridConnectionId, input RelayServiceConnectionEntity) (*http.Request, error) {
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

// responderForCreateOrUpdateRelayServiceConnectionSlot handles the response to the CreateOrUpdateRelayServiceConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateRelayServiceConnectionSlot(resp *http.Response) (result CreateOrUpdateRelayServiceConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
