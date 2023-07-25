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

type UpdateSlotConfigurationNamesOperationResponse struct {
	HttpResponse *http.Response
	Model        *SlotConfigNamesResource
}

// UpdateSlotConfigurationNames ...
func (c WebAppsClient) UpdateSlotConfigurationNames(ctx context.Context, id SiteId, input SlotConfigNamesResource) (result UpdateSlotConfigurationNamesOperationResponse, err error) {
	req, err := c.preparerForUpdateSlotConfigurationNames(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSlotConfigurationNames", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSlotConfigurationNames", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSlotConfigurationNames(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSlotConfigurationNames", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSlotConfigurationNames prepares the UpdateSlotConfigurationNames request.
func (c WebAppsClient) preparerForUpdateSlotConfigurationNames(ctx context.Context, id SiteId, input SlotConfigNamesResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/slotConfigNames", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSlotConfigurationNames handles the response to the UpdateSlotConfigurationNames request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateSlotConfigurationNames(resp *http.Response) (result UpdateSlotConfigurationNamesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
