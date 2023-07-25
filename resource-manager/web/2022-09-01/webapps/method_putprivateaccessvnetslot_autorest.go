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

type PutPrivateAccessVnetSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateAccess
}

// PutPrivateAccessVnetSlot ...
func (c WebAppsClient) PutPrivateAccessVnetSlot(ctx context.Context, id SlotId, input PrivateAccess) (result PutPrivateAccessVnetSlotOperationResponse, err error) {
	req, err := c.preparerForPutPrivateAccessVnetSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "PutPrivateAccessVnetSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "PutPrivateAccessVnetSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutPrivateAccessVnetSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "PutPrivateAccessVnetSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutPrivateAccessVnetSlot prepares the PutPrivateAccessVnetSlot request.
func (c WebAppsClient) preparerForPutPrivateAccessVnetSlot(ctx context.Context, id SlotId, input PrivateAccess) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateAccess/virtualNetworks", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPutPrivateAccessVnetSlot handles the response to the PutPrivateAccessVnetSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForPutPrivateAccessVnetSlot(resp *http.Response) (result PutPrivateAccessVnetSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
