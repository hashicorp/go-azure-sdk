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

type StartContinuousWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
}

// StartContinuousWebJobSlot ...
func (c WebAppsClient) StartContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (result StartContinuousWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForStartContinuousWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartContinuousWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartContinuousWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartContinuousWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StartContinuousWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartContinuousWebJobSlot prepares the StartContinuousWebJobSlot request.
func (c WebAppsClient) preparerForStartContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (*http.Request, error) {
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

// responderForStartContinuousWebJobSlot handles the response to the StartContinuousWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStartContinuousWebJobSlot(resp *http.Response) (result StartContinuousWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
