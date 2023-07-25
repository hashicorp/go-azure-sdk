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

type UpdateMetadataSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// UpdateMetadataSlot ...
func (c WebAppsClient) UpdateMetadataSlot(ctx context.Context, id SlotId, input StringDictionary) (result UpdateMetadataSlotOperationResponse, err error) {
	req, err := c.preparerForUpdateMetadataSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateMetadataSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateMetadataSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateMetadataSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateMetadataSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateMetadataSlot prepares the UpdateMetadataSlot request.
func (c WebAppsClient) preparerForUpdateMetadataSlot(ctx context.Context, id SlotId, input StringDictionary) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/metadata", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateMetadataSlot handles the response to the UpdateMetadataSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateMetadataSlot(resp *http.Response) (result UpdateMetadataSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
