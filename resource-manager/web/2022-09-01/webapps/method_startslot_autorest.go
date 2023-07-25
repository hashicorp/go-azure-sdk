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

type StartSlotOperationResponse struct {
	HttpResponse *http.Response
}

// StartSlot ...
func (c WebAppsClient) StartSlot(ctx context.Context, id SlotId) (result StartSlotOperationResponse, err error) {
	req, err := c.preparerForStartSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartSlot prepares the StartSlot request.
func (c WebAppsClient) preparerForStartSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStartSlot handles the response to the StartSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStartSlot(resp *http.Response) (result StartSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
