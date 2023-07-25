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

type ApplySlotConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
}

// ApplySlotConfigurationSlot ...
func (c WebAppsClient) ApplySlotConfigurationSlot(ctx context.Context, id SlotId, input CsmSlotEntity) (result ApplySlotConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForApplySlotConfigurationSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApplySlotConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApplySlotConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForApplySlotConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ApplySlotConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForApplySlotConfigurationSlot prepares the ApplySlotConfigurationSlot request.
func (c WebAppsClient) preparerForApplySlotConfigurationSlot(ctx context.Context, id SlotId, input CsmSlotEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/applySlotConfig", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForApplySlotConfigurationSlot handles the response to the ApplySlotConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForApplySlotConfigurationSlot(resp *http.Response) (result ApplySlotConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
