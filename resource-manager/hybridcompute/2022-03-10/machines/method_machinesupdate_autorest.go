package machines

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachinesUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *Machine
}

// MachinesUpdate ...
func (c MachinesClient) MachinesUpdate(ctx context.Context, id MachineId, input MachineUpdate) (result MachinesUpdateOperationResponse, err error) {
	req, err := c.preparerForMachinesUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMachinesUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMachinesUpdate prepares the MachinesUpdate request.
func (c MachinesClient) preparerForMachinesUpdate(ctx context.Context, id MachineId, input MachineUpdate) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForMachinesUpdate handles the response to the MachinesUpdate request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForMachinesUpdate(resp *http.Response) (result MachinesUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
