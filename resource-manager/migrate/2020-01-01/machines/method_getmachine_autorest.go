package machines

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMachineOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMwareMachine
}

// GetMachine ...
func (c MachinesClient) GetMachine(ctx context.Context, id commonids.VMwareSiteMachineId) (result GetMachineOperationResponse, err error) {
	req, err := c.preparerForGetMachine(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetMachine", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetMachine", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMachine(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "GetMachine", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMachine prepares the GetMachine request.
func (c MachinesClient) preparerForGetMachine(ctx context.Context, id commonids.VMwareSiteMachineId) (*http.Request, error) {
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

// responderForGetMachine handles the response to the GetMachine request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForGetMachine(resp *http.Response) (result GetMachineOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
