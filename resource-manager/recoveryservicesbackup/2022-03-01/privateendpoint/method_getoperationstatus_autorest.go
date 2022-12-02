package privateendpoint

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOperationStatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *OperationStatus
}

// GetOperationStatus ...
func (c PrivateEndpointClient) GetOperationStatus(ctx context.Context, id OperationsStatusId) (result GetOperationStatusOperationResponse, err error) {
	req, err := c.preparerForGetOperationStatus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoint.PrivateEndpointClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoint.PrivateEndpointClient", "GetOperationStatus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetOperationStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoint.PrivateEndpointClient", "GetOperationStatus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetOperationStatus prepares the GetOperationStatus request.
func (c PrivateEndpointClient) preparerForGetOperationStatus(ctx context.Context, id OperationsStatusId) (*http.Request, error) {
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

// responderForGetOperationStatus handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (c PrivateEndpointClient) responderForGetOperationStatus(resp *http.Response) (result GetOperationStatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
