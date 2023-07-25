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

type GetContainerLogsZipSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetContainerLogsZipSlot ...
func (c WebAppsClient) GetContainerLogsZipSlot(ctx context.Context, id SlotId) (result GetContainerLogsZipSlotOperationResponse, err error) {
	req, err := c.preparerForGetContainerLogsZipSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContainerLogsZipSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContainerLogsZipSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetContainerLogsZipSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContainerLogsZipSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetContainerLogsZipSlot prepares the GetContainerLogsZipSlot request.
func (c WebAppsClient) preparerForGetContainerLogsZipSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/containerlogs/zip/download", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetContainerLogsZipSlot handles the response to the GetContainerLogsZipSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetContainerLogsZipSlot(resp *http.Response) (result GetContainerLogsZipSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
