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

type GetWebSiteContainerLogsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetWebSiteContainerLogsSlot ...
func (c WebAppsClient) GetWebSiteContainerLogsSlot(ctx context.Context, id SlotId) (result GetWebSiteContainerLogsSlotOperationResponse, err error) {
	req, err := c.preparerForGetWebSiteContainerLogsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebSiteContainerLogsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebSiteContainerLogsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetWebSiteContainerLogsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebSiteContainerLogsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetWebSiteContainerLogsSlot prepares the GetWebSiteContainerLogsSlot request.
func (c WebAppsClient) preparerForGetWebSiteContainerLogsSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/containerlogs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetWebSiteContainerLogsSlot handles the response to the GetWebSiteContainerLogsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetWebSiteContainerLogsSlot(resp *http.Response) (result GetWebSiteContainerLogsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
