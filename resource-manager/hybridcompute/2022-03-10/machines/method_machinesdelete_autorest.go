package machines

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachinesDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// MachinesDelete ...
func (c MachinesClient) MachinesDelete(ctx context.Context, id MachineId) (result MachinesDeleteOperationResponse, err error) {
	req, err := c.preparerForMachinesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMachinesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMachinesDelete prepares the MachinesDelete request.
func (c MachinesClient) preparerForMachinesDelete(ctx context.Context, id MachineId) (*http.Request, error) {
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

// responderForMachinesDelete handles the response to the MachinesDelete request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForMachinesDelete(resp *http.Response) (result MachinesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
