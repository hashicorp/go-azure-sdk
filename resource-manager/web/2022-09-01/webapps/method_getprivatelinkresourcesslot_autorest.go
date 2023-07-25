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

type GetPrivateLinkResourcesSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResourcesWrapper
}

// GetPrivateLinkResourcesSlot ...
func (c WebAppsClient) GetPrivateLinkResourcesSlot(ctx context.Context, id SlotId) (result GetPrivateLinkResourcesSlotOperationResponse, err error) {
	req, err := c.preparerForGetPrivateLinkResourcesSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateLinkResourcesSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateLinkResourcesSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateLinkResourcesSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPrivateLinkResourcesSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateLinkResourcesSlot prepares the GetPrivateLinkResourcesSlot request.
func (c WebAppsClient) preparerForGetPrivateLinkResourcesSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateLinkResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetPrivateLinkResourcesSlot handles the response to the GetPrivateLinkResourcesSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPrivateLinkResourcesSlot(resp *http.Response) (result GetPrivateLinkResourcesSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
