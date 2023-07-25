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

type UpdateConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// UpdateConfigurationSlot ...
func (c WebAppsClient) UpdateConfigurationSlot(ctx context.Context, id SlotId, input SiteConfigResource) (result UpdateConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateConfigurationSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateConfigurationSlot prepares the UpdateConfigurationSlot request.
func (c WebAppsClient) preparerForUpdateConfigurationSlot(ctx context.Context, id SlotId, input SiteConfigResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateConfigurationSlot handles the response to the UpdateConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateConfigurationSlot(resp *http.Response) (result UpdateConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
