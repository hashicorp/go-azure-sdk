package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInstanceInfoSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *WebSiteInstanceStatus
}

// GetInstanceInfoSlot ...
func (c WebAppsClient) GetInstanceInfoSlot(ctx context.Context, id SlotInstanceId) (result GetInstanceInfoSlotOperationResponse, err error) {
	req, err := c.preparerForGetInstanceInfoSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceInfoSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceInfoSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetInstanceInfoSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetInstanceInfoSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetInstanceInfoSlot prepares the GetInstanceInfoSlot request.
func (c WebAppsClient) preparerForGetInstanceInfoSlot(ctx context.Context, id SlotInstanceId) (*http.Request, error) {
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

// responderForGetInstanceInfoSlot handles the response to the GetInstanceInfoSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetInstanceInfoSlot(resp *http.Response) (result GetInstanceInfoSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
