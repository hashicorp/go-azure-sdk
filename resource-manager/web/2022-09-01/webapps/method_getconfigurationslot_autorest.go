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

type GetConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// GetConfigurationSlot ...
func (c WebAppsClient) GetConfigurationSlot(ctx context.Context, id SlotId) (result GetConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForGetConfigurationSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetConfigurationSlot prepares the GetConfigurationSlot request.
func (c WebAppsClient) preparerForGetConfigurationSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetConfigurationSlot handles the response to the GetConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetConfigurationSlot(resp *http.Response) (result GetConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
