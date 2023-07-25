package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteProcessSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteProcessSlot ...
func (c WebAppsClient) DeleteProcessSlot(ctx context.Context, id SlotProcessId) (result DeleteProcessSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteProcessSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteProcessSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteProcessSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteProcessSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteProcessSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteProcessSlot prepares the DeleteProcessSlot request.
func (c WebAppsClient) preparerForDeleteProcessSlot(ctx context.Context, id SlotProcessId) (*http.Request, error) {
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

// responderForDeleteProcessSlot handles the response to the DeleteProcessSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteProcessSlot(resp *http.Response) (result DeleteProcessSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
