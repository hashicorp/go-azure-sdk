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

type GetSourceControlSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *SiteSourceControl
}

// GetSourceControlSlot ...
func (c WebAppsClient) GetSourceControlSlot(ctx context.Context, id SlotId) (result GetSourceControlSlotOperationResponse, err error) {
	req, err := c.preparerForGetSourceControlSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSourceControlSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSourceControlSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSourceControlSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetSourceControlSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSourceControlSlot prepares the GetSourceControlSlot request.
func (c WebAppsClient) preparerForGetSourceControlSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/sourceControls/web", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSourceControlSlot handles the response to the GetSourceControlSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetSourceControlSlot(resp *http.Response) (result GetSourceControlSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
