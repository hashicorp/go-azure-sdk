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

type ResetSlotConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
}

// ResetSlotConfigurationSlot ...
func (c WebAppsClient) ResetSlotConfigurationSlot(ctx context.Context, id SlotId) (result ResetSlotConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForResetSlotConfigurationSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ResetSlotConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ResetSlotConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForResetSlotConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ResetSlotConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForResetSlotConfigurationSlot prepares the ResetSlotConfigurationSlot request.
func (c WebAppsClient) preparerForResetSlotConfigurationSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resetSlotConfig", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResetSlotConfigurationSlot handles the response to the ResetSlotConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForResetSlotConfigurationSlot(resp *http.Response) (result ResetSlotConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
