package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateHybridConnectionSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnection
}

// CreateOrUpdateHybridConnectionSlot ...
func (c WebAppsClient) CreateOrUpdateHybridConnectionSlot(ctx context.Context, id SlotHybridConnectionNamespaceRelayId, input HybridConnection) (result CreateOrUpdateHybridConnectionSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateHybridConnectionSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHybridConnectionSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHybridConnectionSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateHybridConnectionSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateHybridConnectionSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateHybridConnectionSlot prepares the CreateOrUpdateHybridConnectionSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateHybridConnectionSlot(ctx context.Context, id SlotHybridConnectionNamespaceRelayId, input HybridConnection) (*http.Request, error) {
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

// responderForCreateOrUpdateHybridConnectionSlot handles the response to the CreateOrUpdateHybridConnectionSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateHybridConnectionSlot(resp *http.Response) (result CreateOrUpdateHybridConnectionSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
