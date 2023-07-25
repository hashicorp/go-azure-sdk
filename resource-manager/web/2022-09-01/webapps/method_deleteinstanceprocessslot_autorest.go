package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteInstanceProcessSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteInstanceProcessSlot ...
func (c WebAppsClient) DeleteInstanceProcessSlot(ctx context.Context, id SlotInstanceProcessId) (result DeleteInstanceProcessSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteInstanceProcessSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceProcessSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceProcessSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteInstanceProcessSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteInstanceProcessSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteInstanceProcessSlot prepares the DeleteInstanceProcessSlot request.
func (c WebAppsClient) preparerForDeleteInstanceProcessSlot(ctx context.Context, id SlotInstanceProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteInstanceProcessSlot handles the response to the DeleteInstanceProcessSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteInstanceProcessSlot(resp *http.Response) (result DeleteInstanceProcessSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
