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

type GetPrivateAccessSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateAccess
}

// GetPrivateAccessSlot ...
func (c WebAppsClient) GetPrivateAccessSlot(ctx context.Context, id SlotId) (result GetPrivateAccessSlotOperationResponse, err error) {
	req, err := c.preparerForGetPrivateAccessSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateAccessSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateAccessSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateAccessSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateAccessSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateAccessSlot prepares the GetPrivateAccessSlot request.
func (c WebAppsClient) preparerForGetPrivateAccessSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateAccess/virtualNetworks", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetPrivateAccessSlot handles the response to the GetPrivateAccessSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPrivateAccessSlot(resp *http.Response) (result GetPrivateAccessSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
