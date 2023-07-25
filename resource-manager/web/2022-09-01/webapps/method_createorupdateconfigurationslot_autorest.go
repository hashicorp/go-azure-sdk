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

type CreateOrUpdateConfigurationSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteConfigResource
}

// CreateOrUpdateConfigurationSlot ...
func (c WebAppsClient) CreateOrUpdateConfigurationSlot(ctx context.Context, id SlotId, input SiteConfigResource) (result CreateOrUpdateConfigurationSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateConfigurationSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateConfigurationSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateConfigurationSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateConfigurationSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdateConfigurationSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateConfigurationSlot prepares the CreateOrUpdateConfigurationSlot request.
func (c WebAppsClient) preparerForCreateOrUpdateConfigurationSlot(ctx context.Context, id SlotId, input SiteConfigResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/web", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateConfigurationSlot handles the response to the CreateOrUpdateConfigurationSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdateConfigurationSlot(resp *http.Response) (result CreateOrUpdateConfigurationSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
