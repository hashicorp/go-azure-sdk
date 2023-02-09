package operationstatus

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationOperationStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *LocationOperationStatus
}

// LocationOperationStatus ...
func (c OperationStatusClient) LocationOperationStatus(ctx context.Context, id OperationId) (result LocationOperationStatusOperationResponse, err error) {
	req, err := c.preparerForLocationOperationStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationstatus.OperationStatusClient", "LocationOperationStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationstatus.OperationStatusClient", "LocationOperationStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForLocationOperationStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationstatus.OperationStatusClient", "LocationOperationStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForLocationOperationStatus prepares the LocationOperationStatus request.
func (c OperationStatusClient) preparerForLocationOperationStatus(ctx context.Context, id OperationId) (*http.Request, error) {
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

// responderForLocationOperationStatus handles the response to the LocationOperationStatus request. The method always
// closes the http.Response Body.
func (c OperationStatusClient) responderForLocationOperationStatus(resp *http.Response) (result LocationOperationStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
