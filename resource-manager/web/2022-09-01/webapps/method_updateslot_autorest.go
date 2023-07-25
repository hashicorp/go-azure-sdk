package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *Site
}

// UpdateSlot ...
func (c WebAppsClient) UpdateSlot(ctx context.Context, id SlotId, input SitePatchResource) (result UpdateSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateSlot prepares the UpdateSlot request.
func (c WebAppsClient) preparerForUpdateSlot(ctx context.Context, id SlotId, input SitePatchResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateSlot handles the response to the UpdateSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateSlot(resp *http.Response) (result UpdateSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
