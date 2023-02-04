package machines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopMachineOperationResponse struct {
	HttpResponse *http.Response
}

// StopMachine ...
func (c MachinesClient) StopMachine(ctx context.Context, id commonids.VMwareSiteMachineId) (result StopMachineOperationResponse, err error) {
	req, err := c.preparerForStopMachine(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "StopMachine", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "StopMachine", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopMachine(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "StopMachine", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopMachine prepares the StopMachine request.
func (c MachinesClient) preparerForStopMachine(ctx context.Context, id commonids.VMwareSiteMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStopMachine handles the response to the StopMachine request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForStopMachine(resp *http.Response) (result StopMachineOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
