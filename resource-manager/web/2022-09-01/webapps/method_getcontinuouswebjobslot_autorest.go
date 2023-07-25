package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetContinuousWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContinuousWebJob
}

// GetContinuousWebJobSlot ...
func (c WebAppsClient) GetContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (result GetContinuousWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForGetContinuousWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContinuousWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContinuousWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetContinuousWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetContinuousWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetContinuousWebJobSlot prepares the GetContinuousWebJobSlot request.
func (c WebAppsClient) preparerForGetContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetContinuousWebJobSlot handles the response to the GetContinuousWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetContinuousWebJobSlot(resp *http.Response) (result GetContinuousWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
